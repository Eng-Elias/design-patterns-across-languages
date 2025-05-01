from http_request_middleware import (
    Handler, LoggingMiddleware, AuthenticationMiddleware,
    AuthorizationMiddleware, FinalHandler, make_request
)

def setup_middleware_chain() -> Handler:
    """Creates and links the middleware handlers."""
    logger = LoggingMiddleware()
    authenticator = AuthenticationMiddleware()
    authorizer = AuthorizationMiddleware()
    final_handler = FinalHandler()

    # Build the chain: Log -> Auth -> Authorize -> Final
    logger.set_next(authenticator).set_next(authorizer).set_next(final_handler)

    return logger # Return the first handler in the chain

if __name__ == "__main__":
    middleware_chain = setup_middleware_chain()

    print("--- Simulating a successful request ---")
    request1 = make_request("/admin/dashboard", token="valid_token", requires_role="admin")
    response1 = middleware_chain.handle(request1)
    print(f"Response 1: {response1}\n")

    print("--- Simulating a request with invalid authentication ---")
    request2 = make_request("/user/profile", token="invalid_token")
    response2 = middleware_chain.handle(request2)
    print(f"Response 2: {response2}\n")

    print("--- Simulating a request with insufficient authorization ---")
    request3 = make_request("/admin/settings", token="valid_token", requires_role="superadmin")
    response3 = middleware_chain.handle(request3)
    print(f"Response 3: {response3}\n")

    print("--- Simulating a request requiring no specific role ---")
    request4 = make_request("/public/info", token="valid_token")
    response4 = middleware_chain.handle(request4)
    print(f"Response 4: {response4}\n")
