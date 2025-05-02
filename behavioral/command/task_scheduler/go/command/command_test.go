package command

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

// Helper to capture log/stdout output for testing receiver actions
func captureOutput(f func()) string {
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


func TestSendEmailCommand_Execute(t *testing.T) {
	service := &EmailService{}
	recipient := "test@example.com"
	subject := "Test Subject"
	body := "Test Body"
	cmd := NewSendEmailCommand(service, recipient, subject, body)

	var output string
	var err error
	output = captureOutput(func() {
		err = cmd.Execute()
	})


	if err != nil {
		t.Errorf("Execute() returned an unexpected error: %v", err)
	}
	if !strings.Contains(output, fmt.Sprintf("Sending email to %s", recipient)) {
		t.Errorf("Output missing expected email sending message. Got: %s", output)
	}
    if !strings.Contains(output, fmt.Sprintf("Body: %s", body)) {
		t.Errorf("Output missing expected email body. Got: %s", output)
	}
	if !strings.Contains(output, "Email sent.") {
		t.Errorf("Output missing expected email sent confirmation. Got: %s", output)
	}

    // Test nil service
    cmd.Service = nil
    err = cmd.Execute()
    if err == nil {
        t.Errorf("Execute() with nil service should return an error, but got nil")
    }
}

func TestGenerateReportCommand_Execute(t *testing.T) {
	service := &ReportGenerator{}
	reportType := "TestReport"
	outputPath := "/tmp/test.rep"
	cmd := NewGenerateReportCommand(service, reportType, outputPath)

	var output string
	var err error
    output = captureOutput(func() {
		err = cmd.Execute()
	})


	if err != nil {
		t.Errorf("Execute() returned an unexpected error: %v", err)
	}
	if !strings.Contains(output, fmt.Sprintf("Generating %s report...", reportType)) {
		t.Errorf("Output missing expected report generation message. Got: %s", output)
	}
    if !strings.Contains(output, fmt.Sprintf("Saving report to %s", outputPath)) {
		t.Errorf("Output missing expected report saving message. Got: %s", output)
	}
	if !strings.Contains(output, "Report generated.") {
		t.Errorf("Output missing expected report generated confirmation. Got: %s", output)
	}

    // Test nil service
    cmd.Service = nil
    err = cmd.Execute()
    if err == nil {
        t.Errorf("Execute() with nil service should return an error, but got nil")
    }
}


func TestRunDatabaseBackupCommand_Execute(t *testing.T) {
    service := &DatabaseService{}
	backupName := "test_backup_cmd"
	cmd := NewRunDatabaseBackupCommand(service, backupName)

	var output string
	var err error
    output = captureOutput(func() {
		err = cmd.Execute()
	})


	if err != nil {
		t.Errorf("Execute() returned an unexpected error: %v", err)
	}
    if !strings.Contains(output, fmt.Sprintf("Starting database backup '%s'...", backupName)) {
		t.Errorf("Output missing expected backup start message. Got: %s", output)
	}
    if !strings.Contains(output, fmt.Sprintf("Database backup '%s' completed.", backupName)) {
		t.Errorf("Output missing expected backup complete message. Got: %s", output)
	}

     // Test nil service
    cmd.Service = nil
    err = cmd.Execute()
    if err == nil {
        t.Errorf("Execute() with nil service should return an error, but got nil")
    }
}

// Test Stringer implementations (optional but good)
func TestCommand_String(t *testing.T) {
	emailCmd := NewSendEmailCommand(nil, "a@b.c", "Subj", "Body")
	reportCmd := NewGenerateReportCommand(nil, "Type", "/path")
	backupCmd := NewRunDatabaseBackupCommand(nil, "Name")

	if !strings.Contains(emailCmd.String(), "SendEmailCommand") { t.Errorf("Bad string format for email cmd") }
	if !strings.Contains(reportCmd.String(), "GenerateReportCommand") { t.Errorf("Bad string format for report cmd") }
	if !strings.Contains(backupCmd.String(), "RunDatabaseBackupCommand") { t.Errorf("Bad string format for backup cmd") }
}
