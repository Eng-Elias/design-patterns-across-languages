from configuration_manager import ConfigurationManager

def main():
    """Demonstrates the usage of the Singleton ConfigurationManager."""
    print("--- Singleton Pattern - Configuration Manager Demo (Python) ---")

    # Get the singleton instance
    print("\nGetting Configuration Manager instance 1...")
    config_manager1 = ConfigurationManager()

    # Access configuration
    print(f"API Key (Instance 1): {config_manager1.get_setting('api_key')}")

    # Modify configuration through the first instance
    print("\nSetting 'timeout' via Instance 1...")
    config_manager1.set_setting("timeout", 30)
    print(f"Timeout (Instance 1): {config_manager1.get_setting('timeout')}")

    # Get the singleton instance again
    print("\nGetting Configuration Manager instance 2...")
    config_manager2 = ConfigurationManager()

    # Verify it's the same instance
    print(f"\nInstance 1 and Instance 2 are the same: {config_manager1 is config_manager2}")

    # Access the modified configuration through the second instance
    print(f"Timeout (Instance 2): {config_manager2.get_setting('timeout')}")
    print(f"API Key (Instance 2): {config_manager2.get_setting('api_key')}")

    print("\nAll settings:")
    print(config_manager2.get_all_settings())

if __name__ == "__main__":
    main()
