package main

import (
	"fmt"
	"iterator_pattern_data_stream_processor_go/data_stream_processor"
	"log"
)

func main() {
	// Create a DataStream
	dataStream := data_stream_processor.NewDataStream()

	// Add chunks of data (using interface{} for mixed types) - Unified Data
	dataStream.AddChunk("Chunk 1")
	dataStream.AddChunk(map[string]interface{}{"type": "data", "payload": 123})
	dataStream.AddChunk([]int{1, 2, 3, 4, 5})
	dataStream.AddChunk("Chunk 4")

	// Get an iterator
	it := dataStream.CreateIterator()

	// Iterate through the stream
	fmt.Println("Iterating through the stream:")
	for it.HasNext() {
		chunk, err := it.Next()
		if err != nil {
			log.Printf("Error getting next chunk: %v", err) // Log error but continue if possible
			break                                            // Stop if Next returns an error
		}
		// Use type assertion or reflection for specific handling if needed
		// For demonstration, just print the value and its type
		fmt.Printf("  Processing chunk: %v (Type: %T)\n", chunk, chunk)
	}

	fmt.Printf("Data stream created with %d chunks.\n\n", dataStream.GetCount())

	// Demonstrate getting a specific chunk
	fmt.Println("\nDemonstrating Get():")
	indexToGet := 2
	chunk, err := dataStream.Get(indexToGet)
	if err != nil {
		fmt.Printf("  Error getting chunk at index %d: %v\n", indexToGet, err)
	} else {
		fmt.Printf("  Chunk at index %d: %v (Type: %T)\n", indexToGet, chunk, chunk)
	}

	// Attempt to get an out-of-bounds chunk
	fmt.Println("\nAttempting to get out-of-bounds chunk:")
	indexOutOfBounds := dataStream.GetCount() + 1
	_, err = dataStream.Get(indexOutOfBounds)
	if err != nil {
		fmt.Printf("  Caught expected error getting index %d: %v\n", indexOutOfBounds, err)
	} else {
		fmt.Println("  Did not get expected error for out-of-bounds Get()")
	}
}