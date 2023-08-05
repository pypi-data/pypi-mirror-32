"""
Most of the clustering procedure is shared between COBRAS_dtw and COBRAS_kshape, this is captured in the COBRAS class.
The following methods are specific to each variant:

- create_superinstance: super-instances for COBRAS_dtw and COBRAS_kshape are different, this method
                        simply creates a super-instance of the appropriate type
- split_superinstance: this is also different or COBRAS_dtw and COBRAS_kshape, the first uses
                        spectral clustering, the second kshape to split a super-instance
"""

import abc
import itertools
import random
import time
import copy

import numpy as np
from cobras_ts.cluster import Cluster

from cobras_ts.clustering import Clustering

class COBRAS:
    def __init__(self, data, querier, max_questions, train_indices=None, store_intermediate_results=True):
        """COBRAS clustering

        :param data: Data set numpy array of size (nb_instances,nb_features)
        :param querier: Querier object that answers whether instances are linked through a must-link or cannot-link.
        :param max_questions: Maximum number of questions that are asked. Clustering stops after this.
        :param train_indices:
            Default: None
        """
        self.data = data
        self.querier = querier
        self.max_questions = max_questions
        self.store_intermediate_results = store_intermediate_results

        if train_indices is None:
            self.train_indices = range(self.data.shape[0])
        else:
            self.train_indices = train_indices

        self.clustering = None
        self.split_cache = dict()
        self.start = None
        self.intermediate_results = []
        self.ml = None
        self.cl = None

    def cluster(self):
        """Perform clustering.

        :return: if cobras.store_intermediate_results is set to False, this method returns a single cluster labeling
                 if cobras.store_intermediate_results is set to True, this method returns:
                        - a list of clusterings (each item of the type cobras_ts.cluster.Cluster) for each query
                        - a list of timestamps for each query
                        - the list of must-link constraints that was queried
                        - the list of cannot-link constraints that was queried
        """
        self.start = time.time()

        # initially, there is only one super-instance that contains all data indices (i.e. list(range(self.data.shape[0])))
        initial_superinstance = self.create_superinstance(list(range(self.data.shape[0])))

        self.ml = []
        self.cl = []

        self.clustering = Clustering([Cluster([initial_superinstance])])

        # the split level for this initial super-instance is determined,
        # the super-instance is split, and a new cluster is created for each of the newly created superinstances
        initial_k = self.determine_split_level(initial_superinstance, copy.deepcopy(self.clustering))

        superinstances = self.split_superinstance(initial_superinstance,initial_k)
        self.clustering.clusters = []
        for si in superinstances:
            self.clustering.clusters.append(Cluster([si]))

        if self.store_intermediate_results:
            # in results we store a tuple for each constraint number
            # for the first couple of queries (those used to determine the initial splitting level)
            # we do not have a clustering, we just return 'all elements in one cluster'
            for i in range(len(self.ml) + len(self.cl)):
                self.intermediate_results.append((copy.deepcopy(self.clustering), time.time() - self.start, len(self.ml) + len(self.cl)))


        # the first bottom up merging step
        self.merge_containing_clusters(starting_level=True)

        while len(self.ml) + len(self.cl) < self.max_questions:

            self.querier.update_clustering(self.clustering)

            to_split, originating_cluster = self.identify_superinstance_to_split()
            if to_split is None:
                break

            if self.intermediate_results:
                clustering_to_store = copy.deepcopy(self.clustering)

            originating_cluster.super_instances.remove(to_split)
            if len(originating_cluster.super_instances) == 0:
                self.clustering.clusters.remove(originating_cluster)

            split_level = self.determine_split_level(to_split, clustering_to_store)
            new_super_instances = self.split_superinstance(to_split, split_level)
            new_clusters = self.add_new_clusters_from_split(new_super_instances)

            if not new_clusters:
                # it is possible that splitting a super-instance does not lead to a new cluster:
                # e.g. a super-instance constains 2 points, of which one is in the test set
                # in this case, the super-instance can be split into two new ones, but these will be joined
                # again immediately, as we cannot have super-instances containing only test points (these cannot be
                # queried)
                # this case handles this, we simply add the super-instance back to its originating cluster,
                # and set the already_tried flag to make sure we do not keep trying to split this superinstance

                originating_cluster.super_instances.append(to_split)
                to_split.tried_splitting = True

                if originating_cluster not in self.clustering.clusters:
                    self.clustering.clusters.append(originating_cluster)

                continue
            else:
                self.clustering.clusters.extend(new_clusters)

            self.merge_containing_clusters(starting_level=False)

        if self.store_intermediate_results:
            return [clust for clust, _, _ in self.intermediate_results], [runtime for _, runtime, _ in self.intermediate_results], self.ml, self.cl
        else:
            return self.clustering.construct_cluster_labeling()

    @abc.abstractmethod
    def split_superinstance(self, si, k):
        return

    @abc.abstractmethod
    def create_superinstance(self, indices):
        return

    def determine_split_level(self, superinstance, intermediate_clustering):
        # need to make a 'deep copy' here, we will split this one a few times just to determine an appropriate splitting
        # level
        si = self.create_superinstance(superinstance.indices)


        must_link_found = False

        split_level = 0
        while not must_link_found:

            if len(si.indices) == 2:
                new_si = [self.create_superinstance([si.indices[0]]), self.create_superinstance([si.indices[1]])]
            else:
                new_si = self.split_superinstance(si,2)

            new_clusters = []
            for si in new_si:
                new_clusters.append(Cluster([si]))

            if len(new_clusters) == 1:
                # we cannot split any further along this branch, we reached the splitting level
                split_level = max([split_level, 1])
                split_n = 2 ** int(split_level)
                return min(len(si.indices),split_n)

            x = new_clusters[0]
            y = new_clusters[1]
            bc1, bc2 = x.get_comparison_points(y)
            pt1 = min([bc1.representative_idx, bc2.representative_idx])
            pt2 = max([bc1.representative_idx, bc2.representative_idx])

            if self.querier.query_points(pt1,pt2):
                self.ml.append((pt1, pt2))
                must_link_found = True
                if self.store_intermediate_results:
                    self.intermediate_results.append((intermediate_clustering, time.time() - self.start, len(self.ml) + len(self.cl)))
            else:
                self.cl.append((pt1, pt2))
                split_level += 1
                if self.store_intermediate_results:
                    self.intermediate_results.append((intermediate_clustering, time.time() - self.start, len(self.ml) + len(self.cl)))

            si_to_choose = []
            if len(x.super_instances[0].train_indices) >= 2:
                si_to_choose.append(x.super_instances[0])
            if len(y.super_instances[0].train_indices) >= 2:
                si_to_choose.append(y.super_instances[0])

            if len(si_to_choose) == 0:
                split_level = max([split_level, 1])
                split_n = 2 ** int(split_level)
                return min(len(si.indices), split_n)

            si = random.choice(si_to_choose)

        split_level = max([split_level, 1])
        split_n = 2 ** int(split_level)
        return min(len(si.indices), split_n)

    def add_new_clusters_from_split(self, si):
        new_clusters = []
        for x in si:
            new_clusters.append(Cluster([x]))

        if len(new_clusters) == 1:
            return None
        else:
            return new_clusters

    def merge_containing_clusters(self, starting_level=False):

        clustering_to_store = copy.deepcopy(self.clustering)

        merged = True
        while merged and len(self.ml) + len(self.cl) < self.max_questions:

            clusters_to_consider = [cluster for cluster in self.clustering.clusters if not cluster.is_finished]

            cluster_pairs = itertools.combinations(clusters_to_consider, 2)
            cluster_pairs = [x for x in cluster_pairs if
                             not x[0].cannot_link_to_other_cluster(x[1], self.cl)]
            cluster_pairs = sorted(cluster_pairs, key=lambda x: x[0].distance_to(x[1]))

            merged = False
            for x, y in cluster_pairs:

                if x.cannot_link_to_other_cluster(y, self.cl):
                    continue

                bc1, bc2 = x.get_comparison_points(y)
                pt1 = min([bc1.representative_idx, bc2.representative_idx])
                pt2 = max([bc1.representative_idx, bc2.representative_idx])


                if (pt1, pt2) in self.ml:
                    x.super_instances.extend(y.super_instances)
                    self.clustering.clusters.remove(y)
                    merged = True
                    break

                if len(self.ml) + len(self.cl) == self.max_questions:
                    break

                if self.querier.query_points(pt1,pt2):
                    x.super_instances.extend(y.super_instances)
                    self.clustering.clusters.remove(y)
                    self.ml.append((pt1, pt2))
                    merged = True

                    if self.store_intermediate_results:
                        if starting_level:
                            # if it is the first merging step there is no previous clustering that is being refined,
                            # so temporary results are the ones being constructed now
                            self.intermediate_results.append(
                                (clustering_to_store, time.time() - self.start,
                                 len(self.ml) + len(self.cl)))
                        else:
                            self.intermediate_results.append((clustering_to_store, time.time() - self.start, len(self.ml) + len(self.cl)))

                    break
                else:
                    self.cl.append((pt1, pt2))

                    if self.store_intermediate_results:
                        if starting_level:
                            self.intermediate_results.append(
                                (clustering_to_store, time.time() - self.start,
                                 len(self.ml) + len(self.cl)))
                        else:
                            self.intermediate_results.append((clustering_to_store, time.time() - self.start, len(self.ml) + len(self.cl)))



            if self.store_intermediate_results and not starting_level:
                self.intermediate_results[-1] = (copy.deepcopy(self.clustering), time.time() - self.start,
                                                 len(self.ml) + len(self.cl))

    def identify_superinstance_to_split(self):

        if len(self.clustering.clusters) == 1 and len(self.clustering.clusters[0].super_instances) == 1:
            return self.clustering.clusters[0].super_instances[0], self.clustering.clusters[0]

        superinstance_to_split = None
        max_heur = -np.inf
        originating_cluster = None

        for cluster in self.clustering.clusters:

            if cluster.is_pure:
                continue

            if cluster.is_finished:
                continue

            for superinstance in cluster.super_instances:
                if superinstance.tried_splitting:
                    continue

                if len(superinstance.indices) > max_heur:
                    superinstance_to_split = superinstance
                    max_heur = len(superinstance.indices)
                    originating_cluster = cluster

        if superinstance_to_split is None:
            return None, None


        return superinstance_to_split, originating_cluster


