# Map Rendering System - Flyweight Pattern

## Problem

When rendering a large number of vehicles on a map, each vehicle requires:

- A unique identifier
- Current state (position, speed, status, heading)
- Visual representation (icon)

If we create a separate icon object for each vehicle, we'll quickly consume a large amount of memory, especially when many vehicles share the same type (e.g., multiple cars using the same car icon).

## Solution / Implementation

The **Flyweight** pattern helps reduce memory usage by sharing common parts of objects between multiple objects. In this case, we'll share vehicle icons between vehicles of the same type.

### Key Components

1. **VehicleType**: Enum defining different types of vehicles (CAR, BUS, TRUCK, MOTORCYCLE)
2. **VehicleStatus**: Enum defining vehicle states (IDLE, MOVING, STOPPED, OFFLINE)
3. **VehicleIcon**: The flyweight object that contains shared icon data
4. **VehicleIconFactory**: Manages the creation and sharing of vehicle icons
5. **Vehicle**: Contains unique state and references the shared icon
6. **MapRenderer**: Manages multiple vehicles and their rendering

### Flyweight Implementation

The `VehicleIconFactory` ensures that:

- Only one icon instance is created per vehicle type
- The same icon instance is shared between all vehicles of the same type
- Memory usage is optimized by reusing icon data

## Setup

Instructions assume you are in the `structural/flyweight/map_rendering` directory.

### Python

```bash
# No specific setup required, uses standard libraries.
```

### TypeScript

```bash
# Install Node.js/npm if you haven't already.
cd typescript
npm install
```

### Go

```bash
# Ensure Go is installed.
# The go.mod file defines the module.
```

## How to Run

Instructions assume you are in the `structural/flyweight/map_rendering` directory.

### Python

```bash
cd python
python main.py
```

### TypeScript

```bash
cd typescript
npm start
```

### Go

```bash
cd go
go run main.go
```

## How to Test

Instructions assume you are in the `structural/flyweight/map_rendering` directory.

### Python

```bash
cd python
python -m unittest test_map_rendering.py
```

### TypeScript

```bash
cd typescript
npm test
```

### Go

```bash
cd go
go test -v ./...
```
