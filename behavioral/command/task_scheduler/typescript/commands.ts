// --- Command Interface ---
export interface Command {
  execute(): Promise<void>; // Using Promise for potential async operations
}

// --- Receivers (Perform the actual work) ---
export class EmailService {
  public async sendEmail(
    recipient: string,
    subject: string,
    body: string
  ): Promise<void> {
    console.log(`Sending email to ${recipient} with subject '${subject}'...`);
    console.log(`Body: ${body}`);
    // Simulate async work
    await new Promise((resolve) => setTimeout(resolve, 50));
    console.log("Email sent.");
  }
}

export class ReportGenerator {
  public async generateReport(
    reportType: string,
    outputPath: string
  ): Promise<void> {
    console.log(`Generating ${reportType} report...`);
    console.log(`Saving report to ${outputPath}...`);
    // Simulate async work
    await new Promise((resolve) => setTimeout(resolve, 100));
    console.log("Report generated.");
  }
}

export class DatabaseService {
  public async runBackup(backupName: string): Promise<void> {
    console.log(`Starting database backup '${backupName}'...`);
    // Simulate async work
    await new Promise((resolve) => setTimeout(resolve, 150));
    console.log(`Database backup '${backupName}' completed.`);
  }
}

// --- Concrete Commands ---
export class SendEmailCommand implements Command {
  constructor(
    private service: EmailService,
    private recipient: string,
    private subject: string,
    private body: string
  ) {}

  public async execute(): Promise<void> {
    await this.service.sendEmail(this.recipient, this.subject, this.body);
  }

  // Optional: For logging/debugging purposes
  public getDescription(): string {
    return `SendEmailCommand(to: ${this.recipient}, subject: ${this.subject})`;
  }
}

export class GenerateReportCommand implements Command {
  constructor(
    private service: ReportGenerator,
    private reportType: string,
    private outputPath: string
  ) {}

  public async execute(): Promise<void> {
    await this.service.generateReport(this.reportType, this.outputPath);
  }

  public getDescription(): string {
    return `GenerateReportCommand(type: ${this.reportType}, path: ${this.outputPath})`;
  }
}

export class RunDatabaseBackupCommand implements Command {
  constructor(private service: DatabaseService, private backupName: string) {}

  public async execute(): Promise<void> {
    await this.service.runBackup(this.backupName);
  }

  public getDescription(): string {
    return `RunDatabaseBackupCommand(name: ${this.backupName})`;
  }
}
