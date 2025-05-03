import { DataStream } from './data_stream_processor';

// Create a data stream (specify the type of data, e.g., any or a specific type)
const dataStream = new DataStream<any>();

// Add some data chunks
dataStream.addChunk("Chunk 1");
dataStream.addChunk({ type: "data", payload: 123 }); // Match payload from python
dataStream.addChunk([1, 2, 3, 4, 5]); // Match array from python
dataStream.addChunk("Chunk 4");

// Process the stream using the iterator
console.log("Processing data stream chunks:");
const iterator = dataStream.createIterator();

while (iterator.hasNext()) {
    const result = iterator.next();
    if (!result.done) {
        const value = result.value;
        // Use JSON.stringify only for objects/arrays for better readability
        const output = typeof value === 'object' && value !== null ? JSON.stringify(value) : value;
        console.log(`  Processing: ${output}`);
    }
}

// Demonstrate getting count
console.log(`\nTotal chunks in stream: ${dataStream.getCount()}`);

// Demonstrate getting a specific chunk
try {
    const thirdChunk = dataStream.get(2);
    const thirdChunkOutput = typeof thirdChunk === 'object' && thirdChunk !== null ? JSON.stringify(thirdChunk) : thirdChunk;
    console.log(`\nThe third chunk is: ${thirdChunkOutput}`);
} catch (error) {
    if (error instanceof Error) {
        console.error(error.message);
    } else {
        console.error('An unknown error occurred when getting a chunk.');
    }
}

// Attempt to get an out-of-bounds chunk
console.log("\nAttempting to get out-of-bounds chunk:");
try {
    dataStream.get(10); // Index likely out of bounds
    console.log("  Did not get expected error for out-of-bounds get()");
} catch (error) {
    if (error instanceof Error) {
        console.log(`  Caught expected error: ${error.message}`);
    } else {
        console.error('  An unknown error occurred when getting out-of-bounds chunk.');
    }
}