package data_stream_processor

import "errors"

// Iterator Interface defines methods for traversing elements.
type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
}

// Aggregate Interface defines a method for creating an iterator.
type Aggregate interface {
	CreateIterator() Iterator
}

// DataStream holds the collection of data chunks.
type DataStream struct {
	dataChunks []interface{}
}

// NewDataStream creates a new DataStream.
func NewDataStream() *DataStream {
	return &DataStream{
		dataChunks: make([]interface{}, 0),
	}
}

// AddChunk adds a data chunk to the stream.
func (ds *DataStream) AddChunk(chunk interface{}) {
	ds.dataChunks = append(ds.dataChunks, chunk)
}

// Get retrieves a chunk by its index.
func (ds *DataStream) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(ds.dataChunks) {
		return nil, errors.New("index out of bounds")
	}
	return ds.dataChunks[index], nil
}

// GetCount returns the number of chunks in the stream.
func (ds *DataStream) GetCount() int {
	return len(ds.dataChunks)
}

// CreateIterator creates an iterator for the DataStream.
func (ds *DataStream) CreateIterator() Iterator {
	return &StreamIterator{
		stream:   ds,
		position: 0,
	}
}

// StreamIterator implements the Iterator interface for DataStream.
type StreamIterator struct {
	stream   *DataStream
	position int
}

// HasNext checks if there are more elements to iterate.
func (it *StreamIterator) HasNext() bool {
	return it.position < it.stream.GetCount()
}

// Next returns the next element in the stream.
func (it *StreamIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("no more elements")
	}
	value, err := it.stream.Get(it.position)
	if err != nil {
		// This should ideally not happen if HasNext is true, but check anyway
		return nil, err
	}
	it.position++
	return value, nil
}