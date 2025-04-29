package logging_adapter

import (
	"fmt"
	"strings"
)

// --- Adaptee Interface --- (Interface defining required adaptee methods)
// This allows us to depend on abstraction, not implementation
type AdapteeLogger interface {
	Record(severity string, message string)
}

// --- Adaptee --- (The incompatible third-party library)
type ThirdPartyLogger struct{}

// Record is the incompatible logging method.
// *ThirdPartyLogger implicitly satisfies the AdapteeLogger interface.
func (tpl *ThirdPartyLogger) Record(severity string, message string) {
	fmt.Printf("[3rdPartyLogger - %s]: %s\n", strings.ToUpper(severity), message)
}

// --- Target --- (The interface the client code expects)
type Logger interface {
	LogInfo(message string)
	LogWarning(message string)
	LogError(message string)
}

// --- Adapter ---
type LoggerAdapter struct {
	adaptee AdapteeLogger // Holds an instance satisfying the AdapteeLogger interface
}

// NewLoggerAdapter creates a new adapter instance.
// Accepts any type that satisfies the AdapteeLogger interface.
func NewLoggerAdapter(adaptee AdapteeLogger) *LoggerAdapter {
	return &LoggerAdapter{adaptee: adaptee}
}

// LogInfo implements the Logger interface by calling the adaptee's Record method.
func (a *LoggerAdapter) LogInfo(message string) {
	a.adaptee.Record("info", message)
}

// LogWarning implements the Logger interface.
func (a *LoggerAdapter) LogWarning(message string) {
	a.adaptee.Record("warning", message)
}

// LogError implements the Logger interface.
func (a *LoggerAdapter) LogError(message string) {
	a.adaptee.Record("error", message)
}

// --- Client Code ---
type ApplicationService struct {
	logger Logger // Depends on the Target interface
}

// NewApplicationService creates a new service instance.
func NewApplicationService(logger Logger) *ApplicationService {
	return &ApplicationService{logger: logger}
}

// PerformOperation uses the Logger interface to log messages.
func (s *ApplicationService) PerformOperation(data string) {
	s.logger.LogInfo(fmt.Sprintf("Starting operation with data: %s", data))
	defer func() {
		if r := recover(); r != nil {
			s.logger.LogError(fmt.Sprintf("Operation failed unexpectedly: %v", r))
		}
	}()

	// Simulate an operation
	if data == "" {
		// Using fmt.Errorf to create an error object
		err := fmt.Errorf("data cannot be empty")
		s.logger.LogError(fmt.Sprintf("Operation failed: %v", err))
		return // Stop processing on error
	}
	if len(data) < 5 {
		s.logger.LogWarning(fmt.Sprintf("Data '%s' is quite short.", data))
	}

	// ... perform actual operation ...
	// Simulate success if no error occurred
	s.logger.LogInfo("Operation completed successfully.")
}
