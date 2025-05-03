// behavioral/memento/document_editor/typescript/memento.ts

// The Memento interface provides a way to retrieve the memento's metadata.
// It doesn't expose the Originator's state directly.
export interface IMemento {
    getState(): string;
    // Optional: Add metadata like name or date if needed
    // getName(): string;
    // getDate(): Date;
}

// Concrete Memento contains the infrastructure for storing the Originator's state.
export class ConcreteMemento implements IMemento {
    private state: string;
    // private date: Date; // Example metadata

    constructor(state: string) {
        this.state = state;
        // this.date = new Date();
    }

    // The Originator uses this method when restoring its state.
    public getState(): string {
        return this.state;
    }

    // Example metadata getter
    // public getName(): string {
    //     return `${this.date.toISOString().slice(0, 19)} / (${this.state.substring(0, 9)}...)`;
    // }

    // public getDate(): Date {
    //     return this.date;
    // }
}