import { IEvent, EventType } from "./event";

/**
 * Observer interface defining the update method.
 */
export interface IObserver {
  update(event: IEvent): void;
}

/**
 * Concrete observer that logs events.
 */
export class LoggerObserver implements IObserver {
  private name: string;
  private logMessages: string[] = []; // Store logs for testing

  constructor(name: string) {
    this.name = name;
  }

  update(event: IEvent): void {
    const dataStr = JSON.stringify(event.getData());
    const logMsg = `${event.getType()}: ${dataStr} at ${event
      .getTimestamp()
      .toISOString()}`;
    console.log(`LoggerObserver: Received event - ${logMsg}`); // Added log for demo
    this.logMessages.push(logMsg);
  }

  // Helper for testing
  public getLogs(): string[] {
    return this.logMessages;
  }
}

/**
 * Concrete observer that sends notifications (simulated).
 */
export class NotifierObserver implements IObserver {
  private name: string;
  private notifications: string[] = []; // Store notifications for testing

  constructor(name: string) {
    this.name = name;
  }

  update(event: IEvent): void {
    const eventType = event.getType();
    // Compare using EventType enum members
    if (
      eventType === EventType.LogError ||
      eventType === EventType.LogCritical
    ) {
      const notificationMsg = `Notification: Critical event - Type: ${eventType}, Data: ${JSON.stringify(
        event.getData()
      )}, Time: ${event.getTimestamp().toISOString()}`;
      console.log(`NotifierObserver: ${notificationMsg}`); // Added log for demo
      this.notifications.push(notificationMsg);
    }
  }

  // Helper for testing
  public getNotifications(): string[] {
    return this.notifications;
  }
}
