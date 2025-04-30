from request_pipeline import (
    Request,
    BaseHandler,
    AuthMiddleware,
    LoggingMiddleware,
    RateLimitMiddleware
)


def print_response(response):
    """Helper function to print response details."""
    print(f"Status Code: {response.status_code}")
    print(f"Body: {response.body}\n")


def main():
    """Demonstrates the Request Processing Pipeline with different middleware combinations."""

    # Create a sample request
    request = Request(
        headers={"Content-Type": "application/json"},
        body={"data": "test"},
        path="/api/test",
        method="POST"
    )

    print("--- Simple Request Processing ---")
    # Process request with base handler only
    base_handler = BaseHandler()
    response = base_handler.handle(request)
    print_response(response)

    print("--- Request with Auth Middleware ---")
    # Process request with auth middleware
    auth_handler = AuthMiddleware(base_handler)
    response = auth_handler.handle(request)
    print_response(response)

    # Add authorization header and try again
    request.headers["Authorization"] = "Bearer valid-token"
    response = auth_handler.handle(request)
    print_response(response)

    print("--- Request with Rate Limiting ---")
    # Process request with rate limiting
    rate_limit_handler = RateLimitMiddleware(base_handler, max_requests=2)
    for i in range(3):  # Try 3 requests
        print(f"Request {i + 1}:")
        response = rate_limit_handler.handle(request)
        print_response(response)

    print("--- Request with Multiple Middleware ---")
    # Create a chain of middleware: Auth -> RateLimit -> Logging -> BaseHandler
    handler = AuthMiddleware(
        RateLimitMiddleware(
            LoggingMiddleware(
                BaseHandler()
            ),
            max_requests=2
        )
    )

    # Process request through the middleware chain
    response = handler.handle(request)
    print_response(response)


if __name__ == "__main__":
    main() 