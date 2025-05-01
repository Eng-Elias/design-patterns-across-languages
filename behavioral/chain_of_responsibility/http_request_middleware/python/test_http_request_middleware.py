import unittest
from unittest.mock import patch, call # Import call for checking print arguments

from http_request_middleware import (
    LoggingMiddleware, AuthenticationMiddleware, AuthorizationMiddleware,
    FinalHandler, make_request
)

class TestHttpRequestMiddleware(unittest.TestCase):

    # No setUp/tearDown needed for stdout redirection

    def setUp(self):
        # Set up the chain fresh for each test
        self.logger = LoggingMiddleware()
        self.authenticator = AuthenticationMiddleware()
        self.authorizer = AuthorizationMiddleware()
        self.final_handler = FinalHandler()
        # Properly chain the handlers
        self.logger.set_next(self.authenticator)
        self.authenticator.set_next(self.authorizer)
        self.authorizer.set_next(self.final_handler)
        # Start the chain with the logger
        self.chain = self.logger

    @patch('builtins.print') # Patch print within the test method's scope
    def test_successful_request_admin(self, mock_print):
        """Tests a request that should pass all middleware."""
        request = make_request("/admin", token="valid_token", requires_role="admin")
        response = self.chain.handle(request)

        # 1. Check the final response
        self.assertIsNotNone(response)
        self.assertEqual(response['status_code'], 200)
        self.assertIn("Successfully processed request", response['body'])
        self.assertIn("user 123", response['body']) # Check user ID from auth

        # 2. Check the sequence of print calls
        expected_calls = [
            call("[Log] Received request: /admin"),
            call("[Auth] Checking authentication..."),
            call("[Auth] Authentication successful."),
            call("[ACL] Checking authorization..."),
            call("[ACL] Authorization successful (Role: admin)."),
            call("[Final] Processing request..."),
            call("[Final] Request processed.")
        ]
        mock_print.assert_has_calls(expected_calls, any_order=False)
        # Ensure no unexpected prints happened (optional, but good practice)
        self.assertEqual(mock_print.call_count, len(expected_calls))

    @patch('builtins.print')
    def test_authentication_failure(self, mock_print):
        """Tests a request that should fail at the authentication step."""
        request = make_request("/secure", token="invalid_token")
        response = self.chain.handle(request)

        # 1. Check the final response (should be from Auth middleware)
        self.assertIsNotNone(response)
        self.assertEqual(response['status_code'], 401)
        self.assertEqual(response['body'], 'Unauthorized')

        # 2. Check the print calls that occurred
        mock_print.assert_any_call("[Log] Received request: /secure")
        mock_print.assert_any_call("[Auth] Checking authentication...")
        mock_print.assert_any_call("[Auth] Authentication failed. Aborting request.")

        # 3. Check that subsequent middleware prints DID NOT occur
        acl_check_call = call("[ACL] Checking authorization...")
        final_process_call = call("[Final] Processing request...")

        found_acl_check = any(c == acl_check_call for c in mock_print.call_args_list)
        found_final_process = any(c == final_process_call for c in mock_print.call_args_list)

        self.assertFalse(found_acl_check, "ACL middleware print should not have occurred")
        self.assertFalse(found_final_process, "Final handler print should not have occurred")
        # Verify only the expected prints happened
        self.assertEqual(mock_print.call_count, 3)


    @patch('builtins.print')
    def test_authorization_failure(self, mock_print):
        """Tests a request that should pass auth but fail authorization."""
        request = make_request("/admin/power", token="valid_token", requires_role="superadmin")
        response = self.chain.handle(request)

        # 1. Check the final response (should be from ACL middleware)
        self.assertIsNotNone(response)
        self.assertEqual(response['status_code'], 403)
        self.assertEqual(response['body'], 'Forbidden')

        # 2. Check the print calls that occurred
        mock_print.assert_any_call("[Log] Received request: /admin/power")
        mock_print.assert_any_call("[Auth] Checking authentication...")
        mock_print.assert_any_call("[Auth] Authentication successful.")
        mock_print.assert_any_call("[ACL] Checking authorization...")
        mock_print.assert_any_call("[ACL] Authorization failed. User role 'admin' does not match required role 'superadmin'. Aborting.")

        # 3. Check that subsequent middleware prints DID NOT occur
        final_process_call = call("[Final] Processing request...")
        found_final_process = any(c == final_process_call for c in mock_print.call_args_list)
        self.assertFalse(found_final_process, "Final handler print should not have occurred")
        # Verify only the expected prints happened
        self.assertEqual(mock_print.call_count, 5)


    @patch('builtins.print')
    def test_successful_request_no_role_required(self, mock_print):
        """Tests a request that passes all middleware without needing a specific role."""
        request = make_request("/public", token="valid_token")
        response = self.chain.handle(request)

        # 1. Check the final response
        self.assertIsNotNone(response)
        self.assertEqual(response['status_code'], 200)
        self.assertIn("Successfully processed request", response['body'])
        self.assertIn("user 123", response['body']) # Check user ID from auth

        # 2. Check the sequence of print calls
        expected_calls = [
            call("[Log] Received request: /public"),
            call("[Auth] Checking authentication..."),
            call("[Auth] Authentication successful."),
            call("[ACL] Checking authorization..."),
            call("[ACL] No specific role required. Proceeding."),
            call("[Final] Processing request..."),
            call("[Final] Request processed.")
        ]
        mock_print.assert_has_calls(expected_calls, any_order=False)
        self.assertEqual(mock_print.call_count, len(expected_calls))

# Remove the __main__ block if running via 'python -m unittest'
if __name__ == '__main__':
   unittest.main()
