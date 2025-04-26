package molecular_simulation

import (
	"reflect"
	"testing"
	"time"
)

// Helper function to create a prototype for testing, reducing boilerplate
// It uses a short sleep duration for faster tests.
func createTestPrototype(t *testing.T) *MolecularSimulation {
	baseParams := ExperimentParameters{"temperature": 300.0, "pressure": 1.0}
	proto, err := NewMolecularSimulation("TestMoleculeGo", baseParams)
	if err != nil {
		t.Fatalf("Failed to create test prototype: %v", err)
	}
	// Allow setup time (might need adjustment if real sleep is used)
	// time.Sleep(50 * time.Millisecond)
	return proto
}

func TestNewMolecularSimulation_PerformsExpensiveSetup(t *testing.T) {
	t.Parallel()
	proto := createTestPrototype(t)

	if !proto.expensiveSetupDone {
		t.Error("Expected expensiveSetupDone to be true after initialization")
	}
	if len(proto.precomputedStates) == 0 {
		t.Error("Expected precomputedStates to be populated after initialization")
	}
}

func TestClone_CreatesNewInstance(t *testing.T) {
	t.Parallel()
	proto := createTestPrototype(t)
	clone := proto.Clone()

	if clone == proto {
		t.Error("Clone() should return a new instance, not the same object")
	}
	if reflect.TypeOf(clone) != reflect.TypeOf(proto) {
		t.Errorf("Expected clone type %T, got %T", proto, clone)
	}
}

func TestClone_DoesNotRepeatExpensiveSetup(t *testing.T) {
	t.Parallel()
	proto := createTestPrototype(t)

	startTime := time.Now()
	clone := proto.Clone()
	cloneDuration := time.Since(startTime)

	// Check the flag on the clone
	if !clone.expensiveSetupDone {
		t.Error("Clone should have expensiveSetupDone set to true immediately")
	}

	// Check that cloning was fast (much less than the setup time)
	if cloneDuration >= 2*time.Second {
		t.Errorf("Cloning took %v, expected much less than setup time (2s)", cloneDuration)
	}
}

func TestClone_IndependentParameters(t *testing.T) {
	t.Parallel()
	proto := createTestPrototype(t)
	clone := proto.Clone()

	originalTemp := proto.Parameters["temperature"]
	newTemp := originalTemp.(float64) + 50.0

	clone.SetParameter("temperature", newTemp)
	clone.Parameters["newParam"] = "cloneValue"

	// Check original is unchanged
	if proto.Parameters["temperature"] != originalTemp {
		t.Errorf("Modifying clone temperature changed original. Got %v, want %v", proto.Parameters["temperature"], originalTemp)
	}
	if _, exists := proto.Parameters["newParam"]; exists {
		t.Error("Adding parameter to clone affected original parameters")
	}

	// Check clone has new values
	if clone.Parameters["temperature"] != newTemp {
		t.Errorf("Clone temperature not updated. Got %v, want %v", clone.Parameters["temperature"], newTemp)
	}
	if clone.Parameters["newParam"] != "cloneValue" {
		t.Errorf("Clone newParam not set correctly. Got %v, want 'cloneValue'", clone.Parameters["newParam"])
	}
}

func TestClone_SharesLargeData(t *testing.T) {
	t.Parallel()
	proto := createTestPrototype(t)
	clone := proto.Clone()

	if len(clone.precomputedStates) != len(proto.precomputedStates) {
		t.Errorf("Clone precomputedStates length (%d) differs from original (%d)", len(clone.precomputedStates), len(proto.precomputedStates))
	}

	// Check if the underlying array pointers are the same (indicates sharing)
	// This is a common optimization but not strictly guaranteed by the spec for all copy methods.
	protoHeader := (*reflect.SliceHeader)(reflect.ValueOf(proto.precomputedStates).UnsafePointer())
	cloneHeader := (*reflect.SliceHeader)(reflect.ValueOf(clone.precomputedStates).UnsafePointer())

	if protoHeader.Data != cloneHeader.Data {
		t.Error("Clone() did not share the underlying array data for precomputedStates")
	}
}

func TestRun_ExecutesOnOriginalAndClone(t *testing.T) {
	t.Parallel()
	proto := createTestPrototype(t)
	clone := proto.Clone()
	clone.SetParameter("duration", 50.0)

	// Run both and ensure no panics occur (basic check)
	// More advanced checks could involve capturing output or checking results.
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Original Run() panicked: %v", r)
			}
		}()
		proto.Run()
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Clone Run() panicked: %v", r)
			}
		}()
		clone.Run()
	}()

	// Check if clone has data (should not be nil after clone)
	if clone.precomputedStates == nil {
		t.Error("Clone precomputedStates is nil after cloning")
	}
}
