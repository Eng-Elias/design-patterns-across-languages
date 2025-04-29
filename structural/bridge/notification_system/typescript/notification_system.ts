// --- Implementation Interface ---
export interface MessageSender {
  /**
   * Defines the interface for implementation classes (message sending mechanisms).
   */
  sendMessage(subject: string, body: string): void;
}

// --- Concrete Implementations ---
export class EmailSender implements MessageSender {
  /**
   * Concrete implementation for sending messages via Email.
   */
  public sendMessage(subject: string, body: string): void {
    console.log(`--- Sending Email ---`);
    console.log(`Subject: ${subject}`);
    console.log(`Body: ${body}`);
    console.log(`---------------------`);
  }
}

export class SmsSender implements MessageSender {
  /**
   * Concrete implementation for sending messages via SMS.
   */
  public sendMessage(subject: string, body: string): void {
    // SMS usually doesn't have a subject, so we combine them
    console.log(`--- Sending SMS ---`);
    console.log(`To: [PhoneNumber]`); // Placeholder
    console.log(`Message: ${subject} - ${body}`);
    console.log(`-----------------`);
  }
}

export class PushNotificationSender implements MessageSender {
  /**
   * Concrete implementation for sending messages via Push Notification.
   */
  public sendMessage(subject: string, body: string): void {
    console.log(`--- Sending Push Notification ---`);
    console.log(`Title: ${subject}`);
    console.log(`Body: ${body}`);
    console.log(`-------------------------------`);
  }
}

// --- Abstraction ---
export abstract class Notification {
  /**
   * Defines the abstraction's interface and maintains a reference to an implementation object.
   */
  protected sender: MessageSender; // Bridge

  constructor(sender: MessageSender) {
    this.sender = sender;
  }

  public abstract send(message: string): void;
}

// --- Refined Abstractions ---
export class InfoNotification extends Notification {
  /**
   * A specific type of notification (refined abstraction).
   */
  public send(message: string): void {
    const subject = "Info";
    const body = `[INFO] ${message}`;
    console.log(`Preparing '${subject}' notification.`);
    this.sender.sendMessage(subject, body);
  }
}

export class WarningNotification extends Notification {
  /**
   * Another specific type of notification.
   */
  public send(message: string): void {
    const subject = "Warning";
    const body = `[WARNING] ${message}`;
    console.log(`Preparing '${subject}' notification.`);
    this.sender.sendMessage(subject, body);
  }
}

export class UrgentNotification extends Notification {
  /**
   * An urgent notification type.
   */
  public send(message: string): void {
    const subject = "** URGENT **";
    const body = `[URGENT ACTION REQUIRED] ${message}`;
    console.log(`Preparing '${subject}' notification.`);
    this.sender.sendMessage(subject, body);
  }
}
