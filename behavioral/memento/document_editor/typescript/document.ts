// behavioral/memento/document_editor/typescript/document.ts
import { IMemento, ConcreteMemento } from './memento';

// The Originator holds some important state that may change over time.
export class Document {
    private content: string;

    constructor(initialContent: string = "") {
        this.content = initialContent;
        console.log(`Document initialized with: '${this.content}'`);
    }

    // Appends text to the document content.
    public write(text: string): void {
        console.log(`Writing: '${text}'`);
        this.content += text;
        console.log(`Current content: '${this.content}'`);
    }

    // Sets the document content directly (used for restoring).
    public setContent(content: string): void {
        this.content = content;
        console.log(`Content set to: '${this.content}'`);
    }

    // Returns the current content.
    public getContent(): string {
        return this.content;
    }

    // Saves the current state inside a memento.
    public save(): IMemento {
        const memento = new ConcreteMemento(this.content);
        console.log(`Saving state: '${this.content}'`);
        return memento;
    }

    // Restores the Originator's state from a memento object.
    public restore(memento: IMemento): void {
        this.content = memento.getState();
        console.log(`Restoring state to: '${this.content}'`);
    }
}