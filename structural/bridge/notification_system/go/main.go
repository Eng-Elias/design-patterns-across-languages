package main

import (
	ns "bridge_pattern_notification_system_go/notification_system"
	"fmt"
)

func main() {
	fmt.Println("Demonstrating the Bridge pattern for sending notifications in Go.")

	// Create the different implementation (sender) objects
	// We need pointers if we want to use methods associated with the struct type directly,
	// but for interfaces, the value type works fine if the methods are defined on the value receiver.
	// Here, methods are on pointer receivers, so we use pointers or let Go handle the address implicitly.
	// It's generally safer/clearer to work with pointers when methods modify state or when interfaces are involved.
	var emailSender ns.MessageSender = &ns.EmailSender{}
	var smsSender ns.MessageSender = &ns.SmsSender{}
	var pushSender ns.MessageSender = &ns.PushNotificationSender{}

	fmt.Println("\n--- Sending Info Notifications ---")
	infoEmail := ns.NewInfoNotification(emailSender)
	infoEmail.Send("System update scheduled for tonight.")

	infoSms := ns.NewInfoNotification(smsSender)
	infoSms.Send("Short notice: Maintenance window extended by 1 hour.")

	fmt.Println("\n--- Sending Warning Notifications ---")
	warningPush := ns.NewWarningNotification(pushSender)
	warningPush.Send("Disk space reaching 85% capacity on server SRV-01.")

	warningEmail := ns.NewWarningNotification(emailSender)
	warningEmail.Send("API response times are slightly elevated.")

	fmt.Println("\n--- Sending Urgent Notifications ---")
	urgentSms := ns.NewUrgentNotification(smsSender)
	urgentSms.Send("Critical service XYZ is down! Immediate attention required.")

	urgentPush := ns.NewUrgentNotification(pushSender) // Can reuse sender
	urgentPush.Send("Security Breach Detected on User Account 'admin'! Action needed.")

	// Notice how we can easily combine any Notification type (Abstraction)
	// with any Sender type (Implementation) without needing specific types
	// like UrgentSmsNotification or InfoEmailNotification.
}
