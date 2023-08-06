# -*- coding: utf-8 -*-

from __future__ import division
from __future__ import print_function

import os
import sys

# temporary solution for relative imports in case pyod is not installed
# if pyod is installed, no need to use the following line
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import unittest
from sklearn.utils.testing import assert_equal
from sklearn.utils.testing import assert_greater
from sklearn.utils.testing import assert_greater_equal
from sklearn.utils.testing import assert_less_equal
from sklearn.utils.testing import assert_raises
from sklearn.utils.estimator_checks import check_estimator
from sklearn.metrics import roc_auc_score

from pyod.models.knn import KNN
from pyod.utils.data import generate_data


class TestKnn(unittest.TestCase):
    def setUp(self):
        self.n_train = 100
        self.n_test = 50
        self.contamination = 0.1
        self.roc_floor = 0.75
        self.X_train, self.y_train, self.X_test, self.y_test = generate_data(
            n_train=self.n_train, n_test=self.n_test,
            contamination=self.contamination)

        self.clf = KNN(contamination=self.contamination)
        self.clf.fit(self.X_train)

    def test_sklearn_estimator(self):
        check_estimator(self.clf)

    def test_parameters(self):
        if not hasattr(self.clf,
                       'decision_scores_') or self.clf.decision_scores_ is None:
            self.assertRaises(AttributeError, 'decision_scores_ is not set')
        if not hasattr(self.clf, 'labels_') or self.clf.labels_ is None:
            self.assertRaises(AttributeError, 'labels_ is not set')
        if not hasattr(self.clf, 'threshold_') or self.clf.threshold_ is None:
            self.assertRaises(AttributeError, 'threshold_ is not set')
        if not hasattr(self.clf, '_mu') or self.clf._mu is None:
            self.assertRaises(AttributeError, '_mu is not set')
        if not hasattr(self.clf, '_sigma') or self.clf._sigma is None:
            self.assertRaises(AttributeError, '_sigma is not set')

    def test_train_scores(self):
        assert_equal(len(self.clf.decision_scores_), self.X_train.shape[0])

    def test_prediction_scores(self):
        pred_scores = self.clf.decision_function(self.X_test)

        # check score shapes
        assert_equal(pred_scores.shape[0], self.X_test.shape[0])

        # check performance
        assert_greater(roc_auc_score(self.y_test, pred_scores), self.roc_floor)

    def test_prediction_labels(self):
        pred_labels = self.clf.predict(self.X_test)
        assert_equal(pred_labels.shape, self.y_test.shape)

    def test_prediction_proba(self):
        pred_proba = self.clf.predict_proba(self.X_test)
        assert_greater_equal(pred_proba.min(), 0)
        assert_less_equal(pred_proba.max(), 1)

    def test_prediction_proba_linear(self):
        pred_proba = self.clf.predict_proba(self.X_test, method='linear')
        assert_greater_equal(pred_proba.min(), 0)
        assert_less_equal(pred_proba.max(), 1)

    def test_prediction_proba_unify(self):
        pred_proba = self.clf.predict_proba(self.X_test, method='unify')
        assert_greater_equal(pred_proba.min(), 0)
        assert_less_equal(pred_proba.max(), 1)

    def test_prediction_proba_parameter(self):
        with assert_raises(ValueError):
            self.clf.predict_proba(self.X_test, method='something')

    def test_fit_predict(self):
        pred_labels = self.clf.fit_predict(self.X_train)
        assert_equal(pred_labels.shape, self.y_train.shape)

    def test_evaluate(self):
        self.clf.fit_predict_evaluate(self.X_test, self.y_test)

    def tearDown(self):
        pass


class TestKnnMean(unittest.TestCase):

    def setUp(self):
        self.n_train = 100
        self.n_test = 50
        self.contamination = 0.1
        self.roc_floor = 0.75
        self.X_train, self.y_train, self.X_test, self.y_test = generate_data(
            n_train=self.n_train, n_test=self.n_test,
            contamination=self.contamination)

        self.clf = KNN(contamination=self.contamination, method='mean')

    def test_fit(self):
        self.clf.fit(self.X_train)

    def test_decision_function(self):
        self.clf.fit(self.X_train)
        self.clf.decision_function(self.X_train)
        self.clf.decision_function(self.X_test)

    def test_sklearn_estimator(self):
        check_estimator(self.clf)

    def tearDown(self):
        pass


class TestKnnMedian(unittest.TestCase):

    def setUp(self):
        self.n_train = 100
        self.n_test = 50
        self.contamination = 0.1
        self.roc_floor = 0.75
        self.X_train, self.y_train, self.X_test, self.y_test = generate_data(
            n_train=self.n_train, n_test=self.n_test,
            contamination=self.contamination)

        self.clf = KNN(contamination=self.contamination, method='median')

    def test_fit(self):
        self.clf.fit(self.X_train)

    def test_decision_function(self):
        self.clf.fit(self.X_train)
        self.clf.decision_function(self.X_train)
        self.clf.decision_function(self.X_test)

    def test_sklearn_estimator(self):
        check_estimator(self.clf)

    def tearDown(self):
        pass


if __name__ == '__main__':
    unittest.main()

##############################################################################
# samples = [[-1, 0], [0., 0.], [1., 1], [2., 5.], [3, 1]]
#
# clf = Knn()
# clf.fit(samples)
#
# decision_scores_ = clf.decision_function(np.asarray([[2, 3], [6, 8]])).ravel()
# assert (decision_scores_[0] == [2])
# assert (decision_scores_[1] == [5])
# #
# labels = clf.predict(np.asarray([[2, 3], [6, 8]])).ravel()
# assert (labels[0] == [0])
# assert (labels[1] == [1])
