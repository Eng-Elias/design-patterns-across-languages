import threading

class ConfigurationManager:
    """
    Singleton Configuration Manager using a lock for thread safety.

    This implementation uses a class variable to hold the single instance
    and a lock to ensure thread-safe initialization.
    """
    _instance = None
    _lock = threading.Lock()
    _config_data = {}

    def __new__(cls):
        # Double-checked locking pattern
        if cls._instance is None:
            with cls._lock:
                # Check again inside the lock
                if cls._instance is None:
                    cls._instance = super(ConfigurationManager, cls).__new__(cls)
                    # Initialize configuration (e.g., load from file/env)
                    cls._instance._load_config()
        return cls._instance

    def _load_config(self):
        """Simulates loading configuration data."""
        print("Loading configuration...")
        # In a real app, load from a file, database, or environment variables
        self._config_data = {
            "api_key": "dummy_api_key_123",
            "database_url": "localhost:5432",
            "feature_flag_x": True
        }
        print("Configuration loaded.")

    def get_setting(self, key):
        """Get a configuration setting by key."""
        return self._config_data.get(key)

    def set_setting(self, key, value):
        """Set a configuration setting."""
        self._config_data[key] = value

    def get_all_settings(self):
        """Return all configuration settings."""
        return self._config_data.copy()
