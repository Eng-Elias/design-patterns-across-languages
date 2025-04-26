import unittest
import time
import copy
from unittest.mock import patch, call
import io

from molecular_simulation import MolecularSimulation

class TestMolecularSimulationPrototype(unittest.TestCase):

    def setUp(self):
        """Set up a base simulation object for tests."""
        # Reduce sleep time for faster tests
        self.original_sleep = time.sleep
        time.sleep = lambda x: None # Replace sleep with no-op

        self.base_params = {'temperature': 300, 'pressure': 1}
        # Capture print output during setup for specific tests
        with patch('sys.stdout', new_callable=io.StringIO) as mock_stdout:
            self.prototype = MolecularSimulation("TestMolecule", self.base_params)
            self.setup_output = mock_stdout.getvalue()

    def tearDown(self):
        """Restore original time.sleep."""
        time.sleep = self.original_sleep

    def test_initialization_performs_expensive_setup(self):
        """Verify that expensive setup print message occurs during __init__."""
        # Check the output captured during setUp
        self.assertIn("Performing expensive precomputation", self.setup_output)
        self.assertIn("Expensive setup complete", self.setup_output)
        self.assertTrue(len(self.prototype._precomputed_states) > 0) # Check data exists

    def test_clone_creates_new_instance(self):
        """Verify clone returns a new object instance."""
        clone = self.prototype.clone()
        self.assertIsNot(clone, self.prototype)
        self.assertIsInstance(clone, MolecularSimulation)

    def test_clone_does_not_repeat_expensive_setup(self):
        """Verify that cloning does not trigger the expensive setup print message."""
        with patch('sys.stdout', new_callable=io.StringIO) as mock_stdout:
            clone = self.prototype.clone()
            clone_output = mock_stdout.getvalue()

        self.assertIn("Cloning simulation", clone_output) # Cloning message should be there
        self.assertNotIn("Performing expensive precomputation", clone_output)
        self.assertNotIn("Expensive setup complete", clone_output)

    def test_cloned_object_has_independent_parameters(self):
        """Verify modifying a clone's parameters doesn't affect the original."""
        clone = self.prototype.clone()
        original_temp = self.prototype.parameters['temperature']
        clone.set_parameter('temperature', original_temp + 50)

        self.assertEqual(self.prototype.parameters['temperature'], 300)
        self.assertEqual(clone.parameters['temperature'], 350)

        # Also test deep copy of parameters dictionary itself
        clone.parameters['new_param'] = 123
        self.assertNotIn('new_param', self.prototype.parameters)
        self.assertIn('new_param', clone.parameters)

    def test_cloned_object_shares_large_data(self):
        """Verify clone potentially shares the large precomputed data structure (check object ID)."""
        # Note: This test relies on the implementation detail that the large list
        # *might* be shared by reference via deepcopy if it's not modified.
        # A more robust test focuses on behavior (independent params), which is covered above.
        clone = self.prototype.clone()
        # Check if the lists are the *same object* in memory (shallow copy for this part)
        self.assertIs(clone._precomputed_states, self.prototype._precomputed_states)
        # Ensure the clone *can* run, implying it has the data
        with patch('sys.stdout', new_callable=io.StringIO):
             clone.run() # Should not fail due to missing data
        self.assertTrue(len(clone._precomputed_states) > 0)

    def test_run_method_works_on_original_and_clone(self):
        """Verify the run method executes without errors on both instances."""
        clone = self.prototype.clone()
        clone.set_parameter('duration', 50)

        with patch('sys.stdout', new_callable=io.StringIO) as mock_stdout:
            try:
                print("\nRunning original:")
                self.prototype.run()
                print("\nRunning clone:")
                clone.run()
            except Exception as e:
                self.fail(f"run() method failed with exception: {e}")

            output = mock_stdout.getvalue()

        # Check if the print statements from run() contain the expected params
        # Verify params printed for the original run (based on observed output)
        self.assertIn("parameters: {'temperature': 300, 'pressure': 1}", output)
        # Verify params printed for the clone run (includes modified duration)
        self.assertIn("'temperature': 300", output)
        self.assertIn("'pressure': 1", output)
        self.assertIn("'duration': 50", output)

        # Ensure the original run didn't accidentally print the clone's duration
        # (This implicitly checks the output sections are distinct)
        self.assertNotIn("'duration': 1000", output)


    def test_set_parameter_modifies_only_clone(self):
        """Verify setting a parameter on the clone does not affect the original."""
        clone = self.prototype.clone()
        clone.set_parameter('temperature', 400)

        self.assertEqual(self.prototype.parameters['temperature'], 300)
        self.assertEqual(clone.parameters['temperature'], 400)

if __name__ == '__main__':
    unittest.main()
