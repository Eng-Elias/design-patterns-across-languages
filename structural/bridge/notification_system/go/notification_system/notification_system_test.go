package notification_system

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// --- Mock Implementation ---

// MockMessageSender mocks the MessageSender interface
type MockMessageSender struct {
	mock.Mock
}

// SendMessage is the mock implementation of the SendMessage method.
func (m *MockMessageSender) SendMessage(subject string, body string) {
	// This records that the method was called with these arguments.
	m.Called(subject, body)
}

// --- Tests ---

func TestInfoNotification_SendsCorrectly(t *testing.T) {
	// Arrange
	mockSender := new(MockMessageSender) // Create an instance of the mock
	infoNotification := NewInfoNotification(mockSender)
	message := "Test info message."
	expectedSubject := "Info"
	expectedBody := fmt.Sprintf("[INFO] %s", message)

	// Set up expectations on the mock
	// Expect SendMessage to be called once with specific arguments
	mockSender.On("SendMessage", expectedSubject, expectedBody).Return()

	// Act
	infoNotification.Send(message)

	// Assert
	// Verify that the expected methods were called on the mock object.
	mockSender.AssertExpectations(t)
}

func TestWarningNotification_SendsCorrectly(t *testing.T) {
	// Arrange
	mockSender := new(MockMessageSender)
	warningNotification := NewWarningNotification(mockSender)
	message := "Test warning message."
	expectedSubject := "Warning"
	expectedBody := fmt.Sprintf("[WARNING] %s", message)

	mockSender.On("SendMessage", expectedSubject, expectedBody).Return()

	// Act
	warningNotification.Send(message)

	// Assert
	mockSender.AssertExpectations(t)
}

func TestUrgentNotification_SendsCorrectly(t *testing.T) {
	// Arrange
	mockSender := new(MockMessageSender)
	urgentNotification := NewUrgentNotification(mockSender)
	message := "Test urgent message."
	expectedSubject := "** URGENT **"
	expectedBody := fmt.Sprintf("[URGENT ACTION REQUIRED] %s", message)

	mockSender.On("SendMessage", expectedSubject, expectedBody).Return()

	// Act
	urgentNotification.Send(message)

	// Assert
	mockSender.AssertExpectations(t)
}

func TestDifferentNotifications_UseSameSender(t *testing.T) {
	// Arrange
	mockSender := new(MockMessageSender)
	infoNotification := NewInfoNotification(mockSender)
	urgentNotification := NewUrgentNotification(mockSender)

	infoMessage := "Info 1"
	urgentMessage := "Urgent 1"
	expectedInfoSubject := "Info"
	expectedInfoBody := fmt.Sprintf("[INFO] %s", infoMessage)
	expectedUrgentSubject := "** URGENT **"
	expectedUrgentBody := fmt.Sprintf("[URGENT ACTION REQUIRED] %s", urgentMessage)

	// Set up expectations for both calls
	mockSender.On("SendMessage", expectedInfoSubject, expectedInfoBody).Return()
	mockSender.On("SendMessage", expectedUrgentSubject, expectedUrgentBody).Return()

	// Act
	infoNotification.Send(infoMessage)
	urgentNotification.Send(urgentMessage)

	// Assert
	mockSender.AssertExpectations(t)
	// Optionally, assert the number of calls if needed (though AssertExpectations often covers it)
	mockSender.AssertNumberOfCalls(t, "SendMessage", 2)
}
