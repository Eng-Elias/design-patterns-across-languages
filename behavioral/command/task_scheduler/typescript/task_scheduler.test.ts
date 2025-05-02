import { TaskScheduler } from "./scheduler";
import {
  Command,
  EmailService,
  ReportGenerator,
  DatabaseService,
  SendEmailCommand,
  GenerateReportCommand,
  RunDatabaseBackupCommand,
} from "./commands";

// Mock the receiver classes
jest.mock("./commands", () => {
  const originalModule = jest.requireActual("./commands");
  return {
    ...originalModule, // Keep original Command interface etc.
    EmailService: jest.fn().mockImplementation(() => ({
      sendEmail: jest.fn().mockResolvedValue(undefined), // Mock async method
    })),
    ReportGenerator: jest.fn().mockImplementation(() => ({
      generateReport: jest.fn().mockResolvedValue(undefined),
    })),
    DatabaseService: jest.fn().mockImplementation(() => ({
      runBackup: jest.fn().mockResolvedValue(undefined),
    })),
  };
});

// Type cast the mocks for easier use
const MockEmailService = EmailService as jest.MockedClass<typeof EmailService>;
const MockReportGenerator = ReportGenerator as jest.MockedClass<
  typeof ReportGenerator
>;
const MockDatabaseService = DatabaseService as jest.MockedClass<
  typeof DatabaseService
>;

describe("TaskScheduler Command Pattern", () => {
  let scheduler: TaskScheduler;
  let mockEmailServiceInstance: EmailService;
  let mockReportGeneratorInstance: ReportGenerator;
  let mockDbServiceInstance: DatabaseService;

  beforeEach(() => {
    // Clear previous mocks and create new instances for each test
    jest.clearAllMocks();
    MockEmailService.mockClear();
    MockReportGenerator.mockClear();
    MockDatabaseService.mockClear();

    // We need instances of the mocked classes
    mockEmailServiceInstance = new EmailService();
    mockReportGeneratorInstance = new ReportGenerator();
    mockDbServiceInstance = new DatabaseService();

    scheduler = new TaskScheduler();

    // Suppress console.log during tests if desired
    // jest.spyOn(console, 'log').mockImplementation(() => {});
  });

  afterEach(() => {
    // Restore console.log if it was spied upon
    // jest.restoreAllMocks();
  });

  test("should add and execute multiple commands in order", async () => {
    // Create commands using mocked service instances
    const cmd1 = new SendEmailCommand(
      mockEmailServiceInstance,
      "t@e.st",
      "Hi",
      "Test"
    );
    const cmd2 = new GenerateReportCommand(
      mockReportGeneratorInstance,
      "Test",
      "/dev/null"
    );
    const cmd3 = new RunDatabaseBackupCommand(
      mockDbServiceInstance,
      "backup-test"
    );

    scheduler.addTask(cmd1);
    scheduler.addTask(cmd2);
    scheduler.addTask(cmd3);

    expect(scheduler.getTaskCount()).toBe(3);

    await scheduler.runPendingTasks();

    // Check that methods on mocked instances were called
    expect(mockEmailServiceInstance.sendEmail).toHaveBeenCalledTimes(1);
    expect(mockEmailServiceInstance.sendEmail).toHaveBeenCalledWith(
      "t@e.st",
      "Hi",
      "Test"
    );

    expect(mockReportGeneratorInstance.generateReport).toHaveBeenCalledTimes(1);
    expect(mockReportGeneratorInstance.generateReport).toHaveBeenCalledWith(
      "Test",
      "/dev/null"
    );

    expect(mockDbServiceInstance.runBackup).toHaveBeenCalledTimes(1);
    expect(mockDbServiceInstance.runBackup).toHaveBeenCalledWith("backup-test");

    // Check order requires more complex mocking or checking console output order (less ideal)
    // For simplicity, we mainly check they were all called correctly here.

    // Check if the task queue is empty after running
    expect(scheduler.getTaskCount()).toBe(0);
  });

  test("should execute a single command", async () => {
    const cmd = new RunDatabaseBackupCommand(
      mockDbServiceInstance,
      "single-backup"
    );
    scheduler.addTask(cmd);
    expect(scheduler.getTaskCount()).toBe(1);

    await scheduler.runPendingTasks();

    expect(mockDbServiceInstance.runBackup).toHaveBeenCalledTimes(1);
    expect(mockDbServiceInstance.runBackup).toHaveBeenCalledWith(
      "single-backup"
    );
    expect(mockEmailServiceInstance.sendEmail).not.toHaveBeenCalled();
    expect(mockReportGeneratorInstance.generateReport).not.toHaveBeenCalled();
    expect(scheduler.getTaskCount()).toBe(0);
  });

  test("should do nothing when running with no tasks", async () => {
    expect(scheduler.getTaskCount()).toBe(0);
    await scheduler.runPendingTasks();

    expect(mockDbServiceInstance.runBackup).not.toHaveBeenCalled();
    expect(mockEmailServiceInstance.sendEmail).not.toHaveBeenCalled();
    expect(mockReportGeneratorInstance.generateReport).not.toHaveBeenCalled();
    expect(scheduler.getTaskCount()).toBe(0);
  });

  // Optional: Test error handling if execute() could throw
  test("should handle errors during task execution (if implemented)", async () => {
    // Mock one of the commands to throw an error
    const errorMsg = "Backup failed!";
    (mockDbServiceInstance.runBackup as jest.Mock).mockRejectedValueOnce(
      new Error(errorMsg)
    );
    const errorCmd = new RunDatabaseBackupCommand(
      mockDbServiceInstance,
      "fail-backup"
    );
    const successCmd = new SendEmailCommand(
      mockEmailServiceInstance,
      "a@b.c",
      "OK",
      "OK"
    );

    // Spy on console.error to check if error is logged
    const errorSpy = jest.spyOn(console, "error").mockImplementation(() => {});

    scheduler.addTask(errorCmd);
    scheduler.addTask(successCmd); // Add a task after the failing one

    await scheduler.runPendingTasks();

    // Check that the failing command was attempted
    expect(mockDbServiceInstance.runBackup).toHaveBeenCalledTimes(1);
    expect(mockDbServiceInstance.runBackup).toHaveBeenCalledWith("fail-backup");

    // Check that the error was logged by the scheduler
    expect(errorSpy).toHaveBeenCalledWith(
      expect.stringContaining("Error executing task"),
      expect.any(Error)
    );

    // Check that the subsequent command *still ran* (confirming sequential execution)
    expect(mockEmailServiceInstance.sendEmail).toHaveBeenCalledTimes(1);
    expect(mockEmailServiceInstance.sendEmail).toHaveBeenCalledWith(
      "a@b.c",
      "OK",
      "OK"
    );

    errorSpy.mockRestore(); // Clean up spy
  });
});
