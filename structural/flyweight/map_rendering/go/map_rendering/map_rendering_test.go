package map_rendering

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapRendering(t *testing.T) {
	// Create a new map renderer
	renderer := NewMapRenderer()

	// Add some vehicles
	renderer.AddVehicle("car1", Car)
	renderer.AddVehicle("bus1", Bus)
	renderer.AddVehicle("truck1", Truck)
	renderer.AddVehicle("car2", Car) // Same type as car1

	// Update vehicle states
	renderer.UpdateVehicle(
		"car1",
		Position{37.7749, -122.4194},
		60,
		Moving,
		90,
	)

	renderer.UpdateVehicle(
		"bus1",
		Position{37.7833, -122.4167},
		30,
		Stopped,
		180,
	)

	renderer.UpdateVehicle(
		"truck1",
		Position{37.7750, -122.4183},
		45,
		Moving,
		270,
	)

	renderer.UpdateVehicle(
		"car2",
		Position{37.7740, -122.4200},
		55,
		Moving,
		0,
	)

	// Render the map
	renderedMap := renderer.RenderMap()
	assert.Len(t, renderedMap, 4, "Should render all vehicles")

	// Check memory efficiency
	assert.Equal(t, 4, renderer.GetVehicleCount(), "Should have 4 vehicles")
	assert.Equal(t, 3, renderer.GetIconCount(), "Should have 3 unique icons (car, bus, truck)")

	// Test icon sharing
	car1 := renderer.vehicles["car1"]
	car2 := renderer.vehicles["car2"]
	assert.Equal(t, car1.icon, car2.icon, "Cars should share the same icon instance")
}

func TestVehicleIconFactory(t *testing.T) {
	factory := NewVehicleIconFactory()

	// Get icons for different vehicle types
	carIcon1 := factory.GetIcon(Car)
	carIcon2 := factory.GetIcon(Car)
	busIcon := factory.GetIcon(Bus)

	// Test icon sharing
	assert.Equal(t, carIcon1, carIcon2, "Same vehicle type should return the same icon instance")
	assert.NotEqual(t, carIcon1, busIcon, "Different vehicle types should return different icon instances")

	// Test icon count
	assert.Equal(t, 2, factory.GetIconCount(), "Should have 2 unique icons")

	// Test cache clearing
	factory.ClearCache()
	assert.Equal(t, 0, factory.GetIconCount(), "Cache should be empty after clearing")
} 