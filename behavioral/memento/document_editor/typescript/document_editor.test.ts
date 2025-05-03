// behavioral/memento/document_editor/typescript/document_editor.test.ts
import { Document } from './document';
import { History } from './history';
// No need to import Memento directly for tests if History manages it

describe('Document Editor Memento Pattern', () => {
    let doc: Document;
    let history: History;

    // Suppress console.log during tests
    let consoleSpy: jest.SpyInstance;
    beforeAll(() => {
        consoleSpy = jest.spyOn(console, 'log').mockImplementation(() => {});
    });
    afterAll(() => {
        consoleSpy.mockRestore();
    });
     // Reset before each test
    beforeEach(() => {
        // Re-enable logging for setup phase if needed, or keep suppressed
        // consoleSpy.mockRestore();
        // consoleSpy = jest.spyOn(console, 'log').mockImplementation(() => {});
    });


    test('should initialize with content and save initial state', () => {
        doc = new Document("Initial");
        history = new History(doc);
        expect(doc.getContent()).toBe("Initial");
        expect(history.getHistoryLength()).toBe(1);
        expect(history.getLastHistoryState()).toBe("Initial");
    });

    test('should save new states correctly', () => {
        doc = new Document();
        history = new History(doc); // Saves initial "" state
        doc.write("First edit");
        history.save();
        expect(doc.getContent()).toBe("First edit");
        expect(history.getHistoryLength()).toBe(2);
        expect(history.getLastHistoryState()).toBe("First edit");
    });

    test('should undo states correctly', () => {
        doc = new Document("A");
        history = new History(doc); // History: ["A"]
        doc.write("B");
        history.save();        // History: ["A", "AB"]
        doc.write("C");
        history.save();        // History: ["A", "AB", "ABC"]

        expect(doc.getContent()).toBe("ABC");
        history.undo();        // Restore "AB". History: ["A", "AB"], Redo: ["ABC"]
        expect(doc.getContent()).toBe("AB");
        expect(history.getHistoryLength()).toBe(2);
        expect(history.getRedoStackLength()).toBe(1);
        expect(history.getLastRedoState()).toBe("ABC");


        history.undo();        // Restore "A". History: ["A"], Redo: ["ABC", "AB"]
        expect(doc.getContent()).toBe("A");
        expect(history.getHistoryLength()).toBe(1);
        expect(history.getRedoStackLength()).toBe(2);

        history.undo();        // Cannot undo further
        expect(doc.getContent()).toBe("A"); // State remains "A"
        expect(history.getHistoryLength()).toBe(1);
        expect(history.getRedoStackLength()).toBe(2); // Redo stack unchanged
    });

    test('should redo states correctly', () => {
        doc = new Document("A");
        history = new History(doc); // History: ["A"]
        doc.write("B");
        history.save();        // History: ["A", "AB"]
        doc.write("C");
        history.save();        // History: ["A", "AB", "ABC"]

        history.undo();        // Back to "AB". History: ["A", "AB"], Redo: ["ABC"]
        history.undo();        // Back to "A". History: ["A"], Redo: ["ABC", "AB"]

        expect(doc.getContent()).toBe("A");
        history.redo();        // Restore "AB". History: ["A", "AB"], Redo: ["ABC"]
        expect(doc.getContent()).toBe("AB");
        expect(history.getHistoryLength()).toBe(2);
        expect(history.getRedoStackLength()).toBe(1);

        history.redo();        // Restore "ABC". History: ["A", "AB", "ABC"], Redo: []
        expect(doc.getContent()).toBe("ABC");
        expect(history.getHistoryLength()).toBe(3);
        expect(history.getRedoStackLength()).toBe(0);

        history.redo();        // Cannot redo further
        expect(doc.getContent()).toBe("ABC"); // State remains "ABC"
        expect(history.getHistoryLength()).toBe(3);
        expect(history.getRedoStackLength()).toBe(0); // Redo stack unchanged
    });

    test('should clear redo stack on new save after undo', () => {
        doc = new Document("A");
        history = new History(doc); // History: ["A"]
        doc.write("B");
        history.save();        // History: ["A", "AB"]
        doc.write("C");
        history.save();        // History: ["A", "AB", "ABC"]

        history.undo();        // Back to "AB". History: ["A", "AB"], Redo: ["ABC"]
        expect(history.getRedoStackLength()).toBe(1);

        // Make a new edit after undo
        doc.write("D"); // Content becomes "ABD"
        history.save();        // History: ["A", "AB", "ABD"], Redo: [] (cleared)
        expect(doc.getContent()).toBe("ABD");
        expect(history.getHistoryLength()).toBe(3);
        expect(history.getRedoStackLength()).toBe(0); // Redo stack should be empty

        // Try redoing - should fail
        history.redo();
        expect(doc.getContent()).toBe("ABD"); // State remains "ABD"
    });
});