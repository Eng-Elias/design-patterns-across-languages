// behavioral/memento/document_editor/typescript/main.ts
import { Document } from './document';
import { History } from './history';

function main() {
    console.log("--- Memento Pattern Document Editor Demo ---");

    // Create Originator (Document)
    const doc = new Document();

    // Create Caretaker (History), passing the originator
    const history = new History(doc);

    // User starts typing and saving states
    doc.write("Hello");
    history.save();
    history.printHistory();

    doc.write(" World");
    history.save();
    history.printHistory();

    doc.write("!");
    history.save();
    history.printHistory();

    // Undo operations
    console.log("\n--- Undoing ---");
    history.undo(); // Undo "!"
    history.printHistory();

    history.undo(); // Undo " World"
    history.printHistory();

    // Redo operation
    console.log("\n--- Redoing ---");
    history.redo(); // Redo " World"
    history.printHistory();

    // Make another change
    console.log("\n--- Making another change ---");
    doc.write(", How are you?");
    history.save();
    history.printHistory();

    // Try to redo (should fail as redo stack was cleared)
    console.log("\n--- Trying Redo again ---");
    history.redo();
    history.printHistory();

    // Undo back to the beginning
    console.log("\n--- Undoing to the start ---");
    history.undo(); // Undo ", How are you?"
    history.undo(); // Undo " World"
    history.undo(); // Undo "Hello"
    history.printHistory();

    // Try to undo further (should fail)
    console.log("\n--- Trying Undo again ---");
    history.undo();
    history.printHistory();

    console.log("\n--- Demo Complete ---");
}

main();