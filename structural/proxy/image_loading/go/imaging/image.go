package imaging

import (
	"fmt"
	"time"
)

// sleepFunc allows patching time.Sleep for testing
var sleepFunc = time.Sleep

// --- Subject Interface --- //

// Image defines the common interface for RealImage and ProxyImage.
type Image interface {
	Display()
	GetFilename() string
}

// --- Real Subject --- //

// RealImage represents the actual resource-intensive object.
type RealImage struct {
	filename string
	// In a real scenario, this might hold image data, dimensions, etc.
}

// NewRealImage creates a new RealImage instance and simulates loading it.
// This function acts as the constructor and includes the expensive operation.
func NewRealImage(filename string) *RealImage {
	img := &RealImage{
		filename: filename,
	}
	img.loadFromDisk()
	return img
}

// loadFromDisk simulates the expensive operation of loading the image.
func (r *RealImage) loadFromDisk() {
	fmt.Printf("Loading image: '%s' from disk... (Simulating delay)\n", r.filename)
	// Simulate time-consuming operation using the patchable function
	sleepFunc(1500 * time.Millisecond)
	fmt.Printf("Finished loading image: '%s'\n", r.filename)
}

// Display shows the image (assumes it's already loaded).
func (r *RealImage) Display() {
	fmt.Printf("Displaying image: '%s'\n", r.filename)
}

// GetFilename returns the associated filename.
func (r *RealImage) GetFilename() string {
	return r.filename
}
