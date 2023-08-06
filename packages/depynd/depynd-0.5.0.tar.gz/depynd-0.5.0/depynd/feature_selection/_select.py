from sklearn.utils.validation import check_X_y

from depynd.feature_selection import mrmr, mifs


def select(X, y, lamb=0.0, method='mifs', **kwargs):
    """Select effective features in ``X`` on predicting ``y``.

    Parameters
    ----------
    X : array-like, shape (n_samples, d)
        Observations of feature variables.
    y : array-like, shape (n_samples)
        Observations of the target variable.
    lamb: float, default 0.0
        Threshold for independence tests.
    method: {'mifs', 'mrmr'}, default 'mifs'
        Feature selection method.

    Returns
    -------
    indices : list
        Indices of the selected features.
    """
    X, y = check_X_y(X, y)
    if method == 'mifs':
        return mifs(X, y, lamb=lamb, **kwargs)
    elif method == 'mrmr':
        return mrmr(X, y, lamb=lamb, **kwargs)
    else:
        raise NotImplementedError('Method %s is not implemented. Use mifs or mrmr.' % method)
