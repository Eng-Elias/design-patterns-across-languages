import { IObserver } from "./observer";
import { IEvent, Event, EventType } from "./event";

/**
 * The Subject (Observable) interface.
 */
export interface IEventSource {
  attach(observer: IObserver): void;
  detach(observer: IObserver): void;
  notify(event: IEvent): void;
  generateEvent(eventType: EventType, data: Record<string, unknown>): void;
}

/**
 * The Concrete Subject (Observable) class.
 */
export class EventSource implements IEventSource {
  private observers: Set<IObserver> = new Set(); // Use Set for efficient add/remove/check

  public attach(observer: IObserver): void {
    if (!this.observers.has(observer)) {
      console.log(`EventSource: Attaching ${observer.constructor.name}`);
      this.observers.add(observer);
    } else {
      console.log(
        `EventSource: ${observer.constructor.name} already attached.`
      );
    }
  }

  public detach(observer: IObserver): void {
    if (this.observers.has(observer)) {
      console.log(`EventSource: Detaching ${observer.constructor.name}`);
      this.observers.delete(observer);
    } else {
      console.log(
        `EventSource: ${observer.constructor.name} not found for detachment.`
      );
    }
  }

  public notify(event: IEvent): void {
    console.log(
      `EventSource: Notifying observers about event: ${event.getType()}`
    );
    this.observers.forEach((observer) => {
      observer.update(event);
    });
  }

  /**
   * Simulate generating an event and notify observers.
   */
  public generateEvent(
    eventType: EventType,
    data: Record<string, unknown>
  ): void {
    console.log(`\nEventSource: Generating event '${eventType}'...`);
    const event = new Event(eventType, data);
    this.notify(event);
  }

  // Helper for testing
  public getObserverCount(): number {
    return this.observers.size;
  }
}
