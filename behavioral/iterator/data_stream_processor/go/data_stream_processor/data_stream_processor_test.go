package data_stream_processor

import (
	"reflect"
	"testing"
)

func TestEmptyStream(t *testing.T) {
	ds := NewDataStream()

	it := ds.CreateIterator()

	if ds.GetCount() != 0 {
		t.Errorf("Expected count 0 for empty stream, got %d", ds.GetCount())
	}

	if it.HasNext() {
		t.Error("Expected HasNext() to be false for empty stream iterator")
	}

	_, err := it.Next()
	if err == nil {
		t.Error("Expected error when calling Next() on empty stream iterator")
	}

	_, err = ds.Get(0)
	if err == nil {
		t.Error("Expected error when calling Get(0) on empty stream")
	}
}

func TestIteration(t *testing.T) {
	ds := NewDataStream()
	// Unified test data
	expectedChunks := []interface{}{"Chunk A", 123, map[string]string{"key": "value"}, []int{10, 20}}
	for _, chunk := range expectedChunks {
		ds.AddChunk(chunk)
	}

	if ds.GetCount() != len(expectedChunks) {
		t.Errorf("Expected count %d, got %d", len(expectedChunks), ds.GetCount())
	}

	it := ds.CreateIterator()
	collectedChunks := []interface{}{}

	for it.HasNext() {
		chunk, err := it.Next()
		if err != nil {
			t.Fatalf("Unexpected error during iteration: %v", err)
		}
		collectedChunks = append(collectedChunks, chunk)
	}

	if !reflect.DeepEqual(collectedChunks, expectedChunks) {
		t.Errorf("Iterated chunks mismatch.\nExpected: %v\nGot:      %v", expectedChunks, collectedChunks)
	}

	// Test HasNext after iteration
	if it.HasNext() {
		t.Error("Expected HasNext() to be false after full iteration")
	}

	// Test Next after iteration
	_, err := it.Next()
	if err == nil {
		t.Error("Expected error when calling Next() after full iteration")
	}
}

func TestGet(t *testing.T) {
	ds := NewDataStream()
	chunks := []interface{}{"first", "second"}
	ds.AddChunk(chunks[0])
	ds.AddChunk(chunks[1])

	// Test valid Get
	val, err := ds.Get(1)
	if err != nil {
		t.Errorf("Unexpected error getting index 1: %v", err)
	}
	if val != chunks[1] {
		t.Errorf("Expected Get(1) to return '%v', got '%v'", chunks[1], val)
	}

	// Test Get out of bounds (negative)
	_, err = ds.Get(-1)
	if err == nil {
		t.Error("Expected error for Get(-1)")
	}

	// Test Get out of bounds (too large)
	_, err = ds.Get(ds.GetCount())
	if err == nil {
		t.Errorf("Expected error for Get(%d)", ds.GetCount())
	}
}