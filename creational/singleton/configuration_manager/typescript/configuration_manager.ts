/**
 * Singleton Configuration Manager using a private static instance.
 *
 * This implementation relies on TypeScript's class features and a private
 * static member to hold the single instance.
 */
export class ConfigurationManager {
  private static instance: ConfigurationManager | null = null;
  private configData: { [key: string]: any } = {};
  private isLoaded: boolean = false;

  // Private constructor ensures no external instantiation
  private constructor() {
    this.loadConfig();
  }

  /**
   * Gets the single instance of the ConfigurationManager.
   * Creates the instance if it doesn't exist yet.
   */
  public static getInstance(): ConfigurationManager {
    if (ConfigurationManager.instance === null) {
      ConfigurationManager.instance = new ConfigurationManager();
    }
    return ConfigurationManager.instance;
  }

  /**
   * Simulates loading configuration data.
   * In a real app, this could read from a file, env vars, or a config service.
   */
  private loadConfig(): void {
    if (!this.isLoaded) {
      console.log("Loading configuration...");
      // Simulate loading data
      this.configData = {
        apiKey: "dummy_api_key_ts_456",
        serverUrl: "https://api.example-ts.com",
        featureFlagY: false,
      };
      this.isLoaded = true;
      console.log("Configuration loaded.");
    }
  }

  /**
   * Gets a configuration setting by key.
   * @param key - The configuration key.
   * @returns The value associated with the key, or undefined if not found.
   */
  public getSetting(key: string): any {
    return this.configData[key];
  }

  /**
   * Sets a configuration setting.
   * @param key - The configuration key.
   * @param value - The value to set.
   */
  public setSetting(key: string, value: any): void {
    this.configData[key] = value;
  }

  /**
   * Returns a copy of all configuration settings.
   */
  public getAllSettings(): { [key: string]: any } {
    // Return a copy to prevent external modification of the internal state
    return { ...this.configData };
  }
}
