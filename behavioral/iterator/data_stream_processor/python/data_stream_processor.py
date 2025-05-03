from collections.abc import Iterator, Iterable


class StreamIterator(Iterator):
    """Concrete Iterator for the DataStream."""
    _position: int = None
    _stream: 'DataStream' = None

    def __init__(self, stream: 'DataStream'):
        self._stream = stream
        self._position = 0

    def __next__(self):
        """Returns the next element from the stream."""
        try:
            value = self._stream.get(self._position)
            self._position += 1
            return value
        except IndexError:
            raise StopIteration()


class DataStream(Iterable):
    """Concrete Aggregate that stores data chunks and provides an iterator."""

    def __init__(self):
        self._data_chunks = []

    def add_chunk(self, chunk: any) -> None:
        """Adds a data chunk to the stream."""
        self._data_chunks.append(chunk)

    def get(self, index: int) -> any:
        """Gets a data chunk at a specific index."""
        return self._data_chunks[index]

    def __len__(self) -> int:
        """Returns the total number of chunks in the stream."""
        return len(self._data_chunks)

    def __iter__(self) -> StreamIterator:
        """Returns the iterator for the stream."""
        return StreamIterator(self)