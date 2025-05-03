package chat_room

import (
	"fmt"
	"log"
	"sync"
)

// ChatMediator defines the interface for the chat room mediator.
type ChatMediator interface {
	AddUser(user User)
	RemoveUser(user User)
	SendMessage(message string, sender User)
}

// User defines the interface for chat room participants (colleagues).
type User interface {
	Send(message string)
	Receive(message string, senderName string)
	GetName() string
	SetMediator(mediator ChatMediator)
}

// --- Concrete Mediator: ChatRoom ---

// ChatRoom implements the ChatMediator interface.
type ChatRoom struct {
	users map[string]User
	mutex sync.RWMutex // Use RWMutex for better performance if reads are frequent
}

// NewChatRoom creates a new ChatRoom instance.
func NewChatRoom() *ChatRoom {
	// Initialize logger to behave like Python's print (no timestamp/prefix)
	log.SetFlags(0)
	return &ChatRoom{
		users: make(map[string]User),
	}
}

// AddUser adds a user to the chat room.
func (cr *ChatRoom) AddUser(user User) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	userName := user.GetName()
	if _, exists := cr.users[userName]; !exists {
		// Add user if not already present
		cr.users[userName] = user
		user.SetMediator(cr) // Set the mediator for the user
		log.Printf("--- %s added to the chat room. ---", userName) // Match Python output
	} else {
		// Log if the user is already present
		log.Printf("--- %s is already in the chat room. Cannot add again. ---", userName) // Match Python output
	}
}

// RemoveUser removes a user from the chat room.
func (cr *ChatRoom) RemoveUser(user User) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	userName := user.GetName()
	if _, exists := cr.users[userName]; exists {
		delete(cr.users, userName)
		log.Printf("--- %s removed from the chat room. ---", userName) // Match Python output
	} else {
		// Log if the user is not found
		log.Printf("--- %s is not in the chat room. Cannot remove. ---", userName) // Match Python output
	}
}

// SendMessage broadcasts a message to all users except the sender.
func (cr *ChatRoom) SendMessage(message string, sender User) {
	cr.mutex.RLock() // Use RLock for reading the user map
	defer cr.mutex.RUnlock()

	senderName := sender.GetName()
	// Log sender and message here, similar to Python
	log.Printf("--- %s sends message: '%s' ---", senderName, message)

	for name, user := range cr.users {
		// Ensure we don't send the message back to the sender
		if name != senderName {
			user.Receive(message, senderName) // Send synchronously
		}
	}
}

// --- Concrete Colleague: ChatUser ---

// ChatUser implements the User interface.
type ChatUser struct {
	mediator ChatMediator
	name     string
}

// NewChatUser creates a new ChatUser instance.
func NewChatUser(name string) *ChatUser {
	return &ChatUser{
		name: name,
	}
}

// SetMediator sets the mediator for the user. Required by ChatRoom.AddUser.
func (cu *ChatUser) SetMediator(mediator ChatMediator) {
	cu.mediator = mediator
}

// GetName returns the user's name.
func (cu *ChatUser) GetName() string {
	return cu.name
}

// Send sends a message via the mediator.
func (cu *ChatUser) Send(message string) {
	if cu.mediator == nil {
		// Use fmt.Printf here as log might not be configured in all contexts
		// Match python's implicit behavior (no output if mediator not set) or add specific error handling if desired
		// For now, let's keep it simple and just return if no mediator
		// fmt.Printf("'%s' cannot send message: Mediator not set.\n", cu.name) // Optional: uncomment for debugging
		return
	}
	// The mediator now logs the sending action in SendMessage
	cu.mediator.SendMessage(message, cu)
}

// Receive receives a message from the mediator.
func (u *ChatUser) Receive(message string, senderName string) {
	u.logReceive(message, senderName)
}

// logReceive handles the actual logging for Receive.
func (u *ChatUser) logReceive(message string, senderName string) {
	// Match Python output format exactly using fmt.Printf for direct stdout
	fmt.Printf("%s received: %s (from %s)\n", u.name, message, senderName)
}
