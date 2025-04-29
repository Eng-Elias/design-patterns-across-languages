import {
  EmailSender,
  SmsSender,
  PushNotificationSender,
  InfoNotification,
  WarningNotification,
  UrgentNotification,
} from "./notification_system";

function main(): void {
  /**
   * Demonstrates the Bridge pattern for sending notifications.
   */

  // Create the different implementation (sender) objects
  const emailSender = new EmailSender();
  const smsSender = new SmsSender();
  const pushSender = new PushNotificationSender();

  console.log("--- Sending Info Notifications ---");
  const infoEmail = new InfoNotification(emailSender);
  infoEmail.send("System update scheduled for tonight.");

  const infoSms = new InfoNotification(smsSender);
  infoSms.send("Short notice: Maintenance window extended by 1 hour.");

  console.log("\n--- Sending Warning Notifications ---");
  const warningPush = new WarningNotification(pushSender);
  warningPush.send("Disk space reaching 85% capacity on server SRV-01.");

  const warningEmail = new WarningNotification(emailSender);
  warningEmail.send("API response times are slightly elevated.");

  console.log("\n--- Sending Urgent Notifications ---");
  const urgentSms = new UrgentNotification(smsSender);
  urgentSms.send("Critical service XYZ is down! Immediate attention required.");

  const urgentEmailPush = new UrgentNotification(pushSender); // Can reuse sender
  urgentEmailPush.send(
    "Security Breach Detected on User Account 'admin'! Action needed."
  );

  // Notice how we can easily combine any Notification type (Abstraction)
  // with any Sender type (Implementation) without needing specific classes
  // like UrgentSmsNotification or InfoEmailNotification.
}

main();
