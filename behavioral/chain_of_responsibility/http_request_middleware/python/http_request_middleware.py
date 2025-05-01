from abc import ABC, abstractmethod
from typing import Any, Dict, Optional

class Handler(ABC):
    """
    Abstract Handler class defines the interface for handling requests and
    linking handlers.
    """
    _next_handler: Optional['Handler'] = None

    def set_next(self, handler: 'Handler') -> 'Handler':
        """Sets the next handler in the chain."""
        self._next_handler = handler
        return handler # Allows chaining

    @abstractmethod
    def handle(self, request: Dict[str, Any]) -> Optional[Dict[str, Any]]:
        """
        Handles the request. Implementation should call self._next_handler.handle(request)
        if processing should continue, otherwise return a result or None to stop.
        """
        # Default implementation if the chain ends without a specific final handler returning a value
        if self._next_handler:
             return self._next_handler.handle(request)
        return request # Or None, depending on desired end-of-chain default


class LoggingMiddleware(Handler):
    """Logs the incoming request."""
    def handle(self, request: Dict[str, Any]) -> Optional[Dict[str, Any]]:
        print(f"[Log] Received request: {request.get('path', 'N/A')}")
        # Explicitly call the next handler if it exists
        if self._next_handler:
            return self._next_handler.handle(request)
        # If no next handler, return the request or a default response
        return request

class AuthenticationMiddleware(Handler):
    """Checks for authentication credentials."""
    def handle(self, request: Dict[str, Any]) -> Optional[Dict[str, Any]]:
        print("[Auth] Checking authentication...")
        headers = request.get('headers', {})
        if headers.get('Authorization') == 'valid_token':
            print("[Auth] Authentication successful.")
            request['user'] = {'id': 123, 'role': 'admin'}
            # Explicitly call the next handler if it exists
            if self._next_handler:
                return self._next_handler.handle(request)
            # If no next handler, return the (modified) request
            return request
        else:
            print("[Auth] Authentication failed. Aborting request.")
            # Stop the chain by returning the error response directly
            return {'status_code': 401, 'body': 'Unauthorized'}

class AuthorizationMiddleware(Handler):
    """Checks if the authenticated user has the required permissions."""
    def handle(self, request: Dict[str, Any]) -> Optional[Dict[str, Any]]:
        print("[ACL] Checking authorization...")
        user = request.get('user')
        required_role = request.get('required_role')

        if not user:
            print("[ACL] No user found in request (likely due to failed auth). Aborting.")
            # Stop the chain
            return {'status_code': 401, 'body': 'Authentication Required'}

        proceed = False
        if required_role and user.get('role') == required_role:
            print(f"[ACL] Authorization successful (Role: {user.get('role')}).")
            proceed = True
        elif not required_role:
             print("[ACL] No specific role required. Proceeding.")
             proceed = True

        if proceed:
             # Explicitly call the next handler if it exists
            if self._next_handler:
                 return self._next_handler.handle(request)
            # If no next handler, return the request
            return request
        else:
            # This case handles: required_role exists, but user role doesn't match
            print(f"[ACL] Authorization failed. User role '{user.get('role')}' does not match required role '{required_role}'. Aborting.")
            # Stop the chain
            return {'status_code': 403, 'body': 'Forbidden'}


class FinalHandler(Handler):
    """The final processing step if the request passes all middleware."""
    def handle(self, request: Dict[str, Any]) -> Optional[Dict[str, Any]]:
        print("[Final] Processing request...")
        # Simulate actual request processing
        user_id = request.get('user', {}).get('id', 'anonymous')
        body = f"Successfully processed request for path: {request.get('path', 'N/A')} for user {user_id}"
        print("[Final] Request processed.")
        # This is the end of the chain, return the final response
        return {'status_code': 200, 'body': body}

# Helper function (no changes needed)
def make_request(path: str, token: Optional[str] = None, requires_role: Optional[str] = None) -> Dict[str, Any]:
    request: Dict[str, Any] = {'path': path, 'headers': {}}
    if token:
        request['headers']['Authorization'] = token
    if requires_role:
        request['required_role'] = requires_role
    return request
