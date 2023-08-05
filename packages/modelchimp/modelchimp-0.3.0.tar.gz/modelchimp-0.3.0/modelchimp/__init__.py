from __future__ import print_function
import requests
import json
import future
import six
from six.moves import cPickle as pickle
import zlib
import sys
import inspect
import os
from . import metrics

# Optional imports
library_base_class = []
try:
    from sklearn.base import BaseEstimator
    library_base_class.append(BaseEstimator)
except ImportError:
    pass

try:
    from keras.engine.training import Model
    from . import mckeras
    library_base_class.append(Model)
except ImportError as e:
    print(e)
    pass


__version__ = '0.3.0'

class PlatformLibraryType(object):
    SKLEARN = '1'
    KERAS = '2'
    CHOICES = (
        (SKLEARN, 'Sklearn'),
        (KERAS, 'Keras')
    )


class Track:
    URL = "https://www.modelchimp.com/"

    def __init__(self, key):
        self._session = requests.Session()
        self._features = []
        self._model = None
        self._evaluation = []
        self._deep_learning_parameters = []
        self._algorithm = ""
        self._http_headers = {}
        self._project_id = ""
        self._platform_library = ""
        self.__authenticate(key)
        self.__parse_ml_code()

    def add_evaluation(self, eval_name, eval_value):
        if not isinstance(eval_name, str):
            raise Exception("Evaluation name should be a string")

        if eval_name == "":
            raise Exception("Evaluation name should not be empty")

        if not ( isinstance(eval_value, int) or isinstance(eval_value, float) ):
            raise Exception("Evaluation value should be a number")

        self._evaluation.append({'key': eval_name, 'value': eval_value})

    def show(self):
        '''
        Prints the details that is going to be synced to the cloud
        '''
        print("\n")
        print("---Model Parameter List---")
        for obj in self._model:
            model_text = "%s : %s" % (obj['key'], obj['value'])
            print(model_text)

        print("\n")
        print("---Evaluation List---")
        for obj in self._evaluation:
            evaluation_text = "%s : %s" % ( obj['key'],
                                            obj['value'])
            print(evaluation_text)

    def save(self, name=None):
        '''
        Save the details to the ModelChimp cloud
        '''

        ml_model_url = Track.URL + 'api/ml-model/'
        result = {
            "name": name,
            "features": json.dumps(self._features),
            "model_parameters": json.dumps(self._model),
            "evaluation_parameters": json.dumps(self._evaluation),
            "deep_learning_parameters": json.dumps(self._deep_learning_parameters),
            "project": self._project_id,
            "algorithm": self._algorithm,
            "platform": "PYTHON",
            "platform_library": self._platform_library
        }

        # Check if its python script
        frame = inspect.stack()[1]
        module = inspect.getmodule(frame[0])
        try:
            filename = module.__file__
            with open(filename, 'rb') as f:
                save_request = self._session.post(ml_model_url, data=result,
                files={"code_file": f}, headers=self._http_headers)
        except AttributeError:
            save_request = self._session.post(ml_model_url, data=result,
                            headers=self._http_headers)

        # if filename.split("/")[-1] != 'interactiveshell.py':
        #     with open(filename, 'rb') as f:
        #         save_request = self._session.post(ml_model_url, data=result,
        #         files={"code_file": f}, headers=self._http_headers)
        # else:
        #     save_request = self._session.post(ml_model_url, data=result,
        #                     headers=self._http_headers)

        if save_request.status_code == 201:
            print("The data was successfully saved")
        else:
            print("The data could not be saved")

    def __authenticate(self, key):
        authentication_url = Track.URL + 'api/decode-key/'
        auth_data = {"key": key}

        request_auth = self._session.post(authentication_url, data=auth_data)
        # Check if the request got authenticated
        if request_auth.status_code != 200:
            raise requests.exceptions.RequestException(
                "Oops...something has gone wrong.")

        # Get the authenticated token and assign it to the header
        token = json.loads(request_auth.text)['token']
        self._http_headers = {'Authorization': 'Token ' + token}
        self._project_id = json.loads(request_auth.text)['project_id']

    def __parse_ml_code(self):
        ml_code_dict = inspect.stack()[2][0].f_locals
        keys = list(ml_code_dict)

        for key in keys:
            value = ml_code_dict[key]
            try:
                if isinstance(value, BaseEstimator):
                    self._platform_library = PlatformLibraryType.SKLEARN
                    self._model = self.__dict_to_kv(value.get_params())
                    self._algorithm = value.__class__.__name__

                if isinstance(value, Model):
                    self._platform_library = PlatformLibraryType.KERAS
                    keras_model_params = value.__dict__
                    self._algorithm = value.__class__.__name__
                    self._deep_learning_parameters = mckeras._get_layer_info(value.layers) if self._algorithm == 'Sequential' else []

                if isinstance(value, str) or isinstance(value, unicode):
                    if mckeras.get_compile_params(value):
                        keras_compile_params = mckeras.get_compile_params(value)

                    if mckeras.get_fit_params(value):
                        keras_fit_params = mckeras.get_fit_params(value)

            except NameError as e:
                pass

        if self._platform_library == PlatformLibraryType.KERAS:
            keras_model_params.update(keras_compile_params)
            keras_model_params.update(keras_fit_params)
            keras_model_params = mckeras._get_relevant_params(keras_model_params, ml_code_dict)
            self._model = self.__dict_to_kv(keras_model_params)


        if self._platform_library == "":
            raise Exception("There are no sklearn or keras packages in the scope")


    def __dict_to_kv(self, dict_val):
        result = [{'key': k, 'value': v} for k, v in dict_val.items()]
        result.sort(key=lambda e: e['key'])

        return result
