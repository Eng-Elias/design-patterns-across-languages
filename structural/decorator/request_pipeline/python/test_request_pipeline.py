import unittest
from request_pipeline import (
    Request,
    BaseHandler,
    AuthMiddleware,
    RateLimitMiddleware
)


class TestRequestPipeline(unittest.TestCase):

    def setUp(self):
        self.request = Request(
            headers={"Content-Type": "application/json"},
            body={"data": "test"},
            path="/api/test",
            method="POST"
        )

    def test_base_handler(self):
        """Test the base handler without any middleware."""
        handler = BaseHandler()
        response = handler.handle(self.request)
        
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.body, {"message": "Request processed successfully"})

    def test_auth_middleware_unauthorized(self):
        """Test AuthMiddleware with missing authorization header."""
        handler = AuthMiddleware(BaseHandler())
        response = handler.handle(self.request)
        
        self.assertEqual(response.status_code, 401)
        self.assertEqual(response.body, {"error": "Unauthorized"})

    def test_auth_middleware_authorized(self):
        """Test AuthMiddleware with valid authorization header."""
        self.request.headers["Authorization"] = "Bearer valid-token"
        handler = AuthMiddleware(BaseHandler())
        response = handler.handle(self.request)
        
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.body, {"message": "Request processed successfully"})

    def test_rate_limit_middleware(self):
        """Test RateLimitMiddleware with multiple requests."""
        handler = RateLimitMiddleware(BaseHandler(), max_requests=2)
        
        # First request should succeed
        response = handler.handle(self.request)
        self.assertEqual(response.status_code, 200)
        
        # Second request should succeed
        response = handler.handle(self.request)
        self.assertEqual(response.status_code, 200)
        
        # Third request should fail
        response = handler.handle(self.request)
        self.assertEqual(response.status_code, 429)
        self.assertEqual(response.body, {"error": "Too many requests"})

    def test_middleware_chain(self):
        """Test multiple middleware components chained together."""
        # Create a chain: Auth -> RateLimit -> BaseHandler
        handler = AuthMiddleware(
            RateLimitMiddleware(
                BaseHandler(),
                max_requests=2
            )
        )
        
        # Test with valid auth
        self.request.headers["Authorization"] = "Bearer valid-token"
        response = handler.handle(self.request)
        self.assertEqual(response.status_code, 200)
        
        # Test with invalid auth
        del self.request.headers["Authorization"]
        response = handler.handle(self.request)
        self.assertEqual(response.status_code, 401)


if __name__ == '__main__':
    unittest.main() 