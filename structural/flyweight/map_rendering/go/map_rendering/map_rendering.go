package map_rendering

import "fmt"

// VehicleType represents the type of vehicle
type VehicleType string

const (
	Car        VehicleType = "car"
	Bus        VehicleType = "bus"
	Truck      VehicleType = "truck"
	Motorcycle VehicleType = "motorcycle"
)

// VehicleStatus represents the current status of a vehicle
type VehicleStatus string

const (
	Idle    VehicleStatus = "idle"
	Moving  VehicleStatus = "moving"
	Stopped VehicleStatus = "stopped"
	Offline VehicleStatus = "offline"
)

// Position represents a geographical position
type Position struct {
	Latitude  float64
	Longitude float64
}

// VehicleIcon represents the icon data for a vehicle type
type VehicleIcon struct {
	Type      VehicleType
	ImageData string // In a real app, this would be the actual image data
	Width     int
	Height    int
}

// VehicleState represents the current state of a vehicle
type VehicleState struct {
	Position Position
	Speed    float64
	Status   VehicleStatus
	Heading  float64 // in degrees
}

// VehicleIconFactory manages the creation and sharing of vehicle icons
type VehicleIconFactory struct {
	icons map[VehicleType]*VehicleIcon
}

// NewVehicleIconFactory creates a new VehicleIconFactory
func NewVehicleIconFactory() *VehicleIconFactory {
	return &VehicleIconFactory{
		icons: make(map[VehicleType]*VehicleIcon),
	}
}

// GetIcon returns the icon for a vehicle type, creating it if necessary
func (f *VehicleIconFactory) GetIcon(vehicleType VehicleType) *VehicleIcon {
	if icon, exists := f.icons[vehicleType]; exists {
		return icon
	}

	// In a real app, this would load the actual image data
	icon := &VehicleIcon{
		Type:      vehicleType,
		ImageData: "icon_data_for_" + string(vehicleType),
		Width:     32,
		Height:    32,
	}
	f.icons[vehicleType] = icon
	return icon
}

// GetIconCount returns the number of unique icons
func (f *VehicleIconFactory) GetIconCount() int {
	return len(f.icons)
}

// ClearCache clears the icon cache
func (f *VehicleIconFactory) ClearCache() {
	f.icons = make(map[VehicleType]*VehicleIcon)
}

// Vehicle represents a vehicle on the map
type Vehicle struct {
	ID    string
	icon  *VehicleIcon
	state VehicleState
}

// NewVehicle creates a new Vehicle
func NewVehicle(id string, vehicleType VehicleType, factory *VehicleIconFactory) *Vehicle {
	return &Vehicle{
		ID:   id,
		icon: factory.GetIcon(vehicleType),
		state: VehicleState{
			Position: Position{0, 0},
			Speed:    0,
			Status:   Idle,
			Heading:  0,
		},
	}
}

// UpdateState updates the vehicle's state
func (v *Vehicle) UpdateState(position Position, speed float64, status VehicleStatus, heading float64) {
	v.state = VehicleState{
		Position: position,
		Speed:    speed,
		Status:   status,
		Heading:  heading,
	}
}

// Render returns a string representation of the rendered vehicle
func (v *Vehicle) Render() string {
	return fmt.Sprintf(
		"Rendering %s at (%.4f, %.4f) with icon %s",
		v.ID,
		v.state.Position.Latitude,
		v.state.Position.Longitude,
		v.icon.ImageData,
	)
}

// MapRenderer manages multiple vehicles on a map
type MapRenderer struct {
	vehicles map[string]*Vehicle
	factory  *VehicleIconFactory
}

// NewMapRenderer creates a new MapRenderer
func NewMapRenderer() *MapRenderer {
	return &MapRenderer{
		vehicles: make(map[string]*Vehicle),
		factory:  NewVehicleIconFactory(),
	}
}

// AddVehicle adds a new vehicle to the map
func (r *MapRenderer) AddVehicle(id string, vehicleType VehicleType) {
	if _, exists := r.vehicles[id]; !exists {
		r.vehicles[id] = NewVehicle(id, vehicleType, r.factory)
	}
}

// UpdateVehicle updates the state of a vehicle
func (r *MapRenderer) UpdateVehicle(id string, position Position, speed float64, status VehicleStatus, heading float64) {
	if vehicle, exists := r.vehicles[id]; exists {
		vehicle.UpdateState(position, speed, status, heading)
	}
}

// RenderMap renders all vehicles on the map
func (r *MapRenderer) RenderMap() []string {
	rendered := make([]string, 0, len(r.vehicles))
	for _, vehicle := range r.vehicles {
		rendered = append(rendered, vehicle.Render())
	}
	return rendered
}

// GetVehicleCount returns the number of vehicles on the map
func (r *MapRenderer) GetVehicleCount() int {
	return len(r.vehicles)
}

// GetIconCount returns the number of unique icons
func (r *MapRenderer) GetIconCount() int {
	return r.factory.GetIconCount()
} 