import { ConfigurationManager } from "./configuration_manager";

function main(): void {
  console.log(
    "--- Singleton Pattern - Configuration Manager Demo (TypeScript) ---"
  );

  // Get the singleton instance
  console.log("\nGetting Configuration Manager instance 1...");
  const configManager1 = ConfigurationManager.getInstance();

  // Access configuration
  console.log(`API Key (Instance 1): ${configManager1.getSetting("apiKey")}`);
  console.log(
    `Server URL (Instance 1): ${configManager1.getSetting("serverUrl")}`
  );

  // Modify configuration through the first instance
  console.log("\nSetting 'retryAttempts' via Instance 1...");
  configManager1.setSetting("retryAttempts", 3);
  console.log(
    `Retry Attempts (Instance 1): ${configManager1.getSetting("retryAttempts")}`
  );

  // Get the singleton instance again
  console.log("\nGetting Configuration Manager instance 2...");
  const configManager2 = ConfigurationManager.getInstance();

  // Verify it's the same instance
  console.log(
    `\nInstance 1 and Instance 2 are the same: ${
      configManager1 === configManager2
    }`
  );

  // Access the modified configuration through the second instance
  console.log(
    `Retry Attempts (Instance 2): ${configManager2.getSetting("retryAttempts")}`
  );
  console.log(`API Key (Instance 2): ${configManager2.getSetting("apiKey")}`);

  console.log("\nAll settings:");
  console.log(configManager2.getAllSettings());
}

// Run the demonstration
main();
