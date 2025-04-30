package main

import (
	"fmt"

	"flyweight_pattern_map_rendering_go/map_rendering"
)

func main() {
	// Create a map renderer
	renderer := map_rendering.NewMapRenderer()

	// Add some vehicles
	renderer.AddVehicle("car1", map_rendering.Car)
	renderer.AddVehicle("bus1", map_rendering.Bus)
	renderer.AddVehicle("truck1", map_rendering.Truck)
	renderer.AddVehicle("car2", map_rendering.Car) // Same type as car1

	// Update vehicle states
	renderer.UpdateVehicle(
		"car1",
		map_rendering.Position{37.7749, -122.4194},
		60,
		map_rendering.Moving,
		90,
	)

	renderer.UpdateVehicle(
		"bus1",
		map_rendering.Position{37.7833, -122.4167},
		30,
		map_rendering.Stopped,
		180,
	)

	renderer.UpdateVehicle(
		"truck1",
		map_rendering.Position{37.7750, -122.4183},
		45,
		map_rendering.Moving,
		270,
	)

	renderer.UpdateVehicle(
		"car2",
		map_rendering.Position{37.7740, -122.4200},
		55,
		map_rendering.Moving,
		0,
	)

	// Render the map
	fmt.Println("Rendered Map:")
	renderedMap := renderer.RenderMap()
	for _, renderedVehicle := range renderedMap {
		fmt.Println(renderedVehicle)
	}

	// Demonstrate memory efficiency
	fmt.Printf("\nMemory Efficiency:\n")
	fmt.Printf("Number of vehicles: %d\n", renderer.GetVehicleCount())
	fmt.Printf("Number of unique icons: %d\n", renderer.GetIconCount())
} 