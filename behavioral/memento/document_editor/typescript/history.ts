// behavioral/memento/document_editor/typescript/history.ts
import { Document } from './document';
import { IMemento } from './memento';

// The Caretaker doesn't depend on the Concrete Memento class.
export class History {
    private history: IMemento[] = [];
    private redoStack: IMemento[] = [];
    private document: Document; // The Originator

    constructor(document: Document) {
        this.document = document;
        // Save initial state
        this.save();
    }

    // Saves the current document state to history.
    public save(): void {
        console.log("History: Saving state...");
        const memento = this.document.save();
        this.history.push(memento);
        // Clear redo stack whenever a new state is saved
        this.redoStack = [];
        console.log("History: State saved.");
    }

    // Restores the previous state.
    public undo(): void {
        if (this.history.length <= 1) { // Need at least initial state + one more
            console.log("History: Cannot undo further.");
            return;
        }

        console.log("History: Undoing...");
        // Pop current state and move it to redo stack
        const currentMemento = this.history.pop();
        if (currentMemento) { // Ensure pop didn't return undefined
             this.redoStack.push(currentMemento);
        }


        // Restore to the previous state (now the last in history)
        const previousMemento = this.history[this.history.length - 1];
        if (previousMemento) { // Ensure history is not empty
             this.document.restore(previousMemento);
             console.log("History: Undo complete.");
        } else {
            console.error("History: Error during undo - history stack is empty unexpectedly.");
        }
    }

    // Restores a previously undone state.
    public redo(): void {
        if (this.redoStack.length === 0) {
            console.log("History: Cannot redo.");
            return;
        }

        console.log("History: Redoing...");
        // Pop the state to redo from the redo stack
        const mementoToRestore = this.redoStack.pop();

        if (mementoToRestore) { // Ensure pop didn't return undefined
            // Restore the document state
            this.document.restore(mementoToRestore);

            // Add the restored state back to the main history
            this.history.push(mementoToRestore);
            console.log("History: Redo complete.");
        } else {
             console.error("History: Error during redo - redo stack is empty unexpectedly.");
        }
    }

    // Prints the states stored in history (for demo).
    public printHistory(): void {
        console.log("History Log:");
        this.history.forEach((memento, index) => {
            console.log(`  ${index}: '${memento.getState()}'`);
        });
        console.log("Redo Stack Log:");
        this.redoStack.forEach((memento, index) => {
             console.log(`  ${index}: '${memento.getState()}'`);
        });
    }

    // Helper for tests
    public getHistoryLength(): number {
        return this.history.length;
    }

    // Helper for tests
    public getRedoStackLength(): number {
        return this.redoStack.length;
    }
     // Helper for tests - Careful: Exposes internal Memento
    public getLastHistoryState(): string | null {
        if (this.history.length === 0) return null;
        return this.history[this.history.length - 1].getState();
    }
     // Helper for tests - Careful: Exposes internal Memento
    public getLastRedoState(): string | null {
        if (this.redoStack.length === 0) return null;
        return this.redoStack[this.redoStack.length - 1].getState();
    }
}