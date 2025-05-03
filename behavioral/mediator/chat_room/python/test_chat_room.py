import unittest
from unittest.mock import MagicMock, call
# Assuming chat_room.py is in the same directory or accessible via PYTHONPATH
from chat_room import ChatRoom, ChatUser, User, ChatMediator

class TestChatRoomMediator(unittest.TestCase):

    def setUp(self):
        """Set up for test methods."""
        self.mediator = ChatRoom()
        # We'll use MagicMock for users in some tests to isolate mediator logic
        # In others, we'll use real ChatUser instances to test integration

    def test_add_user(self):
        """Test adding users to the chat room."""
        mock_user_alice = MagicMock(spec=User)
        mock_user_alice.name = "Alice"
        mock_user_bob = MagicMock(spec=User)
        mock_user_bob.name = "Bob"

        self.mediator.add_user(mock_user_alice)
        self.mediator.add_user(mock_user_bob)

        # Check if users are in the internal dictionary (accessing internal state for testing)
        self.assertIn("Alice", self.mediator._users)
        self.assertIn("Bob", self.mediator._users)
        self.assertEqual(self.mediator._users["Alice"], mock_user_alice)

        # Try adding the same user again (should not add duplicate)
        initial_user_count = len(self.mediator._users)
        self.mediator.add_user(mock_user_alice)
        self.assertEqual(len(self.mediator._users), initial_user_count) # Count shouldn't change

    def test_remove_user(self):
        """Test removing users from the chat room."""
        mock_user_alice = MagicMock(spec=User)
        mock_user_alice.name = "Alice"
        self.mediator.add_user(mock_user_alice)
        self.assertIn("Alice", self.mediator._users)

        self.mediator.remove_user(mock_user_alice)
        self.assertNotIn("Alice", self.mediator._users)

        # Try removing a non-existent user (should not raise error)
        mock_user_charlie = MagicMock(spec=User)
        mock_user_charlie.name = "Charlie"
        try:
            self.mediator.remove_user(mock_user_charlie)
        except Exception as e:
            self.fail(f"Removing non-existent user raised an exception: {e}")


    def test_send_message(self):
        """Test sending a message distributes it correctly."""
        # Use real users this time to check receive calls
        user_alice = ChatUser(self.mediator, "Alice")
        user_bob = ChatUser(self.mediator, "Bob")
        user_charlie = ChatUser(self.mediator, "Charlie")

        # Mock the receive method on users to check if they are called
        user_alice.receive = MagicMock()
        user_bob.receive = MagicMock()
        user_charlie.receive = MagicMock()

        self.mediator.add_user(user_alice)
        self.mediator.add_user(user_bob)
        self.mediator.add_user(user_charlie)

        message = "Hello from Alice!"
        # Alice sends a message
        user_alice.send(message) # This internally calls mediator.send_message

        # Check if Bob and Charlie received the message, but Alice did not
        user_alice.receive.assert_not_called()
        user_bob.receive.assert_called_once_with(message, "Alice")
        user_charlie.receive.assert_called_once_with(message, "Alice")

        # Bob sends a message
        user_bob.receive.reset_mock() # Reset mock for the next assertion
        user_charlie.receive.reset_mock()

        message_bob = "Hi Alice!"
        user_bob.send(message_bob)

        user_alice.receive.assert_called_once_with(message_bob, "Bob")
        user_bob.receive.assert_not_called()
        user_charlie.receive.assert_called_once_with(message_bob, "Bob")


    def test_user_integration(self):
        """Test the interaction between ChatUser and ChatRoom."""
        # This test ensures the User's send method correctly calls the mediator
        mock_mediator = MagicMock(spec=ChatMediator)
        user_dave = ChatUser(mock_mediator, "Dave")
        
        message = "Test message"
        user_dave.send(message)

        # Verify that the user's send method called the mediator's send_message
        mock_mediator.send_message.assert_called_once_with(message, user_dave)


if __name__ == '__main__':
    unittest.main()