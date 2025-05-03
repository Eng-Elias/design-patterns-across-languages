import unittest
from data_stream_processor import DataStream


class TestDataStreamProcessor(unittest.TestCase):

    def test_empty_stream(self):
        stream = DataStream()
        iterator = iter(stream)
        with self.assertRaises(StopIteration):
            next(iterator)
        self.assertEqual(len(stream), 0)

    def test_iteration(self):
        stream = DataStream()
        chunks = ["Chunk A", 123, {"key": "value"}, [10, 20]]
        for chunk in chunks:
            stream.add_chunk(chunk)

        self.assertEqual(len(stream), len(chunks))

        # Test iteration using for loop
        collected_chunks = []
        for chunk in stream:
            collected_chunks.append(chunk)
        self.assertEqual(collected_chunks, chunks)

        # Test explicit iterator usage
        iterator = iter(stream)
        explicit_collected = []
        try:
            while True:
                explicit_collected.append(next(iterator))
        except StopIteration:
            pass
        self.assertEqual(explicit_collected, chunks)

    def test_get_method(self):
        stream = DataStream()
        chunk1 = "First"
        chunk2 = "Second"
        stream.add_chunk(chunk1)
        stream.add_chunk(chunk2)
        self.assertEqual(stream.get(0), chunk1)
        self.assertEqual(stream.get(1), chunk2)
        with self.assertRaises(IndexError):
            stream.get(2)


if __name__ == '__main__':
    unittest.main()