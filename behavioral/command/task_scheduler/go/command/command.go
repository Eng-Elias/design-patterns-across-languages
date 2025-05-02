package command

import (
	"fmt"
	"time"
)

// --- Command Interface ---
type Command interface {
	Execute() error // Commands can potentially fail
}

// --- Receivers (Perform the actual work) ---

// EmailService knows how to send emails.
type EmailService struct{}

func (e *EmailService) SendEmail(recipient, subject, body string) error {
	fmt.Printf("Sending email to %s with subject '%s'...\n", recipient, subject)
	fmt.Printf("Body: %s\n", body)
	// Simulate work/potential failure
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Email sent.")
	return nil // Assume success for example
}

// ReportGenerator knows how to generate reports.
type ReportGenerator struct{}

func (r *ReportGenerator) GenerateReport(reportType, outputPath string) error {
	fmt.Printf("Generating %s report...\n", reportType)
	fmt.Printf("Saving report to %s...\n", outputPath)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Report generated.")
	return nil
}

// DatabaseService knows how to perform database operations.
type DatabaseService struct{}

func (d *DatabaseService) RunBackup(backupName string) error {
	fmt.Printf("Starting database backup '%s'...\n", backupName)
	time.Sleep(150 * time.Millisecond)
	fmt.Printf("Database backup '%s' completed.\n", backupName)
	return nil
}

// --- Concrete Commands ---

// SendEmailCommand encapsulates sending an email.
type SendEmailCommand struct {
	Service   *EmailService // Use pointer to receiver
	Recipient string
	Subject   string
	Body      string
}

// NewSendEmailCommand is a constructor function for SendEmailCommand.
func NewSendEmailCommand(service *EmailService, recipient, subject, body string) *SendEmailCommand {
	return &SendEmailCommand{
		Service:   service,
		Recipient: recipient,
		Subject:   subject,
		Body:      body,
	}
}

func (c *SendEmailCommand) Execute() error {
	if c.Service == nil {
		return fmt.Errorf("email service is nil")
	}
	fmt.Printf("Sending email to %s\n", c.Recipient)
	fmt.Printf("Subject: %s\n", c.Subject) // Added subject print for clarity, adjust test if needed
	fmt.Printf("Body: %s\n", c.Body)
	err := c.Service.SendEmail(c.Recipient, c.Subject, c.Body)
	if err == nil {
		fmt.Println("Email sent.")
	}
	return err
}
func (c *SendEmailCommand) String() string { // Implement Stringer for logging
	return fmt.Sprintf("SendEmailCommand(to:%s, subject:%s)", c.Recipient, c.Subject)
}


// GenerateReportCommand encapsulates generating a report.
type GenerateReportCommand struct {
	Service    *ReportGenerator
	ReportType string
	OutputPath string
}

// NewGenerateReportCommand is a constructor.
func NewGenerateReportCommand(service *ReportGenerator, reportType, outputPath string) *GenerateReportCommand {
	 return &GenerateReportCommand{
		 Service:    service,
		 ReportType: reportType,
		 OutputPath: outputPath,
	 }
}


func (c *GenerateReportCommand) Execute() error {
	if c.Service == nil {
		return fmt.Errorf("report generator service is nil")
	}
	fmt.Printf("Generating %s report...\n", c.ReportType)
	fmt.Printf("Saving report to %s\n", c.OutputPath)
	err := c.Service.GenerateReport(c.ReportType, c.OutputPath)
	if err == nil {
		fmt.Println("Report generated.")
	}
	return err
}
func (c *GenerateReportCommand) String() string {
	return fmt.Sprintf("GenerateReportCommand(type:%s, path:%s)", c.ReportType, c.OutputPath)
}

// RunDatabaseBackupCommand encapsulates running a database backup.
type RunDatabaseBackupCommand struct {
	Service    *DatabaseService
	BackupName string
}

// NewRunDatabaseBackupCommand is a constructor.
func NewRunDatabaseBackupCommand(service *DatabaseService, backupName string) *RunDatabaseBackupCommand {
	return &RunDatabaseBackupCommand{
		Service:    service,
		BackupName: backupName,
	}
}

func (c *RunDatabaseBackupCommand) Execute() error {
	if c.Service == nil {
		return fmt.Errorf("database service is nil")
	}
	fmt.Printf("Starting database backup '%s'...\n", c.BackupName)
	err := c.Service.RunBackup(c.BackupName)
	if err == nil {
		fmt.Printf("Database backup '%s' completed.\n", c.BackupName)
	}
	return err
}
func (c *RunDatabaseBackupCommand) String() string {
	return fmt.Sprintf("RunDatabaseBackupCommand(name:%s)", c.BackupName)
}
