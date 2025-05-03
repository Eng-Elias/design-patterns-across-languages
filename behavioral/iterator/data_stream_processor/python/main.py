from data_stream_processor import DataStream

if __name__ == "__main__":
    # Create a data stream
    data_stream = DataStream()

    # Add some data chunks
    data_stream.add_chunk("Chunk 1")
    data_stream.add_chunk({"type": "data", "payload": 123})
    data_stream.add_chunk([1, 2, 3, 4, 5])
    data_stream.add_chunk("Chunk 4")

    # Process the stream using the iterator (implicitly via for loop)
    print("Processing data stream chunks:")
    for chunk in data_stream:
        print(f"  Processing: {chunk}")

    # Demonstrate getting length
    print(f"\nTotal chunks in stream: {len(data_stream)}")

    # Explicit iterator usage (less common in Python)
    print("\nProcessing again with explicit iterator:")
    iterator = iter(data_stream)
    while True:
        try:
            chunk = next(iterator)
            print(f"  Processing: {chunk}")
        except StopIteration:
            break

    # Demonstrate getting a specific chunk
    try:
        third_chunk = data_stream.get(2)
        print(f"\nThe third chunk is: {third_chunk}")
        fourth_chunk = data_stream.get(3)
        print(f"The fourth chunk is: {fourth_chunk}")
        # Try getting out of bounds
        print("\nAttempting to get out-of-bounds chunk:")
        data_stream.get(10)
    except IndexError as e:
        print(f"  Caught expected error: {e}")