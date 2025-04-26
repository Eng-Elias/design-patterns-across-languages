import time
import copy
import random
from typing import List, Dict, Any

# Define a simple parameter structure (can be more complex)
ExperimentParameters = Dict[str, Any]

class MolecularSimulation:
    """Represents a potentially expensive simulation setup (Concrete Prototype)."""

    def __init__(self, molecule_name: str, base_parameters: ExperimentParameters):
        """Initializes the simulation, including the expensive setup."""
        self.molecule_name = molecule_name
        self.parameters = copy.deepcopy(base_parameters) # Ensure base params are copied
        self._precomputed_states: List[float] = []

        print(f"🧬 Initializing simulation for '{self.molecule_name}'...")
        self._perform_expensive_setup()

    def _perform_expensive_setup(self):
        """Simulates a time-consuming setup process."""
        print(f"⏳ Performing expensive precomputation for '{self.molecule_name}' (takes ~2 seconds)...")
        # Simulate delay
        time.sleep(2)
        # Simulate generating large data
        self._precomputed_states = [random.random() * 100 for _ in range(1_000_000)] # Simulate large data
        print(f"✅ Expensive setup complete for '{self.molecule_name}'. {len(self._precomputed_states)} states computed.")

    def clone(self) -> 'MolecularSimulation':
        """Creates a clone, sharing expensive data and deep copying parameters."""
        print(f"\n🔄 Cloning simulation for '{self.molecule_name}'...")

        # 1. Create a new instance without calling __init__ (to skip expensive setup)
        # Using __new__ is one way, another is a dedicated internal method if needed.
        new_simulation = self.__class__.__new__(self.__class__)

        # 2. Copy essential attributes
        new_simulation.molecule_name = self.molecule_name
        # Deep copy parameters to ensure independence
        new_simulation.parameters = copy.deepcopy(self.parameters)

        # 3. Share the reference to the expensive precomputed data
        new_simulation._precomputed_states = self._precomputed_states

        # Note: Python doesn't explicitly need the '_expensiveSetupDone' flag here
        # because skipping __init__ via __new__ prevents the setup call.

        print(f"    Cloned simulation created. Setup Skipped.")
        return new_simulation

    def set_parameter(self, key: str, value: Any):
        """Modify a specific parameter for this simulation instance."""
        print(f"    Setting parameter '{key}' = {value} for '{self.molecule_name}' simulation")
        self.parameters[key] = value

    def run(self):
        """Runs the simulation using its current parameters and precomputed data."""
        print(f"\n🔬 Running simulation for '{self.molecule_name}' with parameters: {self.parameters}")
        if not self._precomputed_states:
            print("   ❌ Error: Precomputed states not available!")
            return

        # Simulate work using the data and parameters
        temp = self.parameters.get('temperature', 298.15) # Kelvin
        pressure = self.parameters.get('pressure', 1.0) # atm
        duration = self.parameters.get('duration', 100) # picoseconds

        # Example calculation using precomputed data and params
        # (This is just a placeholder for actual simulation logic)
        result_metric = sum(self._precomputed_states[:1000]) * (temp / 273.15) / pressure * (duration / 10)
        print(f"   Simulation complete. Result metric: {result_metric:.2f}")
        print(f"   (Used {len(self._precomputed_states)} precomputed states)")
