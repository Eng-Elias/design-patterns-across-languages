from map_rendering import MapRenderer, VehicleType, VehicleStatus


def main():
    # Create a map renderer
    renderer = MapRenderer()

    # Add some vehicles
    renderer.add_vehicle("car1", VehicleType.CAR)
    renderer.add_vehicle("bus1", VehicleType.BUS)
    renderer.add_vehicle("truck1", VehicleType.TRUCK)
    renderer.add_vehicle("car2", VehicleType.CAR)  # Same type as car1

    # Update vehicle states
    renderer.update_vehicle(
        "car1",
        position=(37.7749, -122.4194),  # San Francisco
        speed=45.0,
        status=VehicleStatus.MOVING,
        heading=90.0
    )

    renderer.update_vehicle(
        "bus1",
        position=(37.7749, -122.4194),
        speed=30.0,
        status=VehicleStatus.STOPPED,
        heading=180.0
    )

    renderer.update_vehicle(
        "truck1",
        position=(37.7749, -122.4194),
        speed=0.0,
        status=VehicleStatus.IDLE,
        heading=0.0
    )

    renderer.update_vehicle(
        "car2",
        position=(37.7749, -122.4194),
        speed=60.0,
        status=VehicleStatus.MOVING,
        heading=270.0
    )

    # Render the map
    print("--- Rendering Map ---")
    for render_output in renderer.render_map():
        print(render_output)

    # Demonstrate memory efficiency
    print("\n--- Memory Efficiency ---")
    print(f"Number of vehicles: {renderer.get_vehicle_count()}")
    print(f"Number of unique icons: {renderer.get_icon_count()}")


if __name__ == "__main__":
    main() 