import unittest
from unittest.mock import MagicMock
from logging_adapter import ThirdPartyLogger, LoggerAdapter, ApplicationService, Logger

class TestLoggingAdapter(unittest.TestCase):

    def test_adapter_calls_adaptee_correctly(self):
        """Verify that adapter methods call the correct adaptee method with the right arguments."""
        # Arrange
        mock_adaptee = MagicMock(spec=ThirdPartyLogger)
        adapter = LoggerAdapter(mock_adaptee)

        # Act
        adapter.log_info("Info message")
        adapter.log_warning("Warning message")
        adapter.log_error("Error message")

        # Assert
        mock_adaptee.record.assert_any_call("info", "Info message")
        mock_adaptee.record.assert_any_call("warning", "Warning message")
        mock_adaptee.record.assert_any_call("error", "Error message")
        self.assertEqual(mock_adaptee.record.call_count, 3)

    def test_application_service_uses_logger(self):
        """Verify that the ApplicationService uses the provided logger interface."""
        # Arrange
        mock_logger = MagicMock(spec=Logger) # Mock the Target interface
        app_service = ApplicationService(mock_logger)

        # Act - Successful operation
        app_service.perform_operation("Valid Data")

        # Assert - Successful operation logging
        mock_logger.log_info.assert_any_call("Starting operation with data: Valid Data")
        mock_logger.log_info.assert_any_call("Operation completed successfully.")
        self.assertEqual(mock_logger.log_info.call_count, 2)
        mock_logger.log_warning.assert_not_called()
        mock_logger.log_error.assert_not_called()

        # Reset mock for next part
        mock_logger.reset_mock()

        # Act - Operation with warning
        app_service.perform_operation("shrt")

        # Assert - Warning logging
        mock_logger.log_info.assert_any_call("Starting operation with data: shrt")
        mock_logger.log_warning.assert_called_once_with("Data 'shrt' is quite short.")
        mock_logger.log_info.assert_any_call("Operation completed successfully.")
        self.assertEqual(mock_logger.log_info.call_count, 2) # Start and End
        mock_logger.log_error.assert_not_called()

        # Reset mock for next part
        mock_logger.reset_mock()

        # Act - Operation with error
        app_service.perform_operation("")

        # Assert - Error logging
        mock_logger.log_info.assert_called_once_with("Starting operation with data: ")
        mock_logger.log_error.assert_called_once_with("Operation failed: Data cannot be empty")
        mock_logger.log_warning.assert_not_called()
        # Check log_info was called only once (start), not for success
        self.assertEqual(mock_logger.log_info.call_count, 1)


if __name__ == '__main__':
    unittest.main()
