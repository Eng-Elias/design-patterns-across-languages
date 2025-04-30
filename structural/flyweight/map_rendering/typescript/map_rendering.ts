// Enums
export enum VehicleType {
  CAR = "car",
  BUS = "bus",
  TRUCK = "truck",
  MOTORCYCLE = "motorcycle",
}

export enum VehicleStatus {
  IDLE = "idle",
  MOVING = "moving",
  STOPPED = "stopped",
  OFFLINE = "offline",
}

// Types
export interface Position {
  latitude: number;
  longitude: number;
}

export interface VehicleIcon {
  type: VehicleType;
  imageData: string; // In a real app, this would be the actual image data
  width: number;
  height: number;
}

export interface VehicleState {
  position: Position;
  speed: number;
  status: VehicleStatus;
  heading: number; // in degrees
}

// Flyweight Factory
export class VehicleIconFactory {
  private static icons: Map<VehicleType, VehicleIcon> = new Map();

  static getIcon(vehicleType: VehicleType): VehicleIcon {
    if (!this.icons.has(vehicleType)) {
      // In a real app, this would load the actual image data
      this.icons.set(vehicleType, {
        type: vehicleType,
        imageData: `icon_data_for_${vehicleType}`,
        width: 32,
        height: 32,
      });
    }
    return this.icons.get(vehicleType)!;
  }

  static getIconCount(): number {
    return this.icons.size;
  }

  static clearCache(): void {
    this.icons.clear();
  }
}

// Vehicle class
export class Vehicle {
  private readonly id: string;
  private readonly icon: VehicleIcon;
  private state: VehicleState;

  constructor(id: string, vehicleType: VehicleType) {
    this.id = id;
    this.icon = VehicleIconFactory.getIcon(vehicleType);
    this.state = {
      position: { latitude: 0, longitude: 0 },
      speed: 0,
      status: VehicleStatus.IDLE,
      heading: 0,
    };
  }

  updateState(
    position: Position,
    speed: number,
    status: VehicleStatus,
    heading: number
  ): void {
    this.state = { position, speed, status, heading };
  }

  render(): string {
    return `Rendering ${this.id} at (${this.state.position.latitude}, ${this.state.position.longitude}) with status ${this.state.status} and icon ${this.icon.imageData}`;
  }
}

// Map Renderer
export class MapRenderer {
  private vehicles: Map<string, Vehicle> = new Map();

  addVehicle(vehicleId: string, vehicleType: VehicleType): void {
    if (!this.vehicles.has(vehicleId)) {
      this.vehicles.set(vehicleId, new Vehicle(vehicleId, vehicleType));
    }
  }

  updateVehicle(
    vehicleId: string,
    position: Position,
    speed: number,
    status: VehicleStatus,
    heading: number
  ): void {
    const vehicle = this.vehicles.get(vehicleId);
    if (vehicle) {
      vehicle.updateState(position, speed, status, heading);
    }
  }

  renderMap(): string[] {
    return Array.from(this.vehicles.values()).map((vehicle) =>
      vehicle.render()
    );
  }

  getVehicleCount(): number {
    return this.vehicles.size;
  }

  getIconCount(): number {
    return VehicleIconFactory.getIconCount();
  }
}
