import { EventSource } from "./event_source";
import { LoggerObserver, NotifierObserver, IObserver } from "./observer";
import { Event, IEvent, EventType } from "./event";

// Mock Observer for testing purposes
const mockObserver1: IObserver = {
  update: jest.fn(),
};

const mockObserver2: IObserver = {
  update: jest.fn(),
};

describe("Event Monitoring - Observer Pattern (TypeScript)", () => {
  let eventSource: EventSource;
  let logger: LoggerObserver;
  let notifier: NotifierObserver;

  beforeEach(() => {
    // Reset mocks before each test
    (mockObserver1.update as jest.Mock).mockClear();
    (mockObserver2.update as jest.Mock).mockClear();

    eventSource = new EventSource();
    logger = new LoggerObserver("TestLogger");
    notifier = new NotifierObserver("TestNotifier");
  });

  test("should attach observers correctly", () => {
    expect(eventSource.getObserverCount()).toBe(0);
    eventSource.attach(mockObserver1);
    expect(eventSource.getObserverCount()).toBe(1);
    eventSource.attach(mockObserver2);
    expect(eventSource.getObserverCount()).toBe(2);
  });

  test("should not attach the same observer twice", () => {
    eventSource.attach(mockObserver1);
    expect(eventSource.getObserverCount()).toBe(1);
    eventSource.attach(mockObserver1); // Try attaching again
    expect(eventSource.getObserverCount()).toBe(1);
  });

  test("should detach observers correctly", () => {
    eventSource.attach(mockObserver1);
    eventSource.attach(mockObserver2);
    expect(eventSource.getObserverCount()).toBe(2);
    eventSource.detach(mockObserver1);
    expect(eventSource.getObserverCount()).toBe(1);
    eventSource.detach(mockObserver2);
    expect(eventSource.getObserverCount()).toBe(0);
  });

  test("should handle detaching non-existent observers gracefully", () => {
    eventSource.attach(mockObserver1);
    expect(eventSource.getObserverCount()).toBe(1);
    eventSource.detach(mockObserver2); // Try detaching unattached observer
    expect(eventSource.getObserverCount()).toBe(1);
  });

  test("should notify all attached observers", () => {
    eventSource.attach(mockObserver1);
    eventSource.attach(mockObserver2);

    const testEvent = new Event(EventType.LogInfo, { data: "test data" });
    eventSource.notify(testEvent);

    expect(mockObserver1.update).toHaveBeenCalledTimes(1);
    expect(mockObserver1.update).toHaveBeenCalledWith(testEvent);
    expect(mockObserver2.update).toHaveBeenCalledTimes(1);
    expect(mockObserver2.update).toHaveBeenCalledWith(testEvent);
  });

  test("generateEvent should create an event and notify observers", (done: jest.DoneCallback) => {
    // Mock the notify method to check if it's called by generateEvent
    const notifySpy = jest.spyOn(eventSource, "notify");
    eventSource.attach(mockObserver1); // Need at least one observer for notify to be called

    eventSource.generateEvent(EventType.LogWarn, { payload: 123 });

    // Wait for the setTimeout in generateEvent to complete
    setTimeout(() => {
      expect(notifySpy).toHaveBeenCalledTimes(1);
      const notifiedEvent = notifySpy.mock.calls[0][0] as IEvent;
      expect(notifiedEvent).toBeInstanceOf(Event);
      expect(notifiedEvent.getType()).toBe(EventType.LogWarn);
      expect(notifiedEvent.getData()).toEqual({ payload: 123 });
      notifySpy.mockRestore(); // Clean up the spy
      done(); // Signal Jest that the async test is complete
    }, 200); // Wait slightly longer than generateEvent's timeout
  });

  test("LoggerObserver should log all events", () => {
    eventSource.attach(logger);
    const event1 = new Event(EventType.LogInfo, { msg: "Info 1" });
    const event2 = new Event(EventType.LogWarn, { msg: "Warn 1" });

    eventSource.notify(event1);
    eventSource.notify(event2);

    const logs = logger.getLogs();
    expect(logs).toHaveLength(2);
    expect(logs[0]).toContain(`LOG_INFO: {"msg":"Info 1"}`);
    expect(logs[1]).toContain(`LOG_WARN: {"msg":"Warn 1"}`);
  });

  test("NotifierObserver should only notify on ERROR or CRITICAL events", () => {
    const notifier = new NotifierObserver("TestNotifier"); // Added name
    const eventInfo = new Event(EventType.LogInfo, { status: "OK" });
    const eventError = new Event(EventType.LogError, { code: 503 });
    const eventCritical = new Event(EventType.LogCritical, { system: "DB" });

    // Attach is not needed for direct update calls
    // eventSource.attach(notifier);

    // Simulate notifications by calling update directly
    notifier.update(eventInfo); // Should not add notification
    notifier.update(eventError); // Should add notification
    notifier.update(eventCritical); // Should add notification

    const notifications = notifier.getNotifications();
    expect(notifications).toHaveLength(2);
    // Check essential parts, ignoring timestamp
    expect(notifications[0]).toContain(`Type: ${EventType.LogError}`);
    expect(notifications[0]).toContain(
      `Data: ${JSON.stringify({ code: 503 })}`
    );
    expect(notifications[1]).toContain(`Type: ${EventType.LogCritical}`);
    expect(notifications[1]).toContain(
      `Data: ${JSON.stringify({ system: "DB" })}`
    );
  });
});
