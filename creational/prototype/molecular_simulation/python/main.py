from molecular_simulation import MolecularSimulation
import time

def main():
    print("--- Prototype Pattern - Molecular Simulation Demo (Python) ---")

    start_time = time.time()

    # 1. Create the initial prototype instance (expensive setup)
    base_params = {
        'temperature': 298.15, # Kelvin
        'pressure': 1.0,      # atm
        'duration': 1000      # picoseconds
    }
    print("\nStep 1: Creating the initial prototype simulation (will perform expensive setup)...")
    prototype_simulation = MolecularSimulation("Water (H2O)", base_params)
    setup_duration = time.time() - start_time
    print(f"Initial setup took {setup_duration:.2f} seconds.")

    # 2. Clone the prototype to create variations (cheap)
    cloning_start_time = time.time()
    print("\nStep 2: Cloning the prototype to create variations...")

    # Clone 1: Higher temperature
    sim_high_temp = prototype_simulation.clone()
    sim_high_temp.set_parameter('temperature', 350.0)

    # Clone 2: Longer duration
    sim_long_duration = prototype_simulation.clone()
    sim_long_duration.set_parameter('duration', 5000)

    # Clone 3: Different pressure
    sim_high_pressure = prototype_simulation.clone()
    sim_high_pressure.set_parameter('pressure', 5.0)
    sim_high_pressure.set_parameter('temperature', 310.0) # Can set multiple

    cloning_duration = time.time() - cloning_start_time
    print(f"\nCloning {3} simulations took {cloning_duration:.4f} seconds (should be very fast).")

    # 3. Run all simulations
    print("\nStep 3: Running the base simulation and all clones...")
    simulations_to_run = [
        prototype_simulation,
        sim_high_temp,
        sim_long_duration,
        sim_high_pressure
    ]

    run_start_time = time.time()
    for i, sim in enumerate(simulations_to_run):
        print(f"\n--- Running Simulation {i+1}/{len(simulations_to_run)} --- ({sim.molecule_name}) ")
        sim.run()

    run_duration = time.time() - run_start_time
    print(f"\nRunning {len(simulations_to_run)} simulations took {run_duration:.2f} seconds.")

    total_duration = time.time() - start_time
    print(f"\n--- Demo Complete ---")
    print(f"Total time: {total_duration:.2f} seconds (Initial Setup: {setup_duration:.2f}s, Cloning: {cloning_duration:.4f}s, Running: {run_duration:.2f}s)")

if __name__ == "__main__":
    main()
