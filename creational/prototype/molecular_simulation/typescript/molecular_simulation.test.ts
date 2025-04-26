import {
  MolecularSimulation,
  ExperimentParameters,
} from "./molecular_simulation";

// Mock setTimeout to avoid actual delays during tests
jest.useFakeTimers();

// Mock console methods to capture output and prevent clutter
const consoleLogSpy = jest.spyOn(console, "log").mockImplementation();
const consoleWarnSpy = jest.spyOn(console, "warn").mockImplementation();
const consoleErrorSpy = jest.spyOn(console, "error").mockImplementation();

describe("MolecularSimulation Prototype", () => {
  let prototype: MolecularSimulation;
  const baseParams: ExperimentParameters = { temperature: 300, pressure: 1 };
  const moleculeName = "TestMoleculeTS";

  // Use beforeAll to create the prototype once, as setup is expensive
  beforeAll(async () => {
    consoleLogSpy.mockClear(); // Clear logs before the initial creation

    // 1. Initiate the async creation, store the promise
    const creationPromise = MolecularSimulation.create(
      moleculeName,
      baseParams
    );

    // 2. Fast-forward timers to complete the simulated async setup (2000ms)
    // This needs to happen *before* awaiting the promise that depends on the timer.
    jest.advanceTimersByTime(2100); // Ensure enough time passes

    // 3. Now await the promise, which should resolve quickly
    prototype = await creationPromise;
  });

  beforeEach(() => {
    // Clear mocks before each individual test
    consoleLogSpy.mockClear();
    consoleWarnSpy.mockClear();
    consoleErrorSpy.mockClear();
  });

  test("should perform expensive setup during initial creation", () => {
    // Note: The actual check happens implicitly in beforeAll.
    // We verify its results here.
    expect(prototype).toBeDefined();
    expect(prototype.moleculeName).toBe(moleculeName);
    // Check if console.log was called with setup messages during beforeAll
    // This requires inspecting the logs *before* the beforeEach clear.
    // Jest spies don't easily retain calls across clears, so we rely on state.
    expect(prototype.getPrecomputedStatesLength()).toBeGreaterThan(0);
    // We can check the initial logs if needed by spying before beforeAll
  });

  test("clone() should create a new instance", () => {
    const clone = prototype.clone();
    expect(clone).toBeInstanceOf(MolecularSimulation);
    expect(clone).not.toBe(prototype); // Should be a different object reference
  });

  test("clone() should not repeat expensive setup", async () => {
    const clone = prototype.clone();
    // Ensure the clone has the data immediately without waiting
    expect(clone.getPrecomputedStatesLength()).toBeGreaterThan(0);

    // Check console output during cloning
    const logOutput = consoleLogSpy.mock.calls.flat().join("\n");
    expect(logOutput).toContain(`Cloning simulation for '${moleculeName}'`);
    expect(logOutput).toContain("Setup Skipped");
    expect(logOutput).not.toContain("Performing expensive precomputation");
    expect(logOutput).not.toContain("Expensive setup complete");

    // Explicitly check the internal flag if accessible (or rely on behavior)
    expect((clone as any)._expensiveSetupDone).toBe(true);
  });

  test("cloned object should have independent parameters", () => {
    const clone = prototype.clone();
    const originalTemp = prototype.parameters["temperature"];
    const newTemp = originalTemp + 50;

    clone.setParameter("temperature", newTemp);
    clone.parameters["newParam"] = "cloneOnly"; // Modify directly for test

    // Original should be unchanged
    expect(prototype.parameters["temperature"]).toBe(originalTemp);
    expect(prototype.parameters["newParam"]).toBeUndefined();

    // Clone should have the new values
    expect(clone.parameters["temperature"]).toBe(newTemp);
    expect(clone.parameters["newParam"]).toBe("cloneOnly");
  });

  test("cloned object should have necessary data for execution", () => {
    const clone = prototype.clone();
    // The clone should have the precomputed data
    expect(clone.getPrecomputedStatesLength()).toBeGreaterThan(0);
    expect(clone.getPrecomputedStatesLength()).toEqual(
      prototype.getPrecomputedStatesLength()
    );

    // Attempt to run the clone - should not error due to missing data
    expect(() => clone.run()).not.toThrow();
    expect(consoleErrorSpy).not.toHaveBeenCalled();
  });

  test("run() method should execute on original and clone", () => {
    const clone = prototype.clone();
    clone.setParameter("duration", 50);

    // Run original
    prototype.run();
    const originalLogOutput = consoleLogSpy.mock.calls
      .slice(-3)
      .flat()
      .join("\n"); // Get last few log lines
    expect(originalLogOutput).toContain(
      `Running simulation for '${moleculeName}'`
    );
    expect(originalLogOutput).toContain("Simulation complete.");
    expect(originalLogOutput).toContain(
      `"temperature":${baseParams.temperature}`
    );
    expect(originalLogOutput).toContain(`"pressure":${baseParams.pressure}`);
    expect(originalLogOutput).not.toContain(`"duration":`); // Ensure duration is NOT logged for original

    consoleLogSpy.mockClear(); // Clear before running clone

    // Run clone
    clone.run();
    const cloneLogOutput = consoleLogSpy.mock.calls.flat().join("\n");
    expect(cloneLogOutput).toContain(
      `Running simulation for '${moleculeName}'`
    );
    expect(cloneLogOutput).toContain("Simulation complete.");
    expect(cloneLogOutput).toContain(`"duration":50`); // Check modified duration

    expect(consoleErrorSpy).not.toHaveBeenCalled();
  });
});
