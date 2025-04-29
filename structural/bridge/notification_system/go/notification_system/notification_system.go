package notification_system

import "fmt"

// --- Implementation Interface ---
type MessageSender interface {
	// Defines the interface for implementation classes (message sending mechanisms).
	SendMessage(subject string, body string)
}

// --- Concrete Implementations ---

type EmailSender struct{}

// SendMessage provides the concrete implementation for sending messages via Email.
func (s *EmailSender) SendMessage(subject string, body string) {
	fmt.Println("--- Sending Email ---")
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Body: %s\n", body)
	fmt.Println("---------------------")
}

type SmsSender struct{}

// SendMessage provides the concrete implementation for sending messages via SMS.
func (s *SmsSender) SendMessage(subject string, body string) {
	// SMS usually doesn't have a subject, so we combine them
	fmt.Println("--- Sending SMS ---")
	fmt.Println("To: [PhoneNumber]") // Placeholder
	fmt.Printf("Message: %s - %s\n", subject, body)
	fmt.Println("-----------------")
}

type PushNotificationSender struct{}

// SendMessage provides the concrete implementation for sending messages via Push Notification.
func (s *PushNotificationSender) SendMessage(subject string, body string) {
	fmt.Println("--- Sending Push Notification ---")
	fmt.Printf("Title: %s\n", subject)
	fmt.Printf("Body: %s\n", body)
	fmt.Println("-------------------------------")
}

// --- Abstraction Interface ---
type Notification interface {
	// Defines the abstraction's interface.
	Send(message string)
}

// --- Base Abstraction Struct (Optional but helps avoid repetition) ---
// Go doesn't have abstract classes, but embedding can achieve similar code reuse.
type BaseNotification struct {
	sender MessageSender // The Bridge: holds the implementation interface
}

// --- Refined Abstractions ---

type InfoNotification struct {
	BaseNotification // Embed the base to get the sender field
}

// NewInfoNotification creates a new InfoNotification.
func NewInfoNotification(sender MessageSender) *InfoNotification {
	return &InfoNotification{BaseNotification{sender: sender}}
}

// Send implements the Notification interface for Info messages.
func (n *InfoNotification) Send(message string) {
	subject := "Info"
	body := fmt.Sprintf("[INFO] %s", message)
	fmt.Printf("Preparing '%s' notification.\n", subject)
	n.sender.SendMessage(subject, body)
}

type WarningNotification struct {
	BaseNotification
}

// NewWarningNotification creates a new WarningNotification.
func NewWarningNotification(sender MessageSender) *WarningNotification {
	return &WarningNotification{BaseNotification{sender: sender}}
}

// Send implements the Notification interface for Warning messages.
func (n *WarningNotification) Send(message string) {
	subject := "Warning"
	body := fmt.Sprintf("[WARNING] %s", message)
	fmt.Printf("Preparing '%s' notification.\n", subject)
	n.sender.SendMessage(subject, body)
}

type UrgentNotification struct {
	BaseNotification
}

// NewUrgentNotification creates a new UrgentNotification.
func NewUrgentNotification(sender MessageSender) *UrgentNotification {
	return &UrgentNotification{BaseNotification{sender: sender}}
}

// Send implements the Notification interface for Urgent messages.
func (n *UrgentNotification) Send(message string) {
	subject := "** URGENT **"
	body := fmt.Sprintf("[URGENT ACTION REQUIRED] %s", message)
	fmt.Printf("Preparing '%s' notification.\n", subject)
	n.sender.SendMessage(subject, body)
}
