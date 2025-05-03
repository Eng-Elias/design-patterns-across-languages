from dataclasses import dataclass
from datetime import datetime
from enum import Enum


class EventType(Enum):
    LOG_INFO = "INFO"
    LOG_WARN = "WARN"
    LOG_ERROR = "ERROR"
    LOG_CRITICAL = "CRITICAL"

@dataclass
class Event:
    """Data class to represent an event."""
    event_type: EventType
    data: dict
    timestamp: datetime = datetime.now()

    def __str__(self):
        return f"[{self.timestamp.strftime('%Y-%m-%d %H:%M:%S')}] {self.event_type}: {self.data}"
