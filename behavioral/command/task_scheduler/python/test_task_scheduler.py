import unittest
from unittest.mock import MagicMock, call, patch

# Import classes from sibling modules
from commands import (
    Command, EmailService, ReportGenerator, DatabaseService,
    SendEmailCommand, GenerateReportCommand, RunDatabaseBackupCommand
)
from scheduler import TaskScheduler

class TestTaskSchedulerCommand(unittest.TestCase):

    def setUp(self):
        # Create mock receivers for each test
        self.mock_email_service = MagicMock(spec=EmailService)
        self.mock_report_generator = MagicMock(spec=ReportGenerator)
        self.mock_db_service = MagicMock(spec=DatabaseService)

        # Create the scheduler instance
        self.scheduler = TaskScheduler()

    @patch('builtins.print') # Mock print to suppress test output clutter
    def test_add_and_run_single_task(self, mock_print):
        """Test adding and running a single command."""
        # Create a command with the mock receiver
        backup_cmd = RunDatabaseBackupCommand(self.mock_db_service, "test_backup")

        # Add the task
        self.scheduler.add_task(backup_cmd)

        # Run tasks
        self.scheduler.run_pending_tasks()

        # Assert that the correct receiver method was called once with correct args
        self.mock_db_service.run_backup.assert_called_once_with("test_backup")
        # Assert other mocks were not called
        self.mock_email_service.send_email.assert_not_called()
        self.mock_report_generator.generate_report.assert_not_called()

    @patch('builtins.print')
    def test_add_and_run_multiple_tasks(self, mock_print):
        """Test adding and running multiple different commands."""
        # Create commands
        email_cmd = SendEmailCommand(self.mock_email_service, "test@dev.null", "Subj", "Body")
        report_cmd = GenerateReportCommand(self.mock_report_generator, "Test Report", "/tmp/test.rpt")
        backup_cmd = RunDatabaseBackupCommand(self.mock_db_service, "multi_backup")

        # Add tasks
        self.scheduler.add_task(email_cmd)
        self.scheduler.add_task(report_cmd)
        self.scheduler.add_task(backup_cmd)

        # Run tasks
        self.scheduler.run_pending_tasks()

        # Assert each mock method was called once with correct args
        self.mock_email_service.send_email.assert_called_once_with("test@dev.null", "Subj", "Body")
        self.mock_report_generator.generate_report.assert_called_once_with("Test Report", "/tmp/test.rpt")
        self.mock_db_service.run_backup.assert_called_once_with("multi_backup")

    @patch('builtins.print')
    def test_run_no_tasks(self, mock_print):
        """Test running the scheduler when no tasks have been added."""
        self.scheduler.run_pending_tasks()

        # Assert that no receiver methods were called
        self.mock_email_service.send_email.assert_not_called()
        self.mock_report_generator.generate_report.assert_not_called()
        self.mock_db_service.run_backup.assert_not_called()
        # Check if the "No tasks to run" message was printed (optional)
        mock_print.assert_any_call("No tasks to run.")

    @patch('builtins.print')
    def test_task_execution_order(self, mock_print):
        """Test if tasks execute in the order they were added (FIFO)."""
        scheduler = TaskScheduler()

        # Create mock commands
        cmd1 = MagicMock(spec=Command)
        cmd2 = MagicMock(spec=Command)
        cmd3 = MagicMock(spec=Command)

        # Add commands to the scheduler
        scheduler.add_task(cmd1)
        scheduler.add_task(cmd2)
        scheduler.add_task(cmd3)

        # Run the scheduler
        scheduler.run_pending_tasks()

        # Expected calls to the execute method of each command instance
        expected_execution_calls = [
            call.execute(),
            call.execute(),
            call.execute()
        ]

        # Check if execute was called on each mock in the correct order
        cmd1.execute.assert_called_once()
        cmd2.execute.assert_called_once()
        cmd3.execute.assert_called_once()
