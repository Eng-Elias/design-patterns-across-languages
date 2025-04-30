package imaging

import "fmt"

// --- Proxy --- //

// ProxyImage acts as a surrogate for RealImage, controlling access and enabling lazy loading.
type ProxyImage struct {
	filename  string
	realImage *RealImage // Pointer to the RealImage, initially nil
}

// NewProxyImage creates a new proxy for an image.
// Note: This does NOT load the RealImage yet.
func NewProxyImage(filename string) *ProxyImage {
	fmt.Printf("ProxyImage created for: '%s' (Real image not loaded yet)\n", filename)
	return &ProxyImage{
		filename:  filename,
		realImage: nil, // Explicitly nil initially
	}
}

// Display handles the display request.
// It loads the RealImage only if it hasn't been loaded yet (Lazy Initialization).
func (p *ProxyImage) Display() {
	if p.realImage == nil {
		fmt.Printf("Proxy for '%s': Real image needs loading.\n", p.filename)
		// Lazy initialization: Create the RealImage object only when needed
		p.realImage = NewRealImage(p.filename)
	} else {
		fmt.Printf("Proxy for '%s': Real image already loaded.\n", p.filename)
	}

	// Delegate the display call to the RealImage object
	p.realImage.Display()
}

// GetFilename returns the filename without loading the real image.
func (p *ProxyImage) GetFilename() string {
	return p.filename
}

// IsLoaded checks if the real image has been loaded (optional helper).
func (p *ProxyImage) IsLoaded() bool {
	return p.realImage != nil
}
