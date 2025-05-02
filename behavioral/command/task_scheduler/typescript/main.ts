import {
  EmailService,
  ReportGenerator,
  DatabaseService,
  SendEmailCommand,
  GenerateReportCommand,
  RunDatabaseBackupCommand,
} from "./commands";
import { TaskScheduler } from "./scheduler";

async function main() {
  // 1. Create Receiver instances
  const emailService = new EmailService();
  const reportService = new ReportGenerator();
  const dbService = new DatabaseService();

  // 2. Create the Invoker (Scheduler)
  const scheduler = new TaskScheduler();

  // 3. Create Concrete Command instances
  const emailCommand = new SendEmailCommand(
    emailService,
    "friend@example.com",
    "Scheduled Check-in",
    "Just checking in as scheduled!"
  );

  const reportCommand = new GenerateReportCommand(
    reportService,
    "Quarterly Performance",
    "/data/reports/q2_2025_perf.xlsx" // Current date is May 2, 2025
  );

  const backupCommand = new RunDatabaseBackupCommand(
    dbService,
    "pre_deploy_backup_20250502" // Current date is May 2, 2025
  );

  // 4. Add Commands to the Scheduler
  scheduler.addTask(emailCommand);
  scheduler.addTask(reportCommand);
  scheduler.addTask(backupCommand);

  // 5. Run the scheduler (execute all queued commands)
  await scheduler.runPendingTasks();

  console.log("\nScheduler finished.");
}

// Execute the main function
main().catch((error) => {
  console.error("An error occurred in the main execution:", error);
});
