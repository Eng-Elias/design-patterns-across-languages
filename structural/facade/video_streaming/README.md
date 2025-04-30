# Video Streaming Service - Facade Pattern

## Problem

Building a video streaming service involves multiple complex subsystems that need to work together:
- Video uploading
- Video encoding
- Storage management
- Content delivery
- User authentication
- Notification handling

Managing these subsystems directly can lead to complex, tightly coupled code that's hard to maintain and extend.

## Solution / Implementation

The **Facade** pattern provides a simplified interface to a complex subsystem of classes. In this case, we'll create a `VideoStreamingService` facade that hides the complexity of the underlying subsystems and provides a simple interface for clients to interact with.

### Subsystems

1. **VideoUploader**: Handles video file uploads
2. **VideoEncoder**: Converts videos to different formats
3. **StorageService**: Manages video storage in the cloud
4. **CDNService**: Distributes videos through a content delivery network
5. **AuthenticationService**: Handles user authentication
6. **NotificationService**: Sends status updates to users

### Facade

The `VideoStreamingService` facade provides a simple interface for:
- Uploading and processing videos
- Streaming videos to users
- Managing user access

This allows clients to interact with the video streaming service without needing to understand the complex interactions between subsystems.

## Setup

Instructions assume you are in the `structural/facade/video_streaming` directory.

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

Instructions assume you are in the `structural/facade/video_streaming` directory.

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

Instructions assume you are in the `structural/facade/video_streaming` directory.

### Python
```bash
cd python
python -m unittest test_video_streaming.py
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