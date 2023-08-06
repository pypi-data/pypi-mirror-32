# -*- coding: utf-8 -*-
"""
@file
@brief Implémente un *transform* qui suit la même API que tout :epkg:`scikit-learn` transform.
"""
import textwrap
import numpy
from .sklearn_base_transform import SkBaseTransform


class SkBaseTransformLearner(SkBaseTransform):
    """
    Un *transform* qui cache un *learner*, il convertit
    la fonction *predict* en *transform*. De cette façon,
    deux learners peuvent être insérés dans le même pipeline.

    .. exref::
        :title: Utiliser deux learners dans un pipeline scikit-learn.
        :tag: sklearn
        :lid: ex-pipe2learner

        Il est impossible d'utiliser deux *learners* dans un pipeline.
        Mais on peut utiliser la classe @see cl SkBaseTransformLearner
        pour maquiller un *learner* en *transform*.

        .. runpython::
            :showcode:

            from sklearn.model_selection import train_test_split
            from sklearn.datasets import load_iris
            from sklearn.linear_model import LogisticRegression
            from sklearn.tree import DecisionTreeClassifier
            from sklearn.metrics import accuracy_score
            from sklearn.pipeline import make_pipeline
            from papierstat.mltricks import SkBaseTransformLearner

            data = load_iris()
            X, y = data.data, data.target
            X_train, X_test, y_train, y_test = train_test_split(X, y)

            try:
                pipe = make_pipeline(LogisticRegression(),
                                     DecisionTreeClassifier())
            except Exception as e:
                print("ERREUR :", e)

            pipe = make_pipeline(SkBaseTransformLearner(LogisticRegression()),
                                 DecisionTreeClassifier())
            pipe.fit(X_train, y_train)
            pred = pipe.predict(X_test)
            score = accuracy_score(y_test, pred)
            print("pipeline avec deux learners :", score)
    """

    def __init__(self, model=None, method=None, **kwargs):
        """
        @param  model   instance d'un learner
        @param  method  méthode à appeler pour transformer les features (voir-ci-dessous)
        @param  kwargs  paramètres

        Options pour le paramètres *method* :

        * ``'predict'``
        * ``'predict_proba'``
        * ``'decision_function'``
        * une fonction

        Si *method is None*, la fonction essaye dans l'ordre
        ``predict_proba`` puis ``predict``.
        """
        super().__init__(**kwargs)
        self.model = model
        if model is None:
            raise ValueError("value cannot be None")
        if method is None:
            for name in {'predict_proba', 'predict', 'transform'}:
                if hasattr(model.__class__, name):
                    method = name
            if method is None:
                raise ValueError(
                    "Unable to guess a default method for '{0}'".format(repr(model)))
        self.method = method
        self._set_method(method)

    def _set_method(self, method):
        """
        Définit la méthode à utiliser pour transmettre
        les variables à la suite du pipeline.
        """
        if isinstance(method, str):
            if method == 'predict':
                self.method_ = self.model.predict
            elif method == 'predict_proba':
                self.method_ = self.model.predict_proba
            elif method == 'decision_function':
                self.method_ = self.model.decision_function
            elif method == 'transform':
                self.method_ = self.model.transform
            else:
                raise ValueError("Unexpected method '{0}'".format(method))
        elif callable(method):
            self.method_ = method
        else:
            raise TypeError(
                "Unable to find the transform method, method={0}".format(method))

    def fit(self, X, y=None, **kwargs):
        """
        Apprends un modèle.

        @param      X               features
        @param      y               cibles
        @param      kwargs          paramètres additionnels
        @return                     self, lui-même
        """
        self.model.fit(X, y=y, **kwargs)
        return self

    def transform(self, X):
        """
        Prédit, souvent cela se résume à appeler la mathode *decision_function*.

        @param      X   features
        @return         prédictions
        """
        res = self.method_(X)
        if len(res.shape) == 1:
            res = res[:, numpy.newaxis]
        return res

    ##############
    # cloning API
    ##############

    def get_params(self, deep=True):
        """
        Returns the parameters mandatory to clone the class.

        @param      deep        unused here
        @return                 dict
        """
        res = self.P.to_dict()
        res['model'] = self.model
        res['method'] = self.method
        if deep:
            par = self.model.get_params(deep)
            for k, v in par.items():
                res["model__" + k] = v
        return res

    def set_params(self, **values):
        """
        Set parameters.

        @param      values      parameters
        """
        if 'model' in values:
            self.model = values['model']
            del values['model']
        elif not hasattr(self, 'model') or self.model is None:
            raise KeyError("Missing key '{0}' in [{1}]".format(
                'model', ', '.join(sorted(values))))
        if 'method' in values:
            self._set_method(values['method'])
            del values['method']
        for k in values:
            if not k.startswith('model__'):
                raise ValueError(
                    "Parameter '{0}' must start with 'model__'.".format(k))
        d = len('model__')
        pars = {k[d:]: v for k, v in values.items()}
        self.model.set_params(**pars)
        if 'method' in values:
            self.method = values['method']
            self._set_method(values['method'])

    #################
    # common methods
    #################

    def __repr__(self):
        """
        usual
        """
        rp = repr(self.model)
        rps = repr(self.P)
        res = "{0}(model={1}, method={2}, {3})".format(
            self.__class__.__name__, rp, self.method, rps)
        return "\n".join(textwrap.wrap(res, subsequent_indent="    "))
