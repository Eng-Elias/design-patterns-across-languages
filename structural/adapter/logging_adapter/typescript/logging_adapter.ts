// --- Adaptee --- (The incompatible third-party library)
export class ThirdPartyLogger {
  /**
   * A hypothetical third-party logger with an incompatible interface.
   */
  public record(severity: string, message: string): void {
    console.log(`[3rdPartyLogger - ${severity.toUpperCase()}]: ${message}`);
  }
}

// --- Target --- (The interface the client code expects)
export interface Logger {
  /**
   * The interface our application expects for logging.
   */
  logInfo(message: string): void;
  logWarning(message: string): void;
  logError(message: string): void;
}

// --- Adapter ---
export class LoggerAdapter implements Logger {
  /**
   * Adapts the ThirdPartyLogger to the Logger interface.
   */
  private adaptee: ThirdPartyLogger;

  constructor(adaptee: ThirdPartyLogger) {
    this.adaptee = adaptee;
  }

  public logInfo(message: string): void {
    this.adaptee.record("info", message);
  }

  public logWarning(message: string): void {
    this.adaptee.record("warning", message);
  }

  public logError(message: string): void {
    this.adaptee.record("error", message);
  }
}

// --- Client Code ---
export class ApplicationService {
  /**
   * A service that uses the Logger interface.
   */
  private logger: Logger;

  constructor(logger: Logger) {
    this.logger = logger;
  }

  public performOperation(data: string): void {
    this.logger.logInfo(`Starting operation with data: ${data}`);
    try {
      // Simulate an operation
      if (!data) {
        throw new Error("Data cannot be empty");
      }
      if (data.length < 5) {
        this.logger.logWarning(`Data '${data}' is quite short.`);
      }
      // ... perform actual operation ...
      this.logger.logInfo("Operation completed successfully.");
    } catch (e: unknown) {
      if (e instanceof Error) {
        this.logger.logError(`Operation failed: ${e.message}`);
      } else {
        this.logger.logError(`Operation failed with unknown error type`);
      }
    }
  }
}
