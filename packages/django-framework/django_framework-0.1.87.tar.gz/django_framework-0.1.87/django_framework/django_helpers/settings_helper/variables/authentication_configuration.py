
from base_variable import BaseVariable

class AuthenticationConfiguration(BaseVariable):
    
    @property
    def REQUIRE_AUTHENTICATION(self):
        var_name = 'require_authentication'
        return self._get_set_variables(var_name = var_name)
    
    
    def _get_require_authentication(self):
        try:
            return self._REQUIRE_AUTHENTICATION, None, None
        except Exception as e:
            print(e)
            print('Defaulting back to True, all things require authentication as needed')
            return True, None, None
        
        
    def _check_require_authentication(self, value):
        self.assert_boolean(value = value)
        return True

    @property
    def AUTHENTICATION_SHOULD_PASSTHROUGH(self):
        var_name = 'authentication_should_passthrough'
        return self._get_set_variables(var_name = var_name)
    
    
    
    @property
    def IS_AUTHENTICATION_SERVER(self):
        var_name = 'is_authentication_server'
        return self._get_set_variables(var_name = var_name)
    
    
    def _get_is_authentication_server(self):
        return self._IS_AUTHENTICATION_SERVER, None, None
    
    def _check_is_authentication_server(self, value):
        self.assert_boolean(value = value)
        return True

    
    def _get_authentication_should_passthrough(self):
        return self._AUTHENTICATION_SHOULD_PASSTHROUGH, None, None
    
    def _check_authentication_should_passthrough(self, value):
        self.assert_boolean(value = value)
        
        if value == True:
            print('-----------------------------------------------------------')
            print('AuthenticationConfiguration:')
            print('WARNING: AUTHENTICATION_SHOULD_PASSTHROUGH is set to True. NO user checks on token for authentication will be done.')
            print('-----------------------------------------------------------')
        return True

