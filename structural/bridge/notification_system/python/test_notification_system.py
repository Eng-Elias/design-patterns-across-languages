import unittest
from unittest.mock import MagicMock
from notification_system import (
    MessageSender, # Import the ABC for spec
    InfoNotification,
    WarningNotification,
    UrgentNotification
)

class TestNotificationSystemBridge(unittest.TestCase):

    def test_info_notification_sends_correctly(self):
        """Verify InfoNotification formats and sends messages via its sender."""
        # Arrange
        mock_sender = MagicMock(spec=MessageSender)
        info_notification = InfoNotification(mock_sender)
        message = "Test info message."
        expected_subject = "Info"
        expected_body = f"[INFO] {message}"

        # Act
        info_notification.send(message)

        # Assert
        mock_sender.send_message.assert_called_once_with(expected_subject, expected_body)

    def test_warning_notification_sends_correctly(self):
        """Verify WarningNotification formats and sends messages via its sender."""
        # Arrange
        mock_sender = MagicMock(spec=MessageSender)
        warning_notification = WarningNotification(mock_sender)
        message = "Test warning message."
        expected_subject = "Warning"
        expected_body = f"[WARNING] {message}"

        # Act
        warning_notification.send(message)

        # Assert
        mock_sender.send_message.assert_called_once_with(expected_subject, expected_body)

    def test_urgent_notification_sends_correctly(self):
        """Verify UrgentNotification formats and sends messages via its sender."""
        # Arrange
        mock_sender = MagicMock(spec=MessageSender)
        urgent_notification = UrgentNotification(mock_sender)
        message = "Test urgent message."
        expected_subject = "** URGENT **"
        expected_body = f"[URGENT ACTION REQUIRED] {message}"

        # Act
        urgent_notification.send(message)

        # Assert
        mock_sender.send_message.assert_called_once_with(expected_subject, expected_body)

    def test_different_notifications_use_same_sender(self):
        """Verify different notification types can use the same sender instance."""
        # Arrange
        mock_sender = MagicMock(spec=MessageSender)
        info_notification = InfoNotification(mock_sender)
        urgent_notification = UrgentNotification(mock_sender)

        # Act
        info_notification.send("Info 1")
        urgent_notification.send("Urgent 1")

        # Assert
        self.assertEqual(mock_sender.send_message.call_count, 2)
        # Check calls were made with expected arguments (order doesn't matter here)
        mock_sender.send_message.assert_any_call("Info", "[INFO] Info 1")
        mock_sender.send_message.assert_any_call("** URGENT **", "[URGENT ACTION REQUIRED] Urgent 1")


if __name__ == '__main__':
    unittest.main()
