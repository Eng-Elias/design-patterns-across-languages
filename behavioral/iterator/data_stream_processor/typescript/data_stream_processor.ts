// Iterator Interface
export interface IIterator<T> {
    next(): { value: T | undefined; done: boolean };
    hasNext(): boolean;
}

// Aggregate Interface
export interface IAggregate<T> {
    createIterator(): IIterator<T>;
}

// Concrete Iterator
export class StreamIterator<T> implements IIterator<T> {
    private stream: DataStream<T>;
    private position: number = 0;

    constructor(stream: DataStream<T>) {
        this.stream = stream;
    }

    next(): { value: T | undefined; done: boolean } {
        if (!this.hasNext()) {
            return { value: undefined, done: true };
        }
        const value = this.stream.get(this.position);
        this.position++;
        return { value: value, done: false };
    }

    hasNext(): boolean {
        return this.position < this.stream.getCount();
    }
}

// Concrete Aggregate
export class DataStream<T> implements IAggregate<T> {
    private dataChunks: T[] = [];

    addChunk(chunk: T): void {
        this.dataChunks.push(chunk);
    }

    get(index: number): T {
        if (index < 0 || index >= this.dataChunks.length) {
            throw new Error("Index out of bounds");
        }
        return this.dataChunks[index];
    }

    getCount(): number {
        return this.dataChunks.length;
    }

    createIterator(): IIterator<T> {
        return new StreamIterator<T>(this);
    }
}