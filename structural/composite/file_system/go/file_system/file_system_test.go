package file_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_Properties(t *testing.T) {
	file := NewFile("test.txt", 100)
	assert.Equal(t, "test.txt", file.GetName())
	assert.Equal(t, 100, file.GetSize())
}

func TestDirectory_Properties(t *testing.T) {
	dir := NewDirectory("docs")
	assert.Equal(t, "docs", dir.GetName())
	// Initial size should be 0
	assert.Equal(t, 0, dir.GetSize())
	// Initial children slice should be non-nil and empty
	assert.NotNil(t, dir.children)
	assert.Empty(t, dir.children)
}

func TestDirectory_AddAndGetChild(t *testing.T) {
	dir := NewDirectory("docs")
	file1 := NewFile("file1.txt", 50)
	file2 := NewFile("file2.txt", 150)

	dir.Add(file1)
	dir.Add(file2)

	assert.Equal(t, file1, dir.GetChild(0))
	assert.Equal(t, file2, dir.GetChild(1))
	assert.Nil(t, dir.GetChild(2), "Accessing out of bounds index should return nil")
	assert.Nil(t, dir.GetChild(-1), "Accessing negative index should return nil")
	assert.Len(t, dir.children, 2)
}

func TestDirectory_RemoveChild(t *testing.T) {
	dir := NewDirectory("docs")
	file1 := NewFile("file1.txt", 50)
	file2 := NewFile("file2.txt", 150)
	file3 := NewFile("other.txt", 200) // A file not added initially

	dir.Add(file1)
	dir.Add(file2)

	assert.Len(t, dir.children, 2)
	assert.Equal(t, file1, dir.GetChild(0))
	assert.Equal(t, file2, dir.GetChild(1))

	// Remove the first element
	removed := dir.Remove(file1)
	assert.True(t, removed, "Remove should return true when element is found")
	assert.Len(t, dir.children, 1)
	assert.Equal(t, file2, dir.GetChild(0), "file2 should now be at index 0")
	assert.Nil(t, dir.GetChild(1))

	// Try removing again (should fail)
	removed = dir.Remove(file1)
	assert.False(t, removed, "Remove should return false for already removed element")
	assert.Len(t, dir.children, 1)

	// Try removing an element never added
	removed = dir.Remove(file3)
	assert.False(t, removed, "Remove should return false for element never added")
	assert.Len(t, dir.children, 1)

	// Remove the remaining element
	removed = dir.Remove(file2)
	assert.True(t, removed)
	assert.Empty(t, dir.children)
	assert.Nil(t, dir.GetChild(0))
}

func TestDirectory_CalculateSize_Simple(t *testing.T) {
	dir := NewDirectory("docs")
	file1 := NewFile("file1.txt", 50)
	file2 := NewFile("file2.txt", 150)

	dir.Add(file1)
	dir.Add(file2)

	assert.Equal(t, 50+150, dir.GetSize())
}

func TestDirectory_CalculateSize_Nested(t *testing.T) {
	root := NewDirectory("root")
	docs := NewDirectory("docs")
	pics := NewDirectory("pics")
	private := NewDirectory("private")

	fileR1 := NewFile("root_file.log", 10)
	fileD1 := NewFile("doc1.txt", 100)
	fileD2 := NewFile("doc2.pdf", 200)
	fileP1 := NewFile("pic1.jpg", 500)
	filePr1 := NewFile("secret.dat", 1000)

	root.Add(fileR1)
	root.Add(docs)
	root.Add(pics)

	docs.Add(fileD1)
	docs.Add(fileD2)
	docs.Add(private)

	pics.Add(fileP1)

	private.Add(filePr1)

	// Check individual sizes
	assert.Equal(t, 10, fileR1.GetSize())
	assert.Equal(t, 100, fileD1.GetSize())
	assert.Equal(t, 200, fileD2.GetSize())
	assert.Equal(t, 500, fileP1.GetSize())
	assert.Equal(t, 1000, filePr1.GetSize())

	// Check directory sizes
	assert.Equal(t, 1000, private.GetSize())
	assert.Equal(t, 500, pics.GetSize())
	assert.Equal(t, 100+200+1000, docs.GetSize())
	assert.Equal(t, 10+(100+200+1000)+500, root.GetSize())
}

// Display method testing is often limited to ensuring no panics,
// unless output capturing is implemented.
func TestDisplay_RunsWithoutPanic(t *testing.T) {
	root := NewDirectory("root")
	docs := NewDirectory("docs")
	file1 := NewFile("file1.txt", 50)
	root.Add(docs)
	docs.Add(file1)

	assert.NotPanics(t, func() {
		root.Display("")
	})
}
