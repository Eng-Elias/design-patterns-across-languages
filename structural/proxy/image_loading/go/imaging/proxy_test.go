package imaging

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// --- Test Helpers --- //

// Store the original sleep function from the imaging package
var originalSleep = sleepFunc

func patchSleep() {
    // Patch the package-level variable in imaging
	sleepFunc = func(d time.Duration) { /* Do nothing, skip sleep */ }
}

func unpatchSleep() {
    // Restore the original sleep function
	sleepFunc = originalSleep
}

// --- ProxyImage Tests --- //

func TestNewProxyImage(t *testing.T) {
	filename := "proxy_init_test.jpg"
	proxy := NewProxyImage(filename)

	assert.NotNil(t, proxy, "Proxy should not be nil")
	assert.Equal(t, filename, proxy.filename, "Proxy should store the correct filename")
	assert.Nil(t, proxy.realImage, "RealImage should be nil initially")
	assert.False(t, proxy.IsLoaded(), "Proxy should report not loaded initially")
}

func TestProxyGetFilename(t *testing.T) {
	filename := "proxy_getfn_test.png"
	proxy := NewProxyImage(filename)

	retrievedFilename := proxy.GetFilename()

	assert.Equal(t, filename, retrievedFilename, "GetFilename should return the correct name")
	// Verify that getting the filename does not load the real image
	assert.Nil(t, proxy.realImage, "RealImage should still be nil after GetFilename")
	assert.False(t, proxy.IsLoaded(), "Proxy should still report not loaded after GetFilename")
}

func TestProxyDisplay_LazyLoadingAndReuse(t *testing.T) {
	patchSleep()       // Speed up loading simulation
	defer unpatchSleep() // Ensure original sleep is restored

	filename := "proxy_display_test.gif"
	proxy := NewProxyImage(filename)

	// 1. Check initial state
	assert.False(t, proxy.IsLoaded(), "Proxy should not be loaded before first display")
	assert.Nil(t, proxy.realImage, "RealImage field should be nil initially")

	// 2. First display call (triggers loading)
	t.Log("Calling Display() for the first time...")
	proxy.Display()

	assert.True(t, proxy.IsLoaded(), "Proxy should be loaded after first display")
	assert.NotNil(t, proxy.realImage, "RealImage field should not be nil after first display")
	assert.Equal(t, filename, proxy.realImage.GetFilename(), "Loaded RealImage should have correct filename")

	// Store the pointer to the loaded RealImage
	loadedRealImage := proxy.realImage

	// 3. Second display call (should reuse)
	t.Log("Calling Display() for the second time...")
	proxy.Display()

	assert.True(t, proxy.IsLoaded(), "Proxy should still be loaded after second display")
	assert.NotNil(t, proxy.realImage, "RealImage field should still not be nil")
	// Crucial check: Ensure the *same* RealImage instance is used
	assert.Same(t, loadedRealImage, proxy.realImage, "Second display call should reuse the existing RealImage instance")
}

func TestProxyIsLoaded(t *testing.T) {
	patchSleep()       // Speed up loading simulation
	defer unpatchSleep() // Ensure original sleep is restored

	filename := "proxy_isloaded_test.bmp"
	proxy := NewProxyImage(filename)

	assert.False(t, proxy.IsLoaded(), "IsLoaded should return false initially")

	// Call display to trigger loading
	proxy.Display()

	assert.True(t, proxy.IsLoaded(), "IsLoaded should return true after display is called")
}

// --- RealImage Tests (Basic Functionality) --- //

func TestNewRealImage(t *testing.T) {
	patchSleep()       // Speed up loading simulation
	defer unpatchSleep() // Ensure original sleep is restored

	filename := "real_image_test.tiff"
	// We expect loading logs during creation
	t.Log("Creating RealImage (expect loading logs)...")
	realImage := NewRealImage(filename)

	assert.NotNil(t, realImage, "RealImage should not be nil")
	assert.Equal(t, filename, realImage.filename, "RealImage should store the correct filename")
}

func TestRealImageDisplay(t *testing.T) {
	patchSleep()       // Speed up loading simulation
	defer unpatchSleep() // Ensure original sleep is restored

	filename := "real_display_test.svg"
	realImage := NewRealImage(filename) // Loading happens here

	// Simply call display - we can't easily assert the fmt.Printf output without capturing stdout
	// The main goal is to ensure it runs without error.
	t.Log("Calling RealImage.Display() (expect display log)...")
	assert.NotPanics(t, func() {
		realImage.Display()
	}, "RealImage.Display() should not panic")
}

func TestRealImageGetFilename(t *testing.T) {
	patchSleep()       // Speed up loading simulation
	defer unpatchSleep() // Ensure original sleep is restored

	filename := "real_getfn_test.ico"
	realImage := NewRealImage(filename)

	assert.Equal(t, filename, realImage.GetFilename(), "RealImage.GetFilename should return the correct filename")
}
