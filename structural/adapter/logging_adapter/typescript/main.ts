import {
  ThirdPartyLogger,
  LoggerAdapter,
  ApplicationService,
} from "./logging_adapter";

function main(): void {
  /**
   * Demonstrates using the Adapter pattern.
   */
  console.log("--- Using the Adapter for the Third-Party Logger ---");

  // Create the Adaptee (the incompatible logger)
  const thirdPartyLogger = new ThirdPartyLogger();

  // Create the Adapter, wrapping the Adaptee
  const loggerAdapter = new LoggerAdapter(thirdPartyLogger);

  // The client code (ApplicationService) uses the standard Logger interface
  // It doesn't know it's talking to an adapter or a third-party logger.
  const appService = new ApplicationService(loggerAdapter);

  console.log("\nPerforming operations:");
  appService.performOperation("ImportantData123");
  console.log("---");
  appService.performOperation("abc"); // Should trigger a warning
  console.log("---");
  appService.performOperation(""); // Should trigger an error
}

main();
