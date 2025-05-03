import { EventSource } from "./event_source";
import { LoggerObserver, NotifierObserver } from "./observer";
import { EventType } from "./event";

/**
 * Utility function to simulate delay.
 */
function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/**
 * Main function to demonstrate the Observer pattern.
 */
async function main() {
  console.log("--- TypeScript Observer Pattern Demo ---");

  const eventSource = new EventSource();
  const logger = new LoggerObserver("FileLogger"); // Provide name
  const notifier = new NotifierObserver("EmailNotifier"); // Provide name

  console.log("\nAttaching observers...");
  eventSource.attach(logger);
  eventSource.attach(notifier);

  console.log("\nGenerating specific events...");
  eventSource.generateEvent(EventType.LogInfo, { msg: "User logged in" });
  await delay(50);
  eventSource.generateEvent(EventType.LogWarn, { msg: "Disk space low" });
  await delay(50);
  eventSource.generateEvent(EventType.LogError, {
    code: 500,
    error: "Database connection failed",
  });
  await delay(50);
  eventSource.generateEvent(EventType.LogCritical, {
    code: 999,
    error: "System meltdown imminent",
  });
  await delay(50);

  // Detach the NotifierObserver (consistent with Python/Go examples)
  console.log("\nDetaching EmailNotifier..."); // Log based on name
  eventSource.detach(notifier);

  console.log("\nGenerating another event...");
  eventSource.generateEvent(EventType.LogInfo, { msg: "User logged out" });
  await delay(50);

  console.log("\n--- Demo Finished ---");
}

// Run the main function
main().catch((error) => {
  console.error("An error occurred:", error);
});
