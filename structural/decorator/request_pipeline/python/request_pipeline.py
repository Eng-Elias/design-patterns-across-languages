import abc
from dataclasses import dataclass
from typing import Dict, Any


@dataclass
class Request:
    """Represents an HTTP request with headers and body."""
    headers: Dict[str, str]
    body: Dict[str, Any]
    path: str
    method: str


@dataclass
class Response:
    """Represents an HTTP response with status code and body."""
    status_code: int
    body: Dict[str, Any]


class RequestHandler(abc.ABC):
    """The common interface for all request handlers."""

    @abc.abstractmethod
    def handle(self, request: Request) -> Response:
        """Process the request and return a response."""
        pass


class BaseHandler(RequestHandler):
    """The base request handler that processes the core request."""

    def handle(self, request: Request) -> Response:
        """Process the request and return a response."""
        # In a real application, this would contain the core business logic
        return Response(
            status_code=200,
            body={"message": "Request processed successfully"}
        )


class Middleware(RequestHandler, abc.ABC):
    """Abstract base class for all middleware decorators."""

    def __init__(self, handler: RequestHandler):
        self._handler = handler

    @abc.abstractmethod
    def handle(self, request: Request) -> Response:
        pass


class AuthMiddleware(Middleware):
    """Middleware that checks authentication before processing the request."""

    def handle(self, request: Request) -> Response:
        # Check if the request has a valid auth token
        if "Authorization" not in request.headers:
            return Response(
                status_code=401,
                body={"error": "Unauthorized"}
            )
        
        # If authenticated, pass to the next handler
        return self._handler.handle(request)


class LoggingMiddleware(Middleware):
    """Middleware that logs request details before processing."""

    def handle(self, request: Request) -> Response:
        # Log request details
        print(f"Request received: {request.method} {request.path}")
        print(f"Headers: {request.headers}")
        print(f"Body: {request.body}")

        # Process the request
        response = self._handler.handle(request)

        # Log response
        print(f"Response: {response.status_code}")
        print(f"Body: {response.body}")

        return response


class RateLimitMiddleware(Middleware):
    """Middleware that implements rate limiting."""

    def __init__(self, handler: RequestHandler, max_requests: int = 100):
        super().__init__(handler)
        self._max_requests = max_requests
        self._request_count = 0

    def handle(self, request: Request) -> Response:
        self._request_count += 1
        
        if self._request_count > self._max_requests:
            return Response(
                status_code=429,
                body={"error": "Too many requests"}
            )
        
        return self._handler.handle(request) 