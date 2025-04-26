import { performance } from "perf_hooks"; // For more precise timing if needed, but console.time works

// Define the structure for parameters
export type ExperimentParameters = Record<string, any>;

// Represents the simulation object that acts as a prototype
export class MolecularSimulation {
  public moleculeName: string;
  public parameters: ExperimentParameters;
  private _precomputedStates: number[] = []; // Simulate large dataset
  private _expensiveSetupDone: boolean = false;

  // Simplified private constructor - only initializes basic properties.
  // State (_precomputedStates, _expensiveSetupDone) is set by create or clone.
  private constructor(moleculeName: string, parameters: ExperimentParameters) {
    this.moleculeName = moleculeName;
    this.parameters = structuredClone(parameters); // Deep clone params early
  }

  // Static method to create the initial prototype instance, ensuring setup runs.
  public static async create(
    moleculeName: string,
    baseParameters: ExperimentParameters
  ): Promise<MolecularSimulation> {
    console.log(`🧬 Initializing simulation for '${moleculeName}'...`);
    // Call simplified constructor
    const instance = new MolecularSimulation(moleculeName, baseParameters);

    // Perform expensive setup - create should be the only entry point needing this.
    // No need to check _expensiveSetupDone here if create is used correctly (once).
    const states = await instance._performExpensiveSetup();

    // Assign the computed states and mark setup as done AFTER it completes.
    instance._precomputedStates = states;
    instance._expensiveSetupDone = true;

    return instance;
  }

  // Simulates a time-consuming setup process and RETURNS the computed states.
  // Removed internal check for _expensiveSetupDone.
  private async _performExpensiveSetup(): Promise<number[]> {
    console.log(
      `⏳ Performing expensive precomputation for '${this.moleculeName}' (takes ~2 seconds)...`
    );

    // Simulate async delay (like file I/O or network request)
    await new Promise((resolve) => setTimeout(resolve, 2000));

    // Simulate generating large data
    const size = 1_000_000;
    const states = Array.from({ length: size }, () => Math.random() * 100);

    // Don't set flags here, just return the data
    console.log(
      `✅ Expensive setup complete for '${this.moleculeName}'. ${states.length} states computed.`
    );
    return states;
  }

  // The Prototype pattern's core method: clone
  public clone(): MolecularSimulation {
    console.log(`\n🔄 Cloning simulation for '${this.moleculeName}'...`);

    // Use the simplified constructor
    const clone = new MolecularSimulation(this.moleculeName, this.parameters); // Constructor deep copies params

    // Manually copy the state from the prototype AFTER construction
    clone._precomputedStates = this._precomputedStates; // Share the large data reference
    clone._expensiveSetupDone = true; // Explicitly mark setup as done for the clone

    console.log(`    Cloned simulation created. Setup Skipped.`);
    return clone;
  }

  // Modify a specific parameter for this simulation instance
  public setParameter(key: string, value: any): void {
    console.log(
      `    Setting parameter '${key}' = ${JSON.stringify(value)} for '${
        this.moleculeName
      }' simulation`
    );
    this.parameters[key] = value;
  }

  // Runs the simulation using its current parameters and precomputed data
  public run(): void {
    console.log(
      `\n🔬 Running simulation for '${
        this.moleculeName
      }' with parameters: ${JSON.stringify(this.parameters)}`
    );
    if (!this._precomputedStates || this._precomputedStates.length === 0) {
      console.error("   ❌ Error: Precomputed states not available!");
      return;
    }

    // Simulate work using the data and parameters
    const temp = this.parameters["temperature"] ?? 298.15; // Kelvin
    const pressure = this.parameters["pressure"] ?? 1.0; // atm
    const duration = this.parameters["duration"] ?? 100; // picoseconds

    // Example calculation (placeholder)
    const resultMetric =
      ((this._precomputedStates.slice(0, 1000).reduce((a, b) => a + b, 0) *
        (temp / 273.15)) /
        pressure) *
      (duration / 10);
    console.log(
      `   Simulation complete. Result metric: ${resultMetric.toFixed(2)}`
    );
    console.log(
      `   (Used ${this._precomputedStates.length} precomputed states)`
    );
  }

  // Getter for testing
  public getPrecomputedStatesLength(): number {
    return this._precomputedStates.length;
  }
}
