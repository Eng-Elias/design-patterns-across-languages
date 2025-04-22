import unittest
import threading
from configuration_manager import ConfigurationManager

class TestConfigurationManager(unittest.TestCase):

    def test_singleton_instance(self):
        """Test that multiple calls return the same instance."""
        instance1 = ConfigurationManager()
        instance2 = ConfigurationManager()
        self.assertIs(instance1, instance2, "ConfigurationManager should return the same instance.")

    def test_config_access(self):
        """Test accessing and modifying configuration."""
        config = ConfigurationManager()
        self.assertEqual(config.get_setting("database_url"), "localhost:5432")
        config.set_setting("new_setting", "test_value")
        self.assertEqual(config.get_setting("new_setting"), "test_value")

    def test_thread_safety(self):
        """Test thread safety of singleton initialization."""
        instances = []
        num_threads = 10

        def get_instance():
            instance = ConfigurationManager()
            instances.append(instance)

        threads = [threading.Thread(target=get_instance) for _ in range(num_threads)]
        for t in threads:
            t.start()
        for t in threads:
            t.join()

        first_instance = instances[0]
        for i in range(1, num_threads):
            self.assertIs(first_instance, instances[i], f"Instance {i} is not the same as the first instance.")

if __name__ == '__main__':
    unittest.main()
