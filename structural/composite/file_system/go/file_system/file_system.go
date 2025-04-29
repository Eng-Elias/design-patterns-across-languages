package file_system

import (
	"fmt"
)

// --- Component Interface ---
type FileSystemComponent interface {
	// The common interface for both leaves (files) and composites (directories).
	GetName() string
	GetSize() int // Returns the size of the component in bytes.
	Display(indent string) // Displays the component's structure.
}

// --- Leaf Struct ---
type File struct {
	// Represents a leaf object (a file) in the composition.
	name string
	size int
}

// NewFile creates a new File instance.
func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Display(indent string) {
	fmt.Printf("%s- %s (%d bytes)\n", indent, f.GetName(), f.GetSize())
}

// --- Composite Struct ---
type Directory struct {
	// Represents a composite object (a directory) that can contain other components.
	name     string
	children []FileSystemComponent // Stores child components (Files or other Directories)
}

// NewDirectory creates a new Directory instance.
func NewDirectory(name string) *Directory {
	return &Directory{
		name:     name,
		children: make([]FileSystemComponent, 0), // Initialize the slice
	}
}

func (d *Directory) GetName() string {
	return d.name
}

// Add adds a child component (file or subdirectory).
func (d *Directory) Add(component FileSystemComponent) {
	d.children = append(d.children, component)
}

// Remove removes a child component. Returns false if not found.
func (d *Directory) Remove(component FileSystemComponent) bool {
	for i, child := range d.children {
		// Pointer comparison works if they are the exact same instance,
		// otherwise need a way to uniquely identify (e.g., compare names if unique within dir)
		if child == component { // Simple comparison, might need adjustment based on usage
			// Remove element by slicing
			d.children = append(d.children[:i], d.children[i+1:]...)
			return true
		}
	}
	return false // Not found
}

// GetChild gets a specific child component. Returns nil if index is out of bounds.
func (d *Directory) GetChild(index int) FileSystemComponent {
	if index < 0 || index >= len(d.children) {
		return nil // Or return an error
	}
	return d.children[index]
}

// GetSize calculates the total size by summing the sizes of all children.
func (d *Directory) GetSize() int {
	totalSize := 0
	for _, child := range d.children {
		totalSize += child.GetSize()
	}
	return totalSize
}

// Display displays the directory and recursively displays its children.
func (d *Directory) Display(indent string) {
	fmt.Printf("%s+ %s (%d bytes total)\n", indent, d.GetName(), d.GetSize())
	newIndent := indent + "  "
	for _, child := range d.children {
		child.Display(newIndent)
	}
}
