from dataclasses import dataclass
from typing import Dict, List
from enum import Enum


class VehicleType(Enum):
    CAR = "car"
    BUS = "bus"
    TRUCK = "truck"
    MOTORCYCLE = "motorcycle"


class VehicleStatus(Enum):
    IDLE = "idle"
    MOVING = "moving"
    STOPPED = "stopped"
    OFFLINE = "offline"


@dataclass
class VehicleIcon:
    type: VehicleType
    image_data: str  # In a real app, this would be the actual image data
    width: int
    height: int


@dataclass
class VehicleState:
    position: tuple[float, float]  # (latitude, longitude)
    speed: float
    status: VehicleStatus
    heading: float  # in degrees


class VehicleIconFactory:
    _icons: Dict[VehicleType, VehicleIcon] = {}

    @classmethod
    def get_icon(cls, vehicle_type: VehicleType) -> VehicleIcon:
        if vehicle_type not in cls._icons:
            # In a real app, this would load the actual image data
            cls._icons[vehicle_type] = VehicleIcon(
                type=vehicle_type,
                image_data=f"icon_data_for_{vehicle_type.value}",
                width=32,
                height=32
            )
        return cls._icons[vehicle_type]


class Vehicle:
    def __init__(self, id: str, vehicle_type: VehicleType):
        self.id = id
        self._icon = VehicleIconFactory.get_icon(vehicle_type)
        self._state = VehicleState(
            position=(0.0, 0.0),
            speed=0.0,
            status=VehicleStatus.IDLE,
            heading=0.0
        )

    def update_state(self, position: tuple[float, float], speed: float,
                    status: VehicleStatus, heading: float) -> None:
        self._state = VehicleState(position, speed, status, heading)

    def render(self) -> str:
        return (
            f"Rendering {self.id} at {self._state.position} with "
            f"icon {self._icon.image_data}"
        )


class MapRenderer:
    def __init__(self):
        self._vehicles: Dict[str, Vehicle] = {}

    def add_vehicle(self, vehicle_id: str, vehicle_type: VehicleType) -> None:
        if vehicle_id not in self._vehicles:
            self._vehicles[vehicle_id] = Vehicle(vehicle_id, vehicle_type)

    def update_vehicle(self, vehicle_id: str, position: tuple[float, float],
                      speed: float, status: VehicleStatus, heading: float) -> None:
        if vehicle_id in self._vehicles:
            self._vehicles[vehicle_id].update_state(
                position, speed, status, heading
            )

    def render_map(self) -> List[str]:
        return [vehicle.render() for vehicle in self._vehicles.values()]

    def get_vehicle_count(self) -> int:
        return len(self._vehicles)

    def get_icon_count(self) -> int:
        return len(VehicleIconFactory._icons) 