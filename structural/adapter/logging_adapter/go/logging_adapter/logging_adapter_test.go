package logging_adapter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// --- Mocks ---

// MockThirdPartyLogger mocks the Adaptee
type MockThirdPartyLogger struct {
	mock.Mock
}

func (m *MockThirdPartyLogger) Record(severity string, message string) {
	m.Called(severity, message)
}

// MockLogger mocks the Target interface
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) LogInfo(message string) {
	m.Called(message)
}

func (m *MockLogger) LogWarning(message string) {
	m.Called(message)
}

func (m *MockLogger) LogError(message string) {
	m.Called(message)
}

// --- Tests ---

func TestLoggerAdapter_CallsAdapteeCorrectly(t *testing.T) {
	// Arrange
	mockAdaptee := new(MockThirdPartyLogger)
	adapter := NewLoggerAdapter(mockAdaptee) // Adapter takes the concrete type pointer

	// Set up expectations
	mockAdaptee.On("Record", "info", "Info message").Return()
	mockAdaptee.On("Record", "warning", "Warning message").Return()
	mockAdaptee.On("Record", "error", "Error message").Return()

	// Act
	adapter.LogInfo("Info message")
	adapter.LogWarning("Warning message")
	adapter.LogError("Error message")

	// Assert
	mockAdaptee.AssertExpectations(t)
	mockAdaptee.AssertNumberOfCalls(t, "Record", 3)
}

func TestApplicationService_UsesLogger(t *testing.T) {
	// Arrange
	mockTargetLogger := new(MockLogger) // Mock the Target interface
	appService := NewApplicationService(mockTargetLogger)

	// --- Act & Assert: Successful operation ---
	validData := "Valid Data"
	mockTargetLogger.On("LogInfo", fmt.Sprintf("Starting operation with data: %s", validData)).Return()
	mockTargetLogger.On("LogInfo", "Operation completed successfully.").Return()

	appService.PerformOperation(validData)

	mockTargetLogger.AssertCalled(t, "LogInfo", fmt.Sprintf("Starting operation with data: %s", validData))
	mockTargetLogger.AssertCalled(t, "LogInfo", "Operation completed successfully.")
mockTargetLogger.AssertNotCalled(t, "LogWarning", mock.Anything) // Pass mock.Anything if you don't care about the message
mockTargetLogger.AssertNotCalled(t, "LogError", mock.Anything)

mockTargetLogger.AssertNumberOfCalls(t, "LogInfo", 2)
	mockTargetLogger.AssertNumberOfCalls(t, "LogWarning", 0)
	mockTargetLogger.AssertNumberOfCalls(t, "LogError", 0)
	mockTargetLogger.AssertExpectations(t) // Ensure all expected calls were made

	// --- Reset mock expectations for next scenario ---
	mockTargetLogger.ExpectedCalls = nil // Clear previous expectations
    // We also need to reset the 'Calls' slice if we want AssertNumberOfCalls to work correctly from zero
    // A simpler way is often to create a new mock for each sub-test, but this works too.
    mockTargetLogger.Calls = nil


	// --- Act & Assert: Operation with warning ---
	shortData := "shrt" // Changed data to be < 5 chars
	mockTargetLogger.On("LogInfo", fmt.Sprintf("Starting operation with data: %s", shortData)).Return()
	mockTargetLogger.On("LogWarning", fmt.Sprintf("Data '%s' is quite short.", shortData)).Return()
	mockTargetLogger.On("LogInfo", "Operation completed successfully.").Return()

	appService.PerformOperation(shortData)

	mockTargetLogger.AssertCalled(t, "LogInfo", fmt.Sprintf("Starting operation with data: %s", shortData))
	mockTargetLogger.AssertCalled(t, "LogWarning", fmt.Sprintf("Data '%s' is quite short.", shortData))
	mockTargetLogger.AssertCalled(t, "LogInfo", "Operation completed successfully.")
mockTargetLogger.AssertNotCalled(t, "LogError", mock.Anything)

	mockTargetLogger.AssertNumberOfCalls(t, "LogInfo", 2) // Start and End
	mockTargetLogger.AssertNumberOfCalls(t, "LogWarning", 1)
	mockTargetLogger.AssertNumberOfCalls(t, "LogError", 0)
	mockTargetLogger.AssertExpectations(t)

    // --- Reset mock expectations for next scenario ---
    mockTargetLogger.ExpectedCalls = nil
    mockTargetLogger.Calls = nil

	// --- Act & Assert: Operation with error ---
	emptyData := ""
	mockTargetLogger.On("LogInfo", fmt.Sprintf("Starting operation with data: %s", emptyData)).Return()
	mockTargetLogger.On("LogError", "Operation failed: data cannot be empty").Return() // Match the specific error message

	appService.PerformOperation(emptyData)

	mockTargetLogger.AssertCalled(t, "LogInfo", fmt.Sprintf("Starting operation with data: %s", emptyData))
	mockTargetLogger.AssertCalled(t, "LogError", "Operation failed: data cannot be empty")
mockTargetLogger.AssertNotCalled(t, "LogWarning", mock.Anything)

	mockTargetLogger.AssertNumberOfCalls(t, "LogInfo", 1) // Only start message
	mockTargetLogger.AssertNumberOfCalls(t, "LogWarning", 0)
	mockTargetLogger.AssertNumberOfCalls(t, "LogError", 1)
	mockTargetLogger.AssertExpectations(t)
}
