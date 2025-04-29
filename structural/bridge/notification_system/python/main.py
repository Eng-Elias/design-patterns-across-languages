from notification_system import (
    EmailSender, SmsSender, PushNotificationSender,
    InfoNotification, WarningNotification, UrgentNotification
)

def main():
    """Demonstrates the Bridge pattern for sending notifications."""

    # Create the different implementation (sender) objects
    email_sender = EmailSender()
    sms_sender = SmsSender()
    push_sender = PushNotificationSender()

    print("--- Sending Info Notifications ---")
    info_email = InfoNotification(email_sender)
    info_email.send("System update scheduled for tonight.")

    info_sms = InfoNotification(sms_sender)
    info_sms.send("Short notice: Maintenance window extended by 1 hour.")

    print("\n--- Sending Warning Notifications ---")
    warning_push = WarningNotification(push_sender)
    warning_push.send("Disk space reaching 85% capacity on server SRV-01.")

    warning_email = WarningNotification(email_sender)
    warning_email.send("API response times are slightly elevated.")

    print("\n--- Sending Urgent Notifications ---")
    urgent_sms = UrgentNotification(sms_sender)
    urgent_sms.send("Critical service XYZ is down! Immediate attention required.")

    urgent_email_push = UrgentNotification(push_sender) # Can reuse sender
    urgent_email_push.send("Security Breach Detected on User Account 'admin'! Action needed.")

    # Notice how we can easily combine any Notification type (Abstraction)
    # with any Sender type (Implementation) without needing specific classes
    # like UrgentSmsNotification or InfoEmailNotification.

if __name__ == "__main__":
    main()
