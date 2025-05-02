package scheduler

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"command_pattern_task_scheduler_go/command"
)

// MockCommand specifically for testing scheduler execution order/calls
type MockCommand struct {
	ID          string
	Executed    bool
	ExecuteFunc func() error // Allow custom execute behavior for testing errors etc.
}

func (m *MockCommand) Execute() error {
	m.Executed = true
    fmt.Printf("MockCommand %s executed.\n", m.ID) // Print to captured output
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc()
	}
	return nil
}

func (m *MockCommand) String() string {
	return fmt.Sprintf("MockCommand(ID:%s)", m.ID)
}

// Helper to capture log/stdout output
func captureSchedulerOutput(f func()) string {
	// Create a pipe
	r, w, _ := os.Pipe()

	// Save original stdout and logger output
	originalStdout := os.Stdout
	originalLogOutput := log.Writer()

	// Redirect stdout and log output to the pipe writer
	os.Stdout = w
	log.SetOutput(w)

	// Channel to capture output from the reading goroutine
	outChan := make(chan string)
	// Start concurrent reader goroutine
	go func() {
		defer r.Close() // Close reader when done
		outputBytes, _ := io.ReadAll(r)
		outChan <- string(outputBytes)
	}()

	// Use defer to restore original outputs
	defer func() {
		os.Stdout = originalStdout
		log.SetOutput(originalLogOutput)
	}()

	// Execute the function whose output we want to capture
	f()

	// Close the writer *after* the function finishes
	// This signals EOF to the reader goroutine, causing io.ReadAll to return
	w.Close()

	// Wait for the reader goroutine to finish and get the output
	capturedOutput := <-outChan
	return capturedOutput
}

func TestTaskScheduler_AddAndRunTasks(t *testing.T) {
	scheduler := NewTaskScheduler()

	cmd1 := &MockCommand{ID: "CMD1"}
	cmd2 := &MockCommand{ID: "CMD2"}

	scheduler.AddTask(cmd1)
	scheduler.AddTask(cmd2)

	if scheduler.GetTaskCount() != 2 {
		t.Errorf("Expected task count to be 2, got %d", scheduler.GetTaskCount())
	}

	output := captureSchedulerOutput(func() {
		scheduler.RunPendingTasks()
	})


	if !cmd1.Executed {
		t.Error("Expected cmd1 to be executed, but it wasn't")
	}
	if !cmd2.Executed {
		t.Error("Expected cmd2 to be executed, but it wasn't")
	}

    // Check output for execution messages
    if !strings.Contains(output, "Executing task [1]: MockCommand(ID:CMD1)") {
        t.Errorf("Output missing execution message for CMD1. Got: %s", output)
    }
     if !strings.Contains(output, "MockCommand CMD1 executed.") {
        t.Errorf("Output missing confirmation message for CMD1. Got: %s", output)
    }

    if !strings.Contains(output, "Executing task [2]: MockCommand(ID:CMD2)") {
        t.Errorf("Output missing execution message for CMD2. Got: %s", output)
    }
      if !strings.Contains(output, "MockCommand CMD2 executed.") {
        t.Errorf("Output missing confirmation message for CMD2. Got: %s", output)
    }


	if scheduler.GetTaskCount() != 0 {
		t.Errorf("Expected task count to be 0 after running, got %d", scheduler.GetTaskCount())
	}
}

func TestTaskScheduler_RunNoTasks(t *testing.T) {
	scheduler := NewTaskScheduler()

	if scheduler.GetTaskCount() != 0 {
		t.Errorf("Expected initial task count to be 0, got %d", scheduler.GetTaskCount())
	}

	output := captureSchedulerOutput(func() {
		scheduler.RunPendingTasks()
	})

	if !strings.Contains(output, "No tasks to run.") {
		t.Errorf("Expected 'No tasks to run.' message, but not found in output: %s", output)
	}
	if scheduler.GetTaskCount() != 0 {
		t.Errorf("Expected task count to remain 0 after running, got %d", scheduler.GetTaskCount())
	}
}

func TestTaskScheduler_TaskExecutionOrder(t *testing.T) {
    scheduler := NewTaskScheduler()
	var executionOrder []string

	// Create mock commands that record their execution order
	cmdA := &MockCommand{ID: "A", ExecuteFunc: func() error { executionOrder = append(executionOrder, "A"); return nil }}
	cmdB := &MockCommand{ID: "B", ExecuteFunc: func() error { executionOrder = append(executionOrder, "B"); return nil }}
	cmdC := &MockCommand{ID: "C", ExecuteFunc: func() error { executionOrder = append(executionOrder, "C"); return nil }}

    scheduler.AddTask(cmdA)
    scheduler.AddTask(cmdB)
    scheduler.AddTask(cmdC)

    captureSchedulerOutput(func() {
		scheduler.RunPendingTasks()
	})


    expectedOrder := []string{"A", "B", "C"}
    if len(executionOrder) != len(expectedOrder) {
        t.Fatalf("Expected execution order length %d, got %d", len(expectedOrder), len(executionOrder))
    }
    for i, id := range expectedOrder {
        if executionOrder[i] != id {
            t.Errorf("Expected execution order at index %d to be %s, got %s. Full order: %v", i, id, executionOrder[i], executionOrder)
        }
    }
}

func TestTaskScheduler_ErrorHandling(t *testing.T) {
     scheduler := NewTaskScheduler()
     errMsg := "Task failed intentionally"

     cmdOk1 := &MockCommand{ID: "OK1"}
     cmdFail := &MockCommand{ID: "FAIL", ExecuteFunc: func() error { return fmt.Errorf("%s", errMsg) }}
     cmdOk2 := &MockCommand{ID: "OK2"}

     scheduler.AddTask(cmdOk1)
     scheduler.AddTask(cmdFail)
     scheduler.AddTask(cmdOk2)

     output := captureSchedulerOutput(func() {
		scheduler.RunPendingTasks()
	 })

     if !cmdOk1.Executed { t.Error("cmdOk1 should have executed") }
     if !cmdFail.Executed { t.Error("cmdFail should have been attempted (marked executed by mock)") }
     if !cmdOk2.Executed { t.Error("cmdOk2 should have executed after failure") }

     // Check if the error message was logged
     if !strings.Contains(output, errMsg) {
         t.Errorf("Expected error message '%s' not found in log output: %s", errMsg, output)
     }
      if !strings.Contains(output, "Error executing task MockCommand(ID:FAIL)") {
         t.Errorf("Expected error log prefix not found in log output: %s", output)
     }

     // Check if OK tasks completed normally
     if !strings.Contains(output, "Task MockCommand(ID:OK1) completed.") { t.Errorf("Missing completion message for OK1")}
     if !strings.Contains(output, "Task MockCommand(ID:OK2) completed.") { t.Errorf("Missing completion message for OK2")}


     if scheduler.GetTaskCount() != 0 {
		 t.Errorf("Expected task count to be 0 after running, got %d", scheduler.GetTaskCount())
	 }
}


// Example test using real commands (relies on command package tests)
// Requires receivers to be instantiated
func TestTaskScheduler_WithRealCommands(t *testing.T) {
    // This test relies more on capturing output from the actual receiver methods
	// It's more of an integration test between scheduler and commands.
	scheduler := NewTaskScheduler()

	// Create real receivers
	emailSvc := &command.EmailService{}
	dbSvc := &command.DatabaseService{}

	// Create real commands
	cmdEmail := command.NewSendEmailCommand(emailSvc, "real@test.com", "Real", "Msg")
	cmdBackup := command.NewRunDatabaseBackupCommand(dbSvc, "real_backup")

	scheduler.AddTask(cmdEmail)
	scheduler.AddTask(cmdBackup)

	output := captureSchedulerOutput(func() {
		scheduler.RunPendingTasks()
	})

	// Check output for receiver messages
	if !strings.Contains(output, "Sending email to real@test.com") {
		t.Error("Missing email service output")
	}
	if !strings.Contains(output, "Starting database backup 'real_backup'") {
		t.Error("Missing database service output")
	}
    if scheduler.GetTaskCount() != 0 {
		t.Errorf("Expected task count to be 0 after running, got %d", scheduler.GetTaskCount())
	}
}
