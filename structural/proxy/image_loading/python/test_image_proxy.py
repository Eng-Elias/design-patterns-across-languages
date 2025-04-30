import unittest
from unittest.mock import patch, MagicMock, call
import time

# Assuming image_proxy.py is in the same directory or accessible via PYTHONPATH
from image_proxy import RealImage, ProxyImage

class TestImageProxy(unittest.TestCase):

    def test_proxy_initialization_does_not_load_real_image(self):
        """Verify that creating a ProxyImage doesn't immediately load the RealImage."""
        filename = "test_init.jpg"
        # Patch __init__ of RealImage to check if it's called during ProxyImage creation
        with patch('image_proxy.RealImage.__init__', return_value=None) as mock_real_init:
            proxy = ProxyImage(filename)
            self.assertIsNone(proxy._real_image) # Internal check, usually avoid if possible
            self.assertFalse(proxy.is_loaded()) # Use public API if available
            mock_real_init.assert_not_called()
            self.assertEqual(proxy.get_filename(), filename)

    def test_proxy_loads_real_image_on_first_display(self):
        """Verify that the RealImage is loaded only when display() is first called."""
        filename = "test_first_display.png"
        proxy = ProxyImage(filename)

        # Mock the RealImage and its methods to track calls
        # We patch the *class* RealImage within the image_proxy module where ProxyImage looks for it.
        with patch('image_proxy.RealImage') as MockRealImageClass:
            # Configure the mock instance that will be returned by RealImage(filename)
            mock_real_instance = MockRealImageClass.return_value
            mock_real_instance.display = MagicMock() # Mock the display method
            mock_real_instance.get_filename.return_value = filename # Mock get_filename

            # Initial state: not loaded
            self.assertFalse(proxy.is_loaded())
            MockRealImageClass.assert_not_called()

            # Call display() for the first time
            proxy.display()

            # Assertions after first display:
            # 1. RealImage constructor should have been called once
            MockRealImageClass.assert_called_once_with(filename)
            # 2. The proxy should now hold the mock real image instance
            self.assertIsNotNone(proxy._real_image) # Check internal state
            self.assertTrue(proxy.is_loaded())
            self.assertIs(proxy._real_image, mock_real_instance)
            # 3. The real image's display method should have been called once
            mock_real_instance.display.assert_called_once()

    def test_proxy_reuses_loaded_real_image_on_subsequent_displays(self):
        """Verify that subsequent calls to display() reuse the loaded RealImage."""
        filename = "test_reuse.gif"
        proxy = ProxyImage(filename)

        with patch('image_proxy.RealImage') as MockRealImageClass:
            mock_real_instance = MockRealImageClass.return_value
            mock_real_instance.display = MagicMock()

            # First call to display (loads the image)
            proxy.display()

            # Reset mocks to check calls during the *second* display
            MockRealImageClass.reset_mock()
            mock_real_instance.display.reset_mock()

            # Second call to display
            proxy.display()

            # Assertions after second display:
            # 1. RealImage constructor should NOT have been called again
            MockRealImageClass.assert_not_called()
            # 2. The real image's display method should have been called once (during this second call)
            mock_real_instance.display.assert_called_once()
            # 3. Check loaded status remains true
            self.assertTrue(proxy.is_loaded())

    def test_proxy_get_filename_does_not_load_real_image(self):
        """Verify calling get_filename() on the proxy doesn't trigger loading."""
        filename = "test_get_filename.bmp"

        with patch('image_proxy.RealImage.__init__', return_value=None) as mock_real_init:
            proxy = ProxyImage(filename)
            # Call get_filename()
            retrieved_filename = proxy.get_filename()

            # Assertions:
            self.assertEqual(retrieved_filename, filename)
            # RealImage constructor should not have been called
            mock_real_init.assert_not_called()
            # Proxy should still not be loaded
            self.assertFalse(proxy.is_loaded())

    # Test the RealImage loading simulation (optional, less critical for Proxy pattern test)
    def test_real_image_loading_simulation(self):
        """Check the simulated loading message and delay in RealImage."""
        filename = "test_real_load.tiff"

        # Patch time.sleep to avoid actual delay during test
        with patch('time.sleep') as mock_sleep:
            # Capture print output
            with patch('builtins.print') as mock_print:
                real_image = RealImage(filename)

                # Assertions:
                # 1. time.sleep was called
                mock_sleep.assert_called_once()
                # 2. Check print calls for loading messages
                expected_calls = [
                    call(f"Loading image: '{filename}' from disk... (Simulating delay)"),
                    call(f"Finished loading image: '{filename}'")
                ]
                mock_print.assert_has_calls(expected_calls, any_order=False)

                # Check display message
                mock_print.reset_mock()
                real_image.display()
                mock_print.assert_called_once_with(f"Displaying image: '{filename}'")
                self.assertEqual(real_image.get_filename(), filename)

if __name__ == '__main__':
    unittest.main()
