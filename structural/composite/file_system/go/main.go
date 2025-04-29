package main

import (
	fs "composite_pattern_file_system_go/file_system" // Alias for clarity
	"fmt"
)

func main() {
	fmt.Println("Demonstrating the Composite pattern using a file system example in Go.")

	// Create some files (Leaf objects) - use NewFile constructor
	file1 := fs.NewFile("document.txt", 1024) // 1 KB
	file2 := fs.NewFile("image.jpg", 5120)    // 5 KB
	file3 := fs.NewFile("archive.zip", 10240) // 10 KB
	file4 := fs.NewFile("report.pdf", 2048)   // 2 KB

	// Create some directories (Composite objects) - use NewDirectory constructor
	root := fs.NewDirectory("root")
	documentsDir := fs.NewDirectory("Documents")
	picturesDir := fs.NewDirectory("Pictures")
	privateDir := fs.NewDirectory("Private")

	// Build the file system tree structure using the Add method
	// Note: Add takes the FileSystemComponent interface type
	root.Add(documentsDir)
	root.Add(picturesDir)

	documentsDir.Add(file1)
	documentsDir.Add(file4)
	documentsDir.Add(privateDir) // Add a subdirectory

	picturesDir.Add(file2)

	privateDir.Add(file3) // Add a file to the subdirectory

	// --- Demonstrate Uniform Treatment --- //

	fmt.Println("\n--- Displaying the entire file system structure ---")
	root.Display("") // Start with no indent

	fmt.Println("\n--- Calculating sizes ---")

	// Calculate size of the entire root directory (Composite)
	fmt.Printf("Total size of '%s': %d bytes\n", root.GetName(), root.GetSize())

	// Calculate size of a subdirectory (Composite)
	fmt.Printf("Total size of '%s': %d bytes\n", documentsDir.GetName(), documentsDir.GetSize())

	// Calculate size of an individual file (Leaf)
	// We can call GetSize() directly on the File object as well
	fmt.Printf("Size of '%s': %d bytes\n", file1.GetName(), file1.GetSize())

	// Display a specific subdirectory (Composite)
	fmt.Println("\n--- Displaying the 'Documents' directory ---")
	documentsDir.Display("") // Start with no indent
}
