import {
  ThirdPartyLogger,
  LoggerAdapter,
  ApplicationService,
  Logger,
} from "./logging_adapter";

// Mock the Adaptee and Target interfaces using Jest
const mockAdaptee = {
  record: jest.fn(),
};

// We can create a more specific mock for the Logger interface if needed,
// but jest.fn() on individual methods works well here.
const mockLogger = {
  logInfo: jest.fn(),
  logWarning: jest.fn(),
  logError: jest.fn(),
};

describe("Logging Adapter Tests", () => {
  beforeEach(() => {
    // Reset mocks before each test
    mockAdaptee.record.mockClear();
    mockLogger.logInfo.mockClear();
    mockLogger.logWarning.mockClear();
    mockLogger.logError.mockClear();
  });

  test("adapter should call adaptee methods correctly", () => {
    // Arrange
    // Use the mock object directly here, no need to mock the constructor
    const adapter = new LoggerAdapter(
      mockAdaptee as unknown as ThirdPartyLogger
    );

    // Act
    adapter.logInfo("Info message");
    adapter.logWarning("Warning message");
    adapter.logError("Error message");

    // Assert
    expect(mockAdaptee.record).toHaveBeenCalledWith("info", "Info message");
    expect(mockAdaptee.record).toHaveBeenCalledWith(
      "warning",
      "Warning message"
    );
    expect(mockAdaptee.record).toHaveBeenCalledWith("error", "Error message");
    expect(mockAdaptee.record).toHaveBeenCalledTimes(3);
  });

  test("ApplicationService should use logger interface correctly", () => {
    // Arrange
    // Provide the mocked Logger interface to the service
    const appService = new ApplicationService(mockLogger as Logger);

    // --- Act & Assert: Successful operation ---
    appService.performOperation("Valid Data");

    expect(mockLogger.logInfo).toHaveBeenCalledWith(
      "Starting operation with data: Valid Data"
    );
    expect(mockLogger.logInfo).toHaveBeenCalledWith(
      "Operation completed successfully."
    );
    expect(mockLogger.logInfo).toHaveBeenCalledTimes(2);
    expect(mockLogger.logWarning).not.toHaveBeenCalled();
    expect(mockLogger.logError).not.toHaveBeenCalled();

    // --- Reset mocks for next scenario ---
    mockLogger.logInfo.mockClear();
    mockLogger.logWarning.mockClear();
    mockLogger.logError.mockClear();

    // --- Act & Assert: Operation with warning ---
    appService.performOperation("shrt");

    expect(mockLogger.logInfo).toHaveBeenCalledWith(
      "Starting operation with data: shrt"
    );
    expect(mockLogger.logWarning).toHaveBeenCalledWith(
      "Data 'shrt' is quite short."
    );
    expect(mockLogger.logInfo).toHaveBeenCalledWith(
      "Operation completed successfully."
    );
    expect(mockLogger.logInfo).toHaveBeenCalledTimes(2); // Start and End
    expect(mockLogger.logWarning).toHaveBeenCalledTimes(1);
    expect(mockLogger.logError).not.toHaveBeenCalled();

    // --- Reset mocks for next scenario ---
    mockLogger.logInfo.mockClear();
    mockLogger.logWarning.mockClear();
    mockLogger.logError.mockClear();

    // --- Act & Assert: Operation with error ---
    appService.performOperation("");

    expect(mockLogger.logInfo).toHaveBeenCalledWith(
      "Starting operation with data: "
    );
    expect(mockLogger.logError).toHaveBeenCalledWith(
      "Operation failed: Data cannot be empty"
    );
    expect(mockLogger.logInfo).toHaveBeenCalledTimes(1); // Only start message
    expect(mockLogger.logWarning).not.toHaveBeenCalled();
    expect(mockLogger.logError).toHaveBeenCalledTimes(1);
  });
});
