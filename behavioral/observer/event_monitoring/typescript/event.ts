export enum EventType {
  LogInfo = "LOG_INFO",
  LogWarn = "LOG_WARN",
  LogError = "LOG_ERROR",
  LogCritical = "LOG_CRITICAL",
}

export interface IEvent {
  getType(): EventType;
  getData(): Record<string, unknown>;
  getTimestamp(): Date;
  toString(): string;
}

export class Event implements IEvent {
  private type: EventType;
  private data: Record<string, unknown>;
  private timestamp: Date;

  constructor(eventType: EventType, data: Record<string, unknown>) {
    this.type = eventType;
    this.data = data;
    this.timestamp = new Date();
  }

  public getType(): EventType {
    return this.type;
  }

  public getData(): Record<string, unknown> {
    return this.data;
  }

  public getTimestamp(): Date {
    return this.timestamp;
  }

  public toString(): string {
    const timeStr = this.timestamp
      .toISOString()
      .replace("T", " ")
      .substring(0, 19);
    return `[${timeStr}] ${this.type}: ${JSON.stringify(this.data)}`;
  }
}
