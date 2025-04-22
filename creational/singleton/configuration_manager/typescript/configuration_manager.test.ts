import { ConfigurationManager } from "./configuration_manager";

describe("ConfigurationManager Singleton", () => {
  // Note: Ideally, we'd reset the singleton state before each test
  // using a dedicated method like _resetInstanceForTesting().
  // Since adding that method via the tool failed, these tests assume
  // the state persists between them, which might not be ideal for true isolation.

  test("should return the same instance on multiple calls", () => {
    const instance1 = ConfigurationManager.getInstance();
    const instance2 = ConfigurationManager.getInstance();
    expect(instance1).toBe(instance2);
    // Check if a default value is loaded (proves initialization happened)
    expect(instance1.getSetting("apiKey")).toBe("dummy_api_key_ts_456");
  });

  test("should initialize configuration only once", () => {
    // Assuming getInstance was called in the previous test,
    // get the instance again and check if initial config persists.
    const instance1 = ConfigurationManager.getInstance();
    const initialApiKey = instance1.getSetting("apiKey");
    expect(initialApiKey).toBe("dummy_api_key_ts_456"); // Ensure it's loaded

    // Modify a setting
    instance1.setSetting("apiKey", "new_temporary_key");
    expect(instance1.getSetting("apiKey")).toBe("new_temporary_key");

    // Get instance again
    const instance2 = ConfigurationManager.getInstance();

    // Verify it's the same instance
    expect(instance1).toBe(instance2);

    // Verify configuration wasn't reloaded. The change should persist
    // because loadConfig should not run again.
    expect(instance2.getSetting("apiKey")).toBe("new_temporary_key");

    // Reset the value for subsequent tests if needed (though ideally reset happens in beforeEach)
    instance2.setSetting("apiKey", "dummy_api_key_ts_456");
  });

  test("should allow getting and setting configuration values", () => {
    const config = ConfigurationManager.getInstance();
    // Use a value known to be loaded initially
    const initialServerUrl = config.getSetting("serverUrl");
    expect(initialServerUrl).toBe("https://api.example-ts.com");

    config.setSetting("newSetting", "testValue");
    expect(config.getSetting("newSetting")).toBe("testValue");

    // Ensure setting doesn't affect unrelated settings
    expect(config.getSetting("serverUrl")).toBe(initialServerUrl);
  });

  test("getAllSettings should return a copy of the config", () => {
    const config = ConfigurationManager.getInstance();
    const settings1 = config.getAllSettings();
    settings1["newKey"] = "newValue"; // Modify the returned object

    const settings2 = config.getAllSettings();
    // The internal config should not have the 'newKey'
    expect(settings2["newKey"]).toBeUndefined();
    // The original returned object should have 'newKey'
    expect(settings1["newKey"]).toBe("newValue");
    // Ensure the objects themselves are different
    expect(settings1).not.toEqual(settings2);
  });
});
