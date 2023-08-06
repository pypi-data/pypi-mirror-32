import json

import arrow
import copy
from django_framework.django_helpers.api_helpers import BaseAPI

from django_framework.django_helpers.exception_helpers.common_exceptions import PermissionError, PermissionAdminError
from django_framework.django_helpers.model_helpers import get_model
from django_framework.django_helpers.serializer_helpers import get_serializer
from django_framework.django_helpers.manager_helpers import get_manager
from django_framework.django_helpers.meta_helpers import get_meta

from django_framework.django_helpers.api_helpers import register_api

from kafka.protocol.api import Response

from django.conf import settings

try:
    CHECK_MODEL_DATA_WITH_PERMISSIONS = settings.CHECK_MODEL_DATA_WITH_PERMISSIONS
except Exception as e:
    CHECK_MODEL_DATA_WITH_PERMISSIONS = True
    print('-----------------------------------------------------------')
    print('django_framework.base_framework.model_api.widget.model_api:')
    print('Error:', e)
    print('WARNING: CHECK_MODEL_DATA_WITH_PERMISSIONS was not set properly, defaulting to True (strongest permissions).')
    print('-----------------------------------------------------------')

    
try:  # 12-17-2017 we currently do not use this variable  becuase it's really really unsafe
    GET_MODEL_WITH_PERMISSIONS = settings.CHECK_MODEL_DATA_WITH_PERMISSIONS
except Exception as e:
    GET_MODEL_WITH_PERMISSIONS = True
    print('-----------------------------------------------------------')
    print('django_framework.base_framework.model_api.widget.model_api:')
    print('Error:', e)
    print('WARNING: GET_MODEL_WITH_PERMISSIONS was not set properly, defaulting to True (strongest permissions).')
    print('-----------------------------------------------------------')



class ModelAPI(BaseAPI):
    
    def __init__(self, request, model_name, model_pk = None, model_uuid = None, admin = False, **kwargs):
        
        self.request_requires_authentication = False # this is used so that we can initiate other variables before checking authentication
        kwargs['request_requires_authentication'] = self.request_requires_authentication # we will authenticate afterwards, since it is dependent on the specific Model.
        
        super(ModelAPI, self).__init__(request=request, admin=admin, **kwargs)
        
        self.model_name = self.clean_model_name(model_name)
        self.model_pk = model_pk
        self.model_uuid = model_uuid
        
        self.Model = get_model(model_name = self.model_name)
        self.Serializer = get_serializer(serializer_name = self.model_name, version = self.version)
        
        self.Manager = get_manager(manager_name = self.model_name)
        self.set_model_meta()

        self.set_requires_validation()
        self.validate_permissions()
    
    def set_requires_validation(self):
        '''We find that this one really doesnt matter anymore?'''
        if self.version == 'admin':
            self.request_requires_authentication = self.Meta.ADMIN_REQUIRE_AUTHENTICATION
        else:
            self.request_requires_authentication = self.Meta.DEFAULT_REQUIRE_AUTHENTICATION

    
    def set_model_meta(self):
        '''This is mainly used for testing as we do not expect people to swap!'''
        ## we think the logic is wrong.  TODO make sure that it is right Jan 18, 2018
        
        user_set_meta = self.request.META.get('HTTP_MODEL_META') or self.request.META.get('MODEL_META')
        if user_set_meta:
            self.Meta = get_meta(meta_name = user_set_meta)
        else:
            self.Meta = get_meta(meta_name = self.model_name)
        

    def validate_permissions(self):
        '''Restrict basic actions based on if logged in and user priveldges!'''

        # we check first if it is "required" to be validated!
        # as it is used to set the later filters...
        # this means that if you are authenticated, it sitll wont filter by username later...
        
        if self.Meta.DEFAULT_ALLOWED_ACTIONS_UNAUTHENTICATED is not None and self.request_method in self.Meta.DEFAULT_ALLOWED_ACTIONS_UNAUTHENTICATED:
            self.request_requires_authentication = False

        try:
            self.is_user_authenticated()  # since we no longer check during initiation, we check explicitely
            return

        except PermissionAdminError as e:
            # if the error is because the user hit an admin endpoint and are not admin,
            # then we always error out. as all admin endpoints SHOULD require authentication?
            raise 
        
        except PermissionError as e:
            # the user is not authenticated, but maybe the endpoint does nto require it!
            if self.request_requires_authentication == False:
                return

        raise PermissionError('This request method is not allowed when not authenticated!')
        
    def run(self):
        if self.request.method == 'GET':
            objs = self.get()

        elif self.request.method == 'POST':
            objs = self.post()
            
        elif self.request.method == 'PUT':
            objs = self.put()
            
        elif self.request.method == 'DELETE':
            objs = self.delete()
                
        self.data = objs
        return objs

    def get(self):
        self._check_request_method_is_allowed()
        objs = self.get_model_with_permissions()
        self._need_to_cache_results = True
        
#         print(self.Manager.get_related_model(models = objs, relationship_name = 'profile'))
        
        return objs
    
    def post(self):
        self._check_request_method_is_allowed()
        
        # check to make sure that the input data is allowed (prevent creation of row to another user)
        self._check_model_data_with_permissions(input_data = self.query_data)
        
        objs = self.Manager.create(Serializer = self.Serializer, data = self.query_data)
        self._set_clear_cache_level(level = None)
        return objs
    
    def put(self):
        self._check_request_method_is_allowed(model_information_required = True)
        objs = self.get_model_with_permissions(exactly_one_result = True)
        
        # check to make sure that input data is allwoed (prevent re-assigning of row to another user)
        self._check_model_data_with_permissions(input_data = self.query_data)
        
        objs = self.Manager.update(Serializer = self.Serializer, data = self.query_data, model = objs[0])
        
        self._set_clear_cache_level(level = None)
        
        return objs

    def delete(self):
        self._check_request_method_is_allowed(model_information_required = True)
        objs = self.get_model_with_permissions(exactly_one_result = True)
        objs = self.Manager.delete(model = objs[0])
        
        self._set_clear_cache_level(level = None)
        
        return objs

    def _check_request_method_is_allowed(self, model_information_required = False):
        '''Check to make sure that everything is allowed
        1.  Check if the Meta allows the method to be run
        2.  Check to make sure , if needed, that a model is specified (mainly for PUT and DELETE)
        '''
        self.Meta.allowed_method(method = self.request.method, version = self.version)
        
        if model_information_required == True:
            if self.model_pk is None and self.model_uuid is None:
                raise ValueError('You cannot change a model without specifying a specific one!')
    
    
    
    def _check_model_data_with_permissions(self, input_data):
        '''We are checking that when a user updates a "parent" field, they also own the parent!
        This prevents POSTing new rows to a random user,
        This prevents updating existing rows to a new user.
        '''
        if self.version == 'admin' or CHECK_MODEL_DATA_WITH_PERMISSIONS == False:
            return # we do no checks onw hat is allowed or not.
        
        # get the parent model name if it exists!
        try:
            parent_model_name = self.Model.PARENT_MODEL_NAME # note that it can be none
            ParentManager = get_manager(manager_name = parent_model_name)
        except Exception as e: #both of them might not exist...
            parent_model_name = None
            ParentManager = None # it doesnt exist.

        if ParentManager == None:
            return  # we could not find a parent. we therefore allow you through! probably not as safe???
        
        # we now get to checking the user input...
        parent_id =  input_data.get(parent_model_name + '_id')
        if parent_id == None:
            return # we are not editing anything that will affect our parent id.

        # get the parent model to make sure it is woned by the user!
        query_params = ParentManager._update_query(query_params={'filter' : {'id' : parent_id}}, param_name = 'relationship', update_dict = {'profile' : [self.user.profile_id]})
        response = ParentManager.get_by_query(query_params = query_params)
        
        if response.count() != 1:
            raise ValueError('You have attempted to create or update a row so that it will no longer be owned by you.  This is not allowed.')
            
        
        
        
    
    def get_model_with_permissions(self, exactly_one_result = False):
        '''Requesting the appropriate model from DB.  We add in a filter to limit by the user if they are not admin'''
        query_params = copy.copy(self.query_params) # when updating query_params, do not want to mess up the original!

        if self.version == "admin": 
            # Admin's are allowed to manipulate everyone!
            pass
        elif self.request_requires_authentication == False:  
            # this means that the method and endpoint doesnt need to be validated (ie no permission restrictions)
            pass
        else:
            # update the query_parmas to include a relationship requirement to this profile!
            query_params = self.Manager._update_query(query_params=query_params, param_name = 'relationship', update_dict = {'profile' : [self.user.profile_id]})
    
        if self.model_pk:
            query_params = self.Manager._update_query(query_params=query_params, param_name = 'filter', update_dict = {'pk' : self.model_pk})
        
        if self.model_uuid:
            query_params = self.Manager._update_query(query_params=query_params, param_name = 'filter', update_dict = {'uuid' : self.model_uuid})
        

        objs = self.Manager.get_by_query(query_params = query_params, query_set = None)
        
        # do quick check if results require at least one.
        if exactly_one_result == True:
            if len(objs) == 0:
                raise ValueError('Returned 0 results when expecting strictly 1 result.')
            elif len(objs) >1:
                raise ValueError('Returned 2 or more results, when expecting strictly 1 result.')
        
        return objs
    
    
    def get_response(self):
        '''Override BaseAPI version of get_response to get autoformatting of data'''
        return self.format_data(data = self.data)
    
    def format_data(self, data):
        '''Serializes data and paginate data. Includes meta data as well.'''
        meta_data = self._set_base_meta_data(len_data = len(data), response_type = self.model_name)

        if self.request_method == 'DELETE':
            pass
        else:
            data = self.serialize_data(data = self._paginate(data))
        return dict(data=data, meta_data = meta_data)
    
register_api(ModelAPI)
