package main

import (
	// Use the correct import paths based on your go.mod file
	"fmt"
	"log"
	"time"

	"command_pattern_task_scheduler_go/command"
	"command_pattern_task_scheduler_go/scheduler"
)

func main() {
	// Important: Replace the import paths above with the ones matching your project structure and go.mod file.
	log.Println("--- Command Pattern: Task Scheduler ---")

	// 1. Create Receiver instances
	emailService := &command.EmailService{}
	reportService := &command.ReportGenerator{}
	dbService := &command.DatabaseService{}

	// 2. Create the Invoker (Scheduler)
	taskScheduler := scheduler.NewTaskScheduler()

	// 3. Create Concrete Command instances
	// Using constructor functions for clarity
	cmd1 := command.NewSendEmailCommand(
		emailService,
		"ceo@mycorp.com",
		"Urgent: Server Status",
		"All production servers are stable.",
	)

	cmd2 := command.NewGenerateReportCommand(
		reportService,
		"End of Day Summary",
		fmt.Sprintf("/var/log/summary_%s.log", time.Now().Format("20060102")), // Current date Friday, May 2, 2025 -> 20250502
	)

	cmd3 := command.NewRunDatabaseBackupCommand(
		dbService,
		fmt.Sprintf("prod_db_snapshot_%s", time.Now().Format("20060102_1504")), // Current time 13:13 -> _1313
	)

	// 4. Add Commands to the Scheduler
	taskScheduler.AddTask(cmd1)
	taskScheduler.AddTask(cmd2)
	taskScheduler.AddTask(cmd3)

	// 5. Run the scheduler
	taskScheduler.RunPendingTasks()

	log.Println("--- Scheduler finished ---")
}
