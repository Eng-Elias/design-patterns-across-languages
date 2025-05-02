from abc import ABC, abstractmethod

# --- Command Interface ---
class Command(ABC):
    """The Command interface declares a method for executing a command."""
    @abstractmethod
    def execute(self) -> None:
        pass

# --- Receivers (Perform the actual work) ---
class EmailService:
    """Knows how to send emails."""
    def send_email(self, recipient: str, subject: str, body: str) -> None:
        print(f"Sending email to {recipient} with subject '{subject}'...")
        print(f"Body: {body}")
        print("Email sent.")

class ReportGenerator:
    """Knows how to generate reports."""
    def generate_report(self, report_type: str, output_path: str) -> None:
        print(f"Generating {report_type} report...")
        print(f"Saving report to {output_path}...")
        print("Report generated.")

class DatabaseService:
    """Knows how to perform database operations."""
    def run_backup(self, backup_name: str) -> None:
        print(f"Starting database backup '{backup_name}'...")
        # Simulate work
        print(f"Database backup '{backup_name}' completed.")


# --- Concrete Commands ---
class SendEmailCommand(Command):
    """Concrete command to send an email."""
    def __init__(self, service: EmailService, recipient: str, subject: str, body: str):
        self._service = service
        self._recipient = recipient
        self._subject = subject
        self._body = body

    def execute(self) -> None:
        self._service.send_email(self._recipient, self._subject, self._body)

class GenerateReportCommand(Command):
    """Concrete command to generate a report."""
    def __init__(self, service: ReportGenerator, report_type: str, output_path: str):
        self._service = service
        self._report_type = report_type
        self._output_path = output_path

    def execute(self) -> None:
        self._service.generate_report(self._report_type, self._output_path)

class RunDatabaseBackupCommand(Command):
    """Concrete command to run a database backup."""
    def __init__(self, service: DatabaseService, backup_name: str):
        self._service = service
        self._backup_name = backup_name

    def execute(self) -> None:
        self._service.run_backup(self._backup_name)
