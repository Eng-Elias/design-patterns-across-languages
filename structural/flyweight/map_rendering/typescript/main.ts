import { MapRenderer, VehicleType, VehicleStatus } from "./map_rendering";

function main(): void {
  // Create a map renderer
  const mapRenderer = new MapRenderer();

  // Add some vehicles
  mapRenderer.addVehicle("car1", VehicleType.CAR);
  mapRenderer.addVehicle("bus1", VehicleType.BUS);
  mapRenderer.addVehicle("truck1", VehicleType.TRUCK);
  mapRenderer.addVehicle("car2", VehicleType.CAR); // Same type as car1

  // Update vehicle states
  mapRenderer.updateVehicle(
    "car1",
    { latitude: 37.7749, longitude: -122.4194 },
    60,
    VehicleStatus.MOVING,
    90
  );

  mapRenderer.updateVehicle(
    "bus1",
    { latitude: 37.7833, longitude: -122.4167 },
    30,
    VehicleStatus.STOPPED,
    180
  );

  mapRenderer.updateVehicle(
    "truck1",
    { latitude: 37.775, longitude: -122.4183 },
    45,
    VehicleStatus.MOVING,
    270
  );

  mapRenderer.updateVehicle(
    "car2",
    { latitude: 37.774, longitude: -122.42 },
    55,
    VehicleStatus.MOVING,
    0
  );

  // Render the map
  const renderedMap = mapRenderer.renderMap();
  console.log("Rendered Map:");
  renderedMap.forEach((renderedVehicle) => {
    console.log(renderedVehicle);
  });

  // Demonstrate memory efficiency
  console.log(`\nMemory Efficiency:`);
  console.log(`Number of vehicles: ${mapRenderer.getVehicleCount()}`);
  console.log(`Number of unique icons: ${mapRenderer.getIconCount()}`);
}

main();
