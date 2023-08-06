# -*- coding: utf-8 -*-
"""
Example of using one class SVM for outlier detection
"""
from __future__ import division
from __future__ import print_function

import os
import sys

# temporary solution for relative imports in case pyod is not installed
# if pyod is installed, no need to use the following line
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import numpy as np
from sklearn.metrics import roc_auc_score

from pyod.models.ocsvm import OCSVM

from pyod.utils.data import generate_data
from pyod.utils.data import visualize
from pyod.utils.utility import precision_n_scores

if __name__ == "__main__":
    contamination = 0.1  # percentage of outliers
    n_train = 200  # number of training points
    n_test = 100  # number of testing points

    X_train, y_train, X_test, y_test = generate_data(
        n_train=n_train, n_test=n_test, contamination=contamination)

    # train one_class_svm detector
    clf_name = 'OneClassSVM'
    clf = OCSVM()
    clf.fit(X_train)

    # get the prediction label and decision_scores_ on the training data
    y_train_pred = clf.labels_  # binary labels (0: inliers, 1: outliers)
    y_train_scores = clf.decision_scores_  # raw outlier scores

    # get the prediction on the test data
    y_test_pred = clf.predict(X_test)  # outlier labels (0 or 1)
    y_test_scores = clf.decision_function(X_test)  # outlier scores

    # evaluate and print the results
    print('{clf_name} Train ROC:{roc}, precision @ rank n:{prn}'.format(
        clf_name=clf_name,
        roc=np.round(roc_auc_score(y_train, y_train_scores), decimals=4),
        prn=np.round(precision_n_scores(y_train, y_train_scores), decimals=4)))

    print('{clf_name} Test ROC:{roc}, precision @ rank n:{prn}'.format(
        clf_name=clf_name,
        roc=np.round(roc_auc_score(y_test, y_test_scores), decimals=4),
        prn=np.round(precision_n_scores(y_test, y_test_scores), decimals=4)))

    # visualize the results
    visualize(clf_name, X_train, y_train, X_test, y_test, y_train_pred,
              y_test_pred, save_figure=False)
