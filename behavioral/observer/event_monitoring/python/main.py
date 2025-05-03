import time
from event import EventType
from event_source import EventSource
from observer import LoggerObserver, NotifierObserver

def main():
    print("--- Python Observer Pattern Demo ---")

    event_source = EventSource()
    # Use names for observers for clarity
    logger = LoggerObserver(name="FileLogger")
    notifier = NotifierObserver(name="EmailNotifier")

    print("\nAttaching observers...")
    event_source.attach(logger)
    event_source.attach(notifier)

    print("\nGenerating specific events...")
    event_source.generate_event(EventType.LOG_INFO, {"msg": "User logged in"})
    time.sleep(0.05) # Simulate delay
    event_source.generate_event(EventType.LOG_WARN, {"msg": "Disk space low"})
    time.sleep(0.05)
    event_source.generate_event(EventType.LOG_ERROR, {"code": 500, "error": "Database connection failed"})
    time.sleep(0.05)
    event_source.generate_event(EventType.LOG_CRITICAL, {"code": 999, "error": "System meltdown imminent"})
    time.sleep(0.05)

    # Detach the NotifierObserver (consistent with Go example)
    print("\nDetaching EmailNotifier...")
    event_source.detach(notifier)

    print("\nGenerating another event...")
    event_source.generate_event(EventType.LOG_INFO, {"msg": "User logged out"})
    time.sleep(0.05)

    print("\n--- Demo Finished ---")

if __name__ == "__main__":
    main()