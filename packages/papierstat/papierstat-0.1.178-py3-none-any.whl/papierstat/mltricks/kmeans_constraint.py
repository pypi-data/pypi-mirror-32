# -*- coding: utf-8 -*-
"""
@file
@brief Implémente la classe @see cl ConstraintKMeans.
"""
import numpy
from sklearn.cluster import KMeans
from sklearn.metrics.pairwise import euclidean_distances
from .kmeans_constraint_ import constraint_kmeans, constraint_predictions


class ConstraintKMeans(KMeans):
    """
    Defines a constraint :epkg:`k-means`.
    Clusters are modified to have an equal size.
    The algorithm is initialized with a regular :epkg:`k-means`
    and continues with a modified version of it.

    Computing the predictions offer a choice.
    The first one is to keep the predictions
    from the regular *k-means* algorithm
    but with the balanced clusters.
    The second is to compute balanced predictions
    over the test set. That implies that the predictions
    for the same observations might change depending
    on the set it belongs to.

    .. runpython::
        :rst:

        from papierstat.datasets.documentation import list_notebooks_rst_links
        links = list_notebooks_rst_links('digressions', 'constraint_kmeans')
        links = ['    * %s' % s for s in links]
        print('\\n'.join(links))
    """

    _strategy_value = {'distance', 'gain'}

    def __init__(self, n_clusters=8, init='k-means++', n_init=10, max_iter=300,
                 tol=0.0001, precompute_distances='auto', verbose=0,
                 random_state=None, copy_x=True, n_jobs=1, algorithm='auto',
                 balanced_predictions=False, strategy='gain', kmeans0=True):
        """
        @param      n_clusters              number of clusters
        @param      init                    used by :epkg:`k-means`
        @param      n_init                  used by :epkg:`k-means`
        @param      max_iter                used by :epkg:`k-means`
        @param      tol                     used by :epkg:`k-means`
        @param      precompute_distances    used by :epkg:`k-means`
        @param      verbose                 used by :epkg:`k-means`
        @param      random_state            used by :epkg:`k-means`
        @param      copy_x                  used by :epkg:`k-means`
        @param      n_jobs                  used by :epkg:`k-means`
        @param      algorithm               used by :epkg:`k-means`
        @param      balanced_predictions    produced balanced prediction
                                            or the regular ones
        @param      strategy                strategy or algorithm used to abide
                                            by the constraint
        @param      kmeans0                 if True, applies *k-means* algorithm first

        The parameter *strategy* determines how
        obseervations should be assigned to a cluster.
        The value can be:

        * ``'distance'``: observations are ranked by distance to a cluster,
          the algorithm assigns first point to the closest center unless it reached
          the maximulmsize
        * ``'gain'``: follows the algorithm described at
           see `Same-size k-Means Variation <https://elki-project.github.io/tutorial/same-size_k_means>`_
        """
        KMeans.__init__(self, n_clusters=n_clusters, init=init, n_init=n_init,
                        max_iter=max_iter, tol=tol, precompute_distances=precompute_distances,
                        verbose=verbose, random_state=random_state, copy_x=copy_x,
                        n_jobs=n_jobs, algorithm=algorithm)
        self.balanced_predictions = balanced_predictions
        self.strategy = strategy
        self.kmeans0 = kmeans0
        if strategy not in ConstraintKMeans._strategy_value:
            raise ValueError('strategy must be in {0}'.format(
                ConstraintKMeans._strategy_value))

    def fit(self, X, y=None, fLOG=None):
        """
        Compute k-means clustering.

        Parameters
        ----------
        X : array-like or sparse matrix, shape=(n_samples, n_features)
            Training instances to cluster. It must be noted that the data
            will be converted to C ordering, which will cause a memory
            copy if the given data is not C-contiguous.

        y : Ignored

        fLOG: logging function
        """
        max_iter = self.max_iter
        self.max_iter //= 2
        if self.kmeans0:
            KMeans.fit(self, X, y)
            state = None
        else:
            state = numpy.random.RandomState(self.random_state)
            labels = state.randint(
                0, self.n_clusters, X.shape[0], dtype=numpy.int32)
            centers = numpy.empty((self.n_clusters, X.shape[1]), dtype=X.dtype)
            choice = state.randint(0, self.n_clusters, self.n_clusters)
            for i, c in enumerate(choice):
                centers[i, :] = X[c, :]
            self.labels_ = labels
            self.cluster_centers_ = centers
            self.inertia_ = float(X.shape[0])
            self.n_iter_ = 0

        self.max_iter = max_iter
        return self.constraint_kmeans(X, state=state, fLOG=fLOG)

    def constraint_kmeans(self, X, state=None, fLOG=None):
        """
        Completes the constraint k-means.

        @param      X       features
        @param      fLOG    logging function
        """
        labels, centers, inertia, iter_ = constraint_kmeans(X, self.labels_, self.cluster_centers_, self.inertia_,
                                                            self.precompute_distances, self.n_iter_, self.max_iter,
                                                            verbose=self.verbose, strategy=self.strategy,
                                                            state=state, fLOG=fLOG)
        self.labels_ = labels
        self.cluster_centers_ = centers
        self.inertia_ = inertia
        self.n_iter_ = iter_
        return self

    def predict(self, X):
        """
        Computes the predictions.

        @param      X       features.
        @return             prediction
        """
        if self.balanced_predictions:
            labels, _, __ = constraint_predictions(
                X, self.cluster_centers_, strategy=self.strategy + '_p')
            return labels
        else:
            return KMeans.predict(self, X)

    def transform(self, X):
        """
        Computes the predictions.

        @param      X       features.
        @return             prediction
        """
        if self.balanced_predictions:
            labels, distances, __ = constraint_predictions(
                X, self.cluster_centers_, strategy=self.strategy)
            # We remove small distances than the chosen clusters
            # due to the constraint, we choose max*2 instead.
            mx = distances.max() * 2
            for i, l in enumerate(labels):
                mi = distances[i, l]
                mmi = distances[i, :].min()
                if mi > mmi:
                    # numpy.nan would be best
                    distances[i, distances[i, :] < mi] = mx
            return distances
        else:
            return KMeans.transform(self, X)

    def score(self, X, y=None):
        """
        Returns the distances to all clusters.

        @param      X       features
        @param      y       unused
        @return             distances
        """
        if self.balanced_predictions:
            _, __, dist_close = constraint_predictions(
                X, self.cluster_centers_, strategy=self.strategy)
            return dist_close
        else:
            return euclidean_distances(self.cluster_centers_, X, squared=True)
