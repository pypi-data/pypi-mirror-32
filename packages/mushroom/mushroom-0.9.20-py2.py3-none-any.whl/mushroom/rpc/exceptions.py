class RpcError(RuntimeError):
    '''
    Base class for all exceptions raised from by the `Engine`.
    '''
    def __init__(self, message=''):
        self.message = message
        super(RpcError, self).__init__(message)


class MethodNotFound(RpcError):
    '''
    This error is raised when a method for a `Request` or `Notification`
    message is not found. This can either happen when a connected client
    tries to call a server method or the server tries to call a method
    on the client side.
    '''
    def __init__(self, method_name):
        self.method_name = method_name
        super(MethodNotFound, self).__init__(method_name)


class RequestTimeout(RpcError):
    '''
    This error is raised when a `Request` message is not answered within
    a specified timeout value. By default the value is set to infinite
    can be set do a different value when making the request.
    '''
    pass


class RequestException(RpcError):
    '''
    This exception is raised when a `Request` message is answered with
    an `Error` message.
    '''

    def __init__(self, data):
        super(RequestException, self).__init__()
        self.data = data


class RequestValidationError(RequestException):
    # FIXME
    pass
