# Import necessary classes from sibling modules
from commands import (
    EmailService, ReportGenerator, DatabaseService,
    SendEmailCommand, GenerateReportCommand, RunDatabaseBackupCommand
)
from scheduler import TaskScheduler

def main():
    # 1. Create Receiver instances
    email_service = EmailService()
    report_service = ReportGenerator()
    db_service = DatabaseService()

    # 2. Create the Invoker (Scheduler)
    scheduler = TaskScheduler()

    # 3. Create Concrete Command instances, linking receivers and parameters
    email_command = SendEmailCommand(
        email_service,
        "user@example.com",
        "Scheduled Notification",
        "This is your scheduled email."
    )

    report_command = GenerateReportCommand(
        report_service,
        "Monthly Sales",
        "/reports/sales_may_2025.pdf"
    )

    backup_command = RunDatabaseBackupCommand(
        db_service,
        "daily_backup_20250502"
    )

    # 4. Add Commands to the Scheduler
    scheduler.add_task(email_command)
    scheduler.add_task(report_command)
    scheduler.add_task(backup_command)

    # 5. Run the scheduler (execute all queued commands)
    scheduler.run_pending_tasks()

if __name__ == "__main__":
    main()
