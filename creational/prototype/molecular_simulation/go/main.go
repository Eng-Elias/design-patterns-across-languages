package main

import (
	"fmt"
	"log"
	"prototype_pattern_molecular_simulation_go/molecular_simulation"
	"time"
)

func main() {
	fmt.Println("--- Prototype Pattern: Molecular Simulation Demo (Go) ---")

	startTime := time.Now()

	// 1. Create the initial prototype instance (expensive setup)
	baseParams := molecular_simulation.ExperimentParameters{
		"temperature": 298.15, // Kelvin (use float64 for Go)
		"pressure":    1.0,    // atm
		"duration":    1000.0, // picoseconds
	}
	fmt.Println("\nStep 1: Creating the initial prototype simulation (will perform expensive setup)...")
	prototypeSimulation, err := molecular_simulation.NewMolecularSimulation("Water (H2O)", baseParams)
	if err != nil {
		log.Fatalf("Failed to create prototype simulation: %v", err)
	}
	setupDuration := time.Since(startTime)
	fmt.Printf("Initial setup took %s.\n", setupDuration)

	// 2. Clone the prototype to create variations (cheap)
	cloningStartTime := time.Now()
	fmt.Println("\nStep 2: Cloning the prototype to create variations...")

	// Clone 1: Higher temperature
	simHighTemp := prototypeSimulation.Clone()
	simHighTemp.SetParameter("temperature", 350.0)

	// Clone 2: Longer duration
	simLongDuration := prototypeSimulation.Clone()
	simLongDuration.SetParameter("duration", 5000.0)

	// Clone 3: Different pressure
	simHighPressure := prototypeSimulation.Clone()
	simHighPressure.SetParameter("pressure", 5.0)
	simHighPressure.SetParameter("temperature", 310.0) // Can set multiple

	cloningDuration := time.Since(cloningStartTime)
	fmt.Printf("\nCloning %d simulations took %s (should be very fast).\n", 3, cloningDuration)

	// 3. Run all simulations
	fmt.Println("\nStep 3: Running the base simulation and all clones...")
	simulationsToRun := []*molecular_simulation.MolecularSimulation{
		prototypeSimulation,
		simHighTemp,
		simLongDuration,
		simHighPressure,
	}

	runStartTime := time.Now()
	for i, sim := range simulationsToRun {
		fmt.Printf("\n--- Running Simulation %d/%d --- (%s) \n", i+1, len(simulationsToRun), sim.MoleculeName)
		sim.Run()
	}

	runDuration := time.Since(runStartTime)
	fmt.Printf("\nRunning %d simulations took %s.\n", len(simulationsToRun), runDuration)

	totalDuration := time.Since(startTime)
	fmt.Println("\n--- Demo Complete ---")
	fmt.Printf("Total time: %s (Initial Setup: %s, Cloning: %s, Running: %s)\n",
		totalDuration, setupDuration, cloningDuration, runDuration)
}
