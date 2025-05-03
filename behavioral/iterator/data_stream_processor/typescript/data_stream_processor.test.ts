import { DataStream } from './data_stream_processor';

describe('DataStream Processor - Iterator Pattern', () => {

  it('should handle an empty stream correctly', () => {
    const stream = new DataStream<any>();
    const iterator = stream.createIterator();

    expect(iterator.hasNext()).toBe(false);
    expect(iterator.next()).toEqual({ value: undefined, done: true });
    expect(stream.getCount()).toBe(0);
  });

  it('should iterate through all elements in the stream', () => {
    const stream = new DataStream<any>();
    const chunks = ["Chunk A", 123, { key: "value" }, [10, 20]];
    chunks.forEach(chunk => stream.addChunk(chunk));

    expect(stream.getCount()).toBe(chunks.length);

    const iterator = stream.createIterator();
    const collectedChunks: any[] = [];

    while (iterator.hasNext()) {
      const result = iterator.next();
      if (!result.done) {
        collectedChunks.push(result.value);
      }
    }

    expect(collectedChunks).toEqual(chunks);
    expect(iterator.hasNext()).toBe(false); // Iterator should be exhausted
    expect(iterator.next()).toEqual({ value: undefined, done: true }); // Calling next again
  });

  it('should correctly report hasNext', () => {
    const stream = new DataStream<string>();
    stream.addChunk("first");
    const iterator = stream.createIterator();

    expect(iterator.hasNext()).toBe(true);
    iterator.next(); // Consume the first element
    expect(iterator.hasNext()).toBe(false);
  });

  it('should retrieve elements using the get method', () => {
    const stream = new DataStream<string>();
    const chunk1 = "First";
    const chunk2 = "Second";
    stream.addChunk(chunk1);
    stream.addChunk(chunk2);

    expect(stream.get(0)).toBe(chunk1);
    expect(stream.get(1)).toBe(chunk2);
  });

  it('should throw an error when getting an element with an invalid index', () => {
    const stream = new DataStream<number>();
    stream.addChunk(1);

    expect(() => stream.get(1)).toThrow("Index out of bounds");
    expect(() => stream.get(-1)).toThrow("Index out of bounds");
  });
});