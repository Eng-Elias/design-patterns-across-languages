package chat_room_test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	cr "mediator_pattern_chat_room_go/chat_room"
)

// --- Mock User for Testing ---

// MockUser implements the User interface for testing purposes.
type MockUser struct {
	name            string
	mediator        cr.ChatMediator
	receivedCalls   []string // Store received messages for assertion
	sendCalled      bool     // Flag to check if Send was called
	setMediatorCalled bool   // Flag to check if SetMediator was called
	mutex           sync.Mutex // Protect receivedCalls
}

func NewMockUser(name string) *MockUser {
	return &MockUser{
		name:          name,
		receivedCalls: make([]string, 0),
	}
}

func (mu *MockUser) Send(message string) {
	mu.sendCalled = true // Mark send as called
	if mu.mediator != nil {
		mu.mediator.SendMessage(message, mu)
	} else {
		fmt.Printf("'%s' cannot send message: Mediator not set.\n", mu.name) // Simulate behavior
	}
}

func (mu *MockUser) Receive(message string, senderName string) {
	mu.mutex.Lock()
	defer mu.mutex.Unlock()
	call := fmt.Sprintf("%s received: %s (from %s)", mu.name, message, senderName) // Match new format
	mu.receivedCalls = append(mu.receivedCalls, call)
}

func (mu *MockUser) GetName() string {
	return mu.name
}

func (mu *MockUser) SetMediator(mediator cr.ChatMediator) {
	mu.mediator = mediator
	mu.setMediatorCalled = true
}

// Helper to check if a specific message was received
func (mu *MockUser) WasCalledWith(message string, senderName string) bool {
	mu.mutex.Lock()
	defer mu.mutex.Unlock()
	expectedCall := fmt.Sprintf("%s received: %s (from %s)", mu.name, message, senderName) // Match new format
	for _, call := range mu.receivedCalls {
		if call == expectedCall {
			return true
		}
	}
	return false
}

// Helper to get the number of times Receive was called
func (mu *MockUser) ReceiveCallCount() int {
	mu.mutex.Lock()
	defer mu.mutex.Unlock()
	return len(mu.receivedCalls)
}

// Helper function to capture stdout
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf) // Temporarily redirect log output
	// Redirect standard fmt output (less common in tests, but useful here)
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // Execute the function

	// Restore outputs
	w.Close()
	os.Stdout = old
	log.SetOutput(os.Stderr) // Restore log output

	captured := buf.String() + readPipe(r)
	return captured
}

func readPipe(r *os.File) string {
	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var readBuf bytes.Buffer
		_, _ = readBuf.ReadFrom(r)
		outC <- readBuf.String()
	}()
	return <-outC
}


// --- Test Functions ---

func TestChatRoom_AddUser(t *testing.T) {
	mediator := cr.NewChatRoom()
	user1 := NewMockUser("Alice")

	captureOutput(func() {
		mediator.AddUser(user1)
	})

	// Check if user was added internally (cannot directly check map easily without export/helper)
	// Instead, check SetMediator was called on the user
	if !user1.setMediatorCalled {
		t.Errorf("Expected SetMediator to be called on user, but it wasn't")
	}

	// Test adding another user
	user2 := NewMockUser("Bob")
	captureOutput(func() {
		mediator.AddUser(user2)
	})
	if !user2.setMediatorCalled {
		t.Errorf("Expected SetMediator to be called on second user, but it wasn't")
	}
}

func TestChatRoom_AddUser_Duplicate(t *testing.T) {
	mediator := cr.NewChatRoom()
	user1 := NewMockUser("Alice")

	captureOutput(func() {
		mediator.AddUser(user1)
	})
	user1.setMediatorCalled = false // Reset flag for the check below

	output := captureOutput(func() {
		mediator.AddUser(user1) // Try adding the same user
	})

	if user1.setMediatorCalled {
		t.Errorf("SetMediator should not be called again for a duplicate user")
	}
	expectedLog := "--- Alice is already in the chat room. Cannot add again. ---" // Updated format
	if !strings.Contains(output, expectedLog) {
		t.Errorf("Expected log message '%s' not found in output: %s", expectedLog, output)
	}
}


func TestChatRoom_RemoveUser(t *testing.T) {
	mediator := cr.NewChatRoom()
	user1 := NewMockUser("Alice")
	user2 := NewMockUser("Bob")
	captureOutput(func() {
		mediator.AddUser(user1)
		mediator.AddUser(user2)
	})


	output := captureOutput(func() {
		mediator.RemoveUser(user1)
	})

	expectedLog := "--- Alice removed from the chat room. ---" // Updated format
	if !strings.Contains(output, expectedLog) {
		t.Errorf("Expected log message '%s' not found in output: %s", expectedLog, output)
	}

	// Verify user1 doesn't receive messages anymore (send from user2)
	user1.receivedCalls = make([]string, 0) // Clear previous potential calls if any
	msg := "Are you still there Alice?"
	captureOutput(func() {
		 user2.Send(msg)
	})

	if user1.ReceiveCallCount() > 0 {
		t.Errorf("Removed user 'Alice' should not have received messages, but received %d", user1.ReceiveCallCount())
	}
}


func TestChatRoom_RemoveUser_NotFound(t *testing.T) {
	mediator := cr.NewChatRoom()
	user1 := NewMockUser("Alice") // User exists but not added
	userNotInRoom := NewMockUser("Charlie")

	// Add Alice
	captureOutput(func() {
		mediator.AddUser(user1)
	})

	output := captureOutput(func() {
		mediator.RemoveUser(userNotInRoom) // Try removing user not in the room
	})

	expectedLog := "--- Charlie is not in the chat room. Cannot remove. ---" // Updated format
	if !strings.Contains(output, expectedLog) {
		t.Errorf("Expected log message '%s' not found in output: %s", expectedLog, output)
	}
}

func TestChatRoom_SendMessage(t *testing.T) {
	mediator := cr.NewChatRoom()
	alice := NewMockUser("Alice")
	bob := NewMockUser("Bob")
	charlie := NewMockUser("Charlie")

	captureOutput(func() {
		mediator.AddUser(alice)
		mediator.AddUser(bob)
		mediator.AddUser(charlie)
	})

	message := "Hello Team!"
	captureOutput(func() {
		alice.Send(message) // Alice sends a message
	})

	// Check if Bob and Charlie received the message, but Alice didn't
	if alice.ReceiveCallCount() != 0 {
		t.Errorf("Sender 'Alice' should not receive their own message, but received %d calls", alice.ReceiveCallCount())
	}
	if bob.ReceiveCallCount() != 1 || !bob.WasCalledWith(message, "Alice") {
		t.Errorf("User 'Bob' should have received message '%s' from 'Alice' exactly once. Calls: %d, Found: %t", message, bob.ReceiveCallCount(), bob.WasCalledWith(message, "Alice"))
	}
	if charlie.ReceiveCallCount() != 1 || !charlie.WasCalledWith(message, "Alice") {
	 t.Errorf("User 'Charlie' should have received message '%s' from 'Alice' exactly once. Calls: %d, Found: %t", message, charlie.ReceiveCallCount(), charlie.WasCalledWith(message, "Alice"))
	}
}


func TestChatUser_Send(t *testing.T) {
	mediator := cr.NewChatRoom() // Use actual ChatRoom
	user := cr.NewChatUser("Tester") // Use actual ChatUser

	captureOutput(func(){
		mediator.AddUser(user) // Add user, which sets the mediator
	})

	output := captureOutput(func(){
		user.Send("Test Send")
	})

	// Check if mediator's SendMessage was effectively called (by checking output logs)
	expectedLog := "--- Tester sends message: 'Test Send' ---" // Updated format
	if !strings.Contains(output, expectedLog) {
		t.Errorf("Expected log message '%s' from mediator not found in output: %s", expectedLog, output)
	}
}


func TestChatUser_Send_NoMediator(t *testing.T) {
	user := cr.NewChatUser("Tester") // User without a mediator set

	output := captureOutput(func() {
		user.Send("Test Send No Mediator")
	})

	if output != "" {
		t.Errorf("Expected no output when sending with no mediator, but got: %s", output)
	}
}