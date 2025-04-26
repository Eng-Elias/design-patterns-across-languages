import { MolecularSimulation } from "./molecular_simulation";

console.log("[main.ts] Script execution started."); // Added for debugging

async function main() {
  console.log("[main.ts] Async main function started."); // Added for debugging
  console.log(
    "--- Prototype Pattern: Molecular Simulation Demo (TypeScript) ---"
  );

  console.time("Total Demo Time");

  // 1. Create the initial prototype instance (expensive setup)
  const baseParams = {
    temperature: 298.15, // Kelvin
    pressure: 1.0, // atm
    duration: 1000, // picoseconds
  };
  console.log(
    "\nStep 1: Creating the initial prototype simulation (will perform expensive setup)..."
  );
  console.time("Initial Setup");
  const prototypeSimulation = await MolecularSimulation.create(
    "Water (H2O)",
    baseParams
  );
  console.timeEnd("Initial Setup");

  // 2. Clone the prototype to create variations (cheap)
  console.log("\nStep 2: Cloning the prototype to create variations...");
  console.time("Cloning Phase");

  // Clone 1: Higher temperature
  const simHighTemp = prototypeSimulation.clone();
  simHighTemp.setParameter("temperature", 350.0); // Boiling point

  // Clone 2: Longer duration
  const simLongDuration = prototypeSimulation.clone();
  simLongDuration.setParameter("duration", 5000);

  // Clone 3: Different pressure
  const simHighPressure = prototypeSimulation.clone();
  simHighPressure.setParameter("pressure", 10.0);
  simHighPressure.setParameter("temperature", 310.0);

  console.timeEnd("Cloning Phase");
  console.log(`(Cloning ${3} simulations was very fast)`);

  // 3. Run all simulations
  console.log("\nStep 3: Running the base simulation and all clones...");
  const simulationsToRun = [
    prototypeSimulation,
    simHighTemp,
    simLongDuration,
    simHighPressure,
  ];

  console.time("Running Phase");
  for (let i = 0; i < simulationsToRun.length; i++) {
    const sim = simulationsToRun[i];
    console.log(
      `\n--- Running Simulation ${i + 1}/${simulationsToRun.length} --- (${
        sim.moleculeName
      }) `
    );
    sim.run();
  }
  console.timeEnd("Running Phase");

  console.log("\n--- Demo Complete ---");
  console.timeEnd("Total Demo Time");
}

// Run the async main function
main().catch((error) => {
  console.error("An error occurred:", error);
});
