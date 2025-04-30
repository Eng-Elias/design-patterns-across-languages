package main

import (
	"fmt"
	"proxy_pattern_image_loading_go/imaging"
	"time"
)

func main() {
	fmt.Println("--- Creating image proxies --- ")
	// Creating proxy objects doesn't load the actual images yet
	// Use the Image interface type for the variables
	image1 := imaging.NewProxyImage("family_vacation_2024.jpeg")
	image2 := imaging.NewProxyImage("company_logo_final.svg")
	image3 := imaging.NewProxyImage("research_paper_scan.pdf") // Example filename

	fmt.Println("\n--- Accessing filenames (should not load) --- ")
	fmt.Printf("Image 1 filename: %s\n", image1.GetFilename())
	fmt.Printf("Image 2 filename: %s\n", image2.GetFilename())

	// Check loaded status (optional method)
	fmt.Printf("Image 1 loaded status: %t\n", image1.IsLoaded()) // false

	fmt.Println("\n--- Requesting display for image 1 (will trigger loading) --- ")
	startTime := time.Now()
	image1.Display()
	duration := time.Since(startTime)
	fmt.Printf("(Display took %v)\n", duration)

	fmt.Printf("Image 1 loaded status: %t\n", image1.IsLoaded()) // true

	fmt.Println("\n--- Requesting display for image 1 AGAIN (should be fast) --- ")
	startTime = time.Now()
	image1.Display() // The real image is already loaded, delegation is fast
	duration = time.Since(startTime)
	fmt.Printf("(Display took %v)\n", duration)

	fmt.Println("\n--- Requesting display for image 2 (will trigger loading) --- ")
	startTime = time.Now()
	image2.Display()
	duration = time.Since(startTime)
	fmt.Printf("(Display took %v)\n", duration)

	fmt.Println("\n--- Requesting display for image 3 (will trigger loading) --- ")
	startTime = time.Now()
	image3.Display()
	duration = time.Since(startTime)
	fmt.Printf("(Display took %v)\n", duration)

	fmt.Println("\nDemo finished.")
}
