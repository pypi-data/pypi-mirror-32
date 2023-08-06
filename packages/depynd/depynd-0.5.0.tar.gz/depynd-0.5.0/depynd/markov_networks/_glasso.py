import numpy as np
from sklearn.preprocessing import scale
from sklearn.covariance import graph_lasso


def glasso(X, lamb, return_precision=False):
    """Learn the structure of Markov random field with the graphical lasso.

    This function internally calls the implementation in scikit-learn.

    Parameters
    ----------
    X : array, shape (n_samples, d)
        Observations of variables.
    lamb : float
        Regularization parameter.
    return_precision : bool, default False
        If True, the estimated precision matrix will be returned instead of adjacency matrix.

    Returns
    ----------
    adj : array, shape (d, d)
        Estimated adjacency matrix (or precision matrix if ``return_precision`` is True) of an MRF.
    """
    cov = np.cov(scale(X), rowvar=False)
    pre = graph_lasso(cov, alpha=lamb)[1]
    if return_precision:
        return pre
    else:
        adj = ~np.isclose(pre, 0)
        adj[np.eye(len(adj), dtype=bool)] = 0
        return adj
