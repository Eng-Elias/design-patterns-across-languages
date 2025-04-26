package molecular_simulation

import (
	"fmt"
	"math/rand"
	"time"
)

// ExperimentParameters defines the type for simulation parameters
type ExperimentParameters map[string]interface{}

// MolecularSimulation represents the prototype object
type MolecularSimulation struct {
	MoleculeName       string
	Parameters         ExperimentParameters
	precomputedStates  []float64 // Slice to hold large data
	expensiveSetupDone bool      // Flag to prevent re-running setup
}

// NewMolecularSimulation creates a new simulation instance and runs the expensive setup.
func NewMolecularSimulation(moleculeName string, baseParameters ExperimentParameters) (*MolecularSimulation, error) {
	fmt.Printf("🧬 Initializing simulation for '%s'...\n", moleculeName)

	// Deep copy initial parameters to prevent external modification
	copiedParams := make(ExperimentParameters)
	for k, v := range baseParameters {
		copiedParams[k] = v // Simple copy, assumes non-nested maps/slices for params
	}

	sim := &MolecularSimulation{
		MoleculeName:       moleculeName,
		Parameters:         copiedParams,
		precomputedStates:  nil, // Initialize as nil
		expensiveSetupDone: false,
	}

	err := sim.performExpensiveSetup()
	if err != nil {
		return nil, fmt.Errorf("failed during expensive setup: %w", err)
	}

	return sim, nil
}

// performExpensiveSetup simulates the time-consuming setup process.
func (s *MolecularSimulation) performExpensiveSetup() error {
	if s.expensiveSetupDone {
		fmt.Printf("⏭️ Expensive setup already completed for '%s'. Skipping.\n", s.MoleculeName)
		return nil
	}
	fmt.Printf("⏳ Performing expensive precomputation for '%s' (takes ~2 seconds)...\n", s.MoleculeName)

	// Simulate delay
	time.Sleep(2 * time.Second)

	// Simulate generating large data
	const dataSize = 1_000_000
	s.precomputedStates = make([]float64, dataSize)
	randGen := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed random generator
	for i := 0; i < dataSize; i++ {
		s.precomputedStates[i] = randGen.Float64() * 100
	}

	s.expensiveSetupDone = true
	fmt.Printf("✅ Expensive setup complete for '%s'. %d states computed.\n", s.MoleculeName, len(s.precomputedStates))
	return nil
}

// Clone creates a copy of the MolecularSimulation object.
// It performs a deep copy of parameters but shares the expensive precomputed data.
func (s *MolecularSimulation) Clone() *MolecularSimulation {
	fmt.Printf("\n🔄 Cloning simulation for '%s'...\n", s.MoleculeName)

	// Create a new map for parameters and copy values (deep copy for this level)
	clonedParams := make(ExperimentParameters, len(s.Parameters))
	for key, value := range s.Parameters {
		// Note: This is a shallow copy for nested structures within parameters.
		// For true deep copy of arbitrary parameters, a library or recursive copy needed.
		clonedParams[key] = value
	}

	// Create the clone, *sharing* the precomputedStates slice
	clone := &MolecularSimulation{
		MoleculeName:       s.MoleculeName, // Name is usually immutable string, copy is fine
		Parameters:         clonedParams,   // Use the newly copied parameter map
		precomputedStates:  s.precomputedStates, // Share the slice reference (points to same underlying array)
		expensiveSetupDone: true,               // Mark setup as done for the clone
	}

	fmt.Printf("    Cloned simulation created. Setup Skipped (expensiveSetupDone=%t).\n", clone.expensiveSetupDone)
	return clone
}

// SetParameter modifies a specific parameter for this simulation instance.
func (s *MolecularSimulation) SetParameter(key string, value interface{}) {
	fmt.Printf("    Setting parameter '%s' = %v for '%s' simulation\n", key, value, s.MoleculeName)
	s.Parameters[key] = value
}

// Run executes the simulation using its current parameters and precomputed data.
func (s *MolecularSimulation) Run() {
	fmt.Printf("\n🔬 Running simulation for '%s' with parameters: %v\n", s.MoleculeName, s.Parameters)
	if len(s.precomputedStates) == 0 {
		fmt.Println("   ❌ Error: Precomputed states not available!")
		return
	}

	// Safely get parameters with defaults
	temp, ok := s.Parameters["temperature"].(float64)
	if !ok {
		temp = 298.15 // Default Kelvin
	}
	pressure, ok := s.Parameters["pressure"].(float64)
	if !ok {
		pressure = 1.0 // Default atm
	}
	duration, ok := s.Parameters["duration"].(float64)
	if !ok {
		duration = 100.0 // Default picoseconds
	}

	// Example calculation (placeholder)
	var sumOfStates float64
	limit := 1000
	if len(s.precomputedStates) < limit {
		limit = len(s.precomputedStates)
	}
	for i := 0; i < limit; i++ {
		sumOfStates += s.precomputedStates[i]
	}

	resultMetric := 0.0
	if pressure != 0 {
		resultMetric = sumOfStates * (temp / 273.15) / pressure * (duration / 10.0)
	}

	fmt.Printf("   Simulation complete. Result metric: %.2f\n", resultMetric)
	fmt.Printf("   (Used %d precomputed states)\n", len(s.precomputedStates))
}

// GetPrecomputedStatesLength is a helper for testing
func (s *MolecularSimulation) GetPrecomputedStatesLength() int {
	if s.precomputedStates == nil {
		return 0
	}
	return len(s.precomputedStates)
}
