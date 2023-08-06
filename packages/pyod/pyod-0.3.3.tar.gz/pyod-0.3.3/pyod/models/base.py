"""
Abstract base class for outlier detector models
"""

from abc import ABC, abstractmethod

import numpy as np
from scipy.stats import rankdata
from scipy.special import erf
from scipy.stats import scoreatpercentile

from sklearn.preprocessing import MinMaxScaler
from sklearn.metrics import roc_auc_score
from sklearn.utils.validation import check_is_fitted

from ..utils.utility import precision_n_scores


class BaseDetector(ABC):
    """
    Abstract class for all outlier detection algorithms.

    :param contamination: The amount of contamination of the data set,
        i.e. the proportion of outliers in the data set. Used when fitting to
        define the threshold on the decision function.
    :type contamination: float in (0., 0.5), optional (default=0.1)
    """

    @abstractmethod
    def __init__(self, contamination=0.1):

        if not (0. < contamination <= 0.5):
            raise ValueError("contamination must be in (0, 0.5], "
                             "got: %f" % contamination)

        self.contamination = contamination

    @abstractmethod
    def decision_function(self, X):
        """
        Predict Anomaly score of X of the base classifiers. The anomaly score
        of an input sample is computed based on different detector algorithms.
        For consistency, outliers have larger anomaly decision_scores.

        :param X: The training input samples. Sparse matrices are accepted only
            if they are supported by the base estimator.
        :type X: numpy array of shape (n_samples, n_features)

        :return: decision_scores: The anomaly score of the input samples.
        :rtype: array, shape (n_samples,)
        """
        pass

    @abstractmethod
    def fit(self, X):
        """
        Fit detector.

        :param X: The training input samples. Sparse matrices are accepted only
            if they are supported by the base estimator.
        :type X: numpy array of shape (n_samples, n_features)

        :return: return self
        :rtype: object
        """

        pass

    def fit_predict(self, X):
        """
        Fit detector and predict if a particular sample is an outlier or not.

        :param X: The input samples
        :type X: numpy array of shape (n_samples, n_features)

        :return: For each observation, tells whether or not
            it should be considered as an outlier according to the fitted model.
            0 stands for inliers and 1 for outliers.
        :rtype: array, shape (n_samples,)
        """

        self.fit(X)
        return self.y_pred

    def predict(self, X):
        """
        Predict if a particular sample is an outlier or not.

        :param X: The input samples
        :type X: numpy array of shape (n_samples, n_features)

        :return: For each observation, tells whether or not
            it should be considered as an outlier according to the fitted
            model. 0 stands for inliers and 1 for outliers.
        :rtype: array, shape (n_samples,)
        """

        check_is_fitted(self, ['decision_scores', 'threshold_', 'y_pred'])

        pred_score = self.decision_function(X)
        return (pred_score > self.threshold_).astype('int').ravel()

    def predict_proba(self, X, method='linear'):
        """
        Predict the probability of a sample being outlier. Two approaches
        are possible:

        1. simply use Min-max conversion to linearly transform the outlier
           decision_scores into the range of [0,1]. The model must be fitted first.
        2. use unifying decision_scores, see reference [1] below.

        [1] Kriegel, H.P., Kroger, P., Schubert, E. and Zimek, A., 2011, April.
        Interpreting and unifying outlier decision_scores. In Proc' SIAM, 2011.

        :param X: The input samples
        :type X: numpy array of shape (n_samples, n_features)

        :param method: probability conversion method. It must be one of
            'linear' or 'unify'.
        :type method: str, optional (default='linear')

        :return: For each observation, return the outlier probability, ranging
            in [0,1]
        :rtype: array, shape (n_samples,)
        """

        check_is_fitted(self, ['decision_scores', 'threshold_', 'y_pred'])

        test_scores = self.decision_function(X)
        train_scores = self.decision_scores

        if method == 'linear':
            scaler = MinMaxScaler().fit(train_scores.reshape(-1, 1))
            proba = scaler.transform(test_scores.reshape(-1, 1))
            return proba.clip(0, 1)

        elif method == 'unify':
            # turn output into probability
            pre_erf_score = (test_scores - self._mu) / (
                    self._sigma * np.sqrt(2))
            erf_score = erf(pre_erf_score)
            proba = erf_score.clip(0)
            return proba
        else:
            raise ValueError(method,
                             'is not a valid probability conversion method')

    def predict_rank(self, X):
        """
        Predict the outlyingness rank of a sample in a fitted model. The
        method is specifically for combining various outlier detectors.

        :param X: The input samples
        :type X: numpy array of shape (n_samples, n_features)

        :return: outlying rank of a sample according to the training data
        :rtype: array, shape (n_samples,)
        """

        check_is_fitted(self, ['decision_scores', 'threshold_', 'y_pred'])

        test_scores = self.decision_function(X)
        train_scores = self.decision_scores

        ranks = np.zeros([X.shape[0], 1])

        for i in range(test_scores.shape[0]):
            train_scores_i = np.append(train_scores.reshape(-1, 1),
                                       test_scores[i])

            ranks[i] = rankdata(train_scores_i)[-1]

        # return normalized ranks
        ranks_norm = ranks / ranks.max()
        return ranks_norm

    def fit_predict_evaluate(self, X, y):
        """
        Fit the detector, predict on samples, and evaluate the model

        :param X: The input samples
        :type X: numpy array of shape (n_samples, n_features)

        :param y: Outlier labels of the input samples
        :type y: array, shape (n_samples,)

        :return: roc score and precision @ rank n score
        :rtype:  tuple (float, float)
        """

        self.fit(X)
        roc = roc_auc_score(y, self.decision_scores)
        prec_n = precision_n_scores(y, self.decision_scores)

        print("roc score:", roc)
        print("precision @ rank n:", prec_n)

        return roc, prec_n

    def _process_decision_scores(self):
        """
        Internal function to calculate key attributes:
        threshold: used to decide the binary label
        y_pred: binary lables of training data

        :return: self
        :rtype: object
        """

        self.threshold_ = scoreatpercentile(self.decision_scores,
                                            100 * (1 - self.contamination))
        self.y_pred = (self.decision_scores > self.threshold_).astype(
            'int').ravel()

        # calculate for predict_proba()

        self._mu = np.mean(self.decision_scores)
        self._sigma = np.std(self.decision_scores)

        return self
