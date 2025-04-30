import { describe, beforeEach, test, expect } from '@jest/globals';
import {
  MapRenderer,
  VehicleType,
  VehicleStatus,
  VehicleIconFactory,
} from "./map_rendering";

describe("MapRendering", () => {
  let renderer: MapRenderer;

  beforeEach(() => {
    renderer = new MapRenderer();
    VehicleIconFactory.clearCache();
  });

  test("should add a new vehicle to the map", () => {
    renderer.addVehicle("car1", VehicleType.CAR);
    expect(renderer.getVehicleCount()).toBe(1);
  });

  test("should update vehicle state", () => {
    renderer.addVehicle("car1", VehicleType.CAR);
    renderer.updateVehicle(
      "car1",
      { latitude: 37.7749, longitude: -122.4194 },
      45.0,
      VehicleStatus.MOVING,
      90.0
    );
    const renderOutput = renderer.renderMap()[0];
    expect(renderOutput).toContain("car1");
    expect(renderOutput).toContain("37.7749");
    expect(renderOutput).toContain("-122.4194");
  });

  test("should share icons between vehicles of the same type", () => {
    renderer.addVehicle("car1", VehicleType.CAR);
    renderer.addVehicle("car2", VehicleType.CAR);
    renderer.addVehicle("bus1", VehicleType.BUS);

    expect(renderer.getVehicleCount()).toBe(3);
    expect(renderer.getIconCount()).toBe(2);
  });

  test("should render all vehicles on the map", () => {
    renderer.addVehicle("car1", VehicleType.CAR);
    renderer.addVehicle("bus1", VehicleType.BUS);

    renderer.updateVehicle(
      "car1",
      { latitude: 37.7749, longitude: -122.4194 },
      45.0,
      VehicleStatus.MOVING,
      90.0
    );
    renderer.updateVehicle(
      "bus1",
      { latitude: 37.7749, longitude: -122.4194 },
      30.0,
      VehicleStatus.STOPPED,
      180.0
    );

    const renderOutputs = renderer.renderMap();
    expect(renderOutputs).toHaveLength(2);
    expect(renderOutputs.some(output => output.includes("car1"))).toBe(true);
    expect(renderOutputs.some(output => output.includes("bus1"))).toBe(true);
  });
});

describe("VehicleIconFactory", () => {
  beforeEach(() => {
    VehicleIconFactory.clearCache();
  });

  test("should create and return icon for new vehicle type", () => {
    const icon = VehicleIconFactory.getIcon(VehicleType.CAR);
    expect(icon.type).toBe(VehicleType.CAR);
    expect(icon.imageData).toBe("icon_data_for_car");
    expect(icon.width).toBe(32);
    expect(icon.height).toBe(32);
  });

  test("should return the same icon instance for the same vehicle type", () => {
    const icon1 = VehicleIconFactory.getIcon(VehicleType.CAR);
    const icon2 = VehicleIconFactory.getIcon(VehicleType.CAR);
    expect(icon1).toBe(icon2);
  });

  test("should return different icon instances for different vehicle types", () => {
    const carIcon = VehicleIconFactory.getIcon(VehicleType.CAR);
    const busIcon = VehicleIconFactory.getIcon(VehicleType.BUS);
    expect(carIcon).not.toBe(busIcon);
    expect(carIcon.type).toBe(VehicleType.CAR);
    expect(busIcon.type).toBe(VehicleType.BUS);
  });

  test("should track the number of unique icons", () => {
    VehicleIconFactory.getIcon(VehicleType.CAR);
    VehicleIconFactory.getIcon(VehicleType.CAR); // Same type
    VehicleIconFactory.getIcon(VehicleType.BUS);
    expect(VehicleIconFactory.getIconCount()).toBe(2);
  });

  test("should clear the icon cache", () => {
    VehicleIconFactory.getIcon(VehicleType.CAR);
    VehicleIconFactory.getIcon(VehicleType.BUS);
    VehicleIconFactory.clearCache();
    expect(VehicleIconFactory.getIconCount()).toBe(0);
  });
});
