import unittest
from unittest.mock import MagicMock, call
from event_source import EventSource
from observer import LoggerObserver, NotifierObserver, Observer
from event import Event, EventType
import time

class TestEventMonitoring(unittest.TestCase):

    def setUp(self):
        self.event_source = EventSource()
        # Use real observers for integration testing, mocks for isolation
        self.logger = LoggerObserver("TestLogger")
        self.notifier = NotifierObserver("TestNotifier")
        self.mock_observer1 = MagicMock(spec=Observer)
        self.mock_observer2 = MagicMock(spec=Observer)

    def test_attach_observer(self):
        self.assertEqual(self.event_source.get_observer_count(), 0)
        self.event_source.attach(self.mock_observer1)
        self.assertEqual(self.event_source.get_observer_count(), 1)
        self.event_source.attach(self.mock_observer2)
        self.assertEqual(self.event_source.get_observer_count(), 2)

    def test_attach_duplicate_observer(self):
        self.event_source.attach(self.mock_observer1)
        self.assertEqual(self.event_source.get_observer_count(), 1)
        self.event_source.attach(self.mock_observer1) # Attach again
        self.assertEqual(self.event_source.get_observer_count(), 1)

    def test_detach_observer(self):
        self.event_source.attach(self.mock_observer1)
        self.event_source.attach(self.mock_observer2)
        self.assertEqual(self.event_source.get_observer_count(), 2)
        self.event_source.detach(self.mock_observer1)
        self.assertEqual(self.event_source.get_observer_count(), 1)
        self.event_source.detach(self.mock_observer2)
        self.assertEqual(self.event_source.get_observer_count(), 0)

    def test_detach_non_existent_observer(self):
        self.event_source.attach(self.mock_observer1)
        self.assertEqual(self.event_source.get_observer_count(), 1)
        self.event_source.detach(self.mock_observer2) # Try detaching unattached observer
        self.assertEqual(self.event_source.get_observer_count(), 1)

    def test_notify_observers(self):
        self.event_source.attach(self.mock_observer1)
        self.event_source.attach(self.mock_observer2)

        event_data = {"key": "value"}
        event = Event(EventType.LOG_INFO, event_data)

        # Directly call _notify for testing notification mechanism
        self.event_source._notify(event)

        self.mock_observer1.update.assert_called_once_with(event)
        self.mock_observer2.update.assert_called_once_with(event)

    def test_generate_event_notifies(self):
        self.event_source.attach(self.mock_observer1)

        # Mock the _notify method to check if it's called by generate_event
        self.event_source._notify = MagicMock()

        self.event_source.generate_event(EventType.LOG_INFO, {"data": "payload"})

        # Check that _notify was called (implicitly checks event creation too)
        self.event_source._notify.assert_called_once()
        # Check the argument passed to _notify was an Event object
        call_args = self.event_source._notify.call_args
        self.assertIsInstance(call_args[0][0], Event)
        self.assertEqual(call_args[0][0].event_type, EventType.LOG_INFO)
        self.assertEqual(call_args[0][0].data, {"data": "payload"})

    def test_logger_observer_update(self):
        self.event_source.attach(self.logger)
        self.assertEqual(len(self.logger.get_logs()), 0)
        self.event_source.generate_event(EventType.LOG_INFO, {"msg": "Info message"})
        self.assertEqual(len(self.logger.get_logs()), 1)
        self.assertIn("LOG_INFO: {'msg': 'Info message'}", self.logger.get_logs()[0])

        self.event_source.generate_event(EventType.LOG_WARN, {"msg": "Warning message"})
        self.assertEqual(len(self.logger.get_logs()), 2)
        self.assertIn("LOG_WARN: {'msg': 'Warning message'}", self.logger.get_logs()[1])

    def test_notifier_observer_update(self):
        self.event_source.attach(self.notifier)
        self.assertEqual(len(self.notifier.get_notifications()), 0)

        # INFO should not trigger notification
        self.event_source.generate_event(EventType.LOG_INFO, {"status": "OK"})
        self.assertEqual(len(self.notifier.get_notifications()), 0)

        # ERROR should trigger notification
        self.event_source.generate_event(EventType.LOG_ERROR, {"code": 503})
        self.assertEqual(len(self.notifier.get_notifications()), 1)
        self.assertIn("Notify: EventType.LOG_ERROR - {'code': 503}", self.notifier.get_notifications()[0])

        # CRITICAL should trigger notification
        self.event_source.generate_event(EventType.LOG_CRITICAL, {"system": "DB"})
        self.assertEqual(len(self.notifier.get_notifications()), 2)
        self.assertIn("Notify: EventType.LOG_CRITICAL - {'system': 'DB'}", self.notifier.get_notifications()[1])


if __name__ == '__main__':
    unittest.main()