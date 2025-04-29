import { File, Directory, FileSystemComponent } from "./file_system";

describe("FileSystem Composite Tests", () => {
  describe("File (Leaf)", () => {
    let file: File;

    beforeEach(() => {
      file = new File("test.txt", 100);
    });

    test("should have correct name and size", () => {
      expect(file.getName()).toBe("test.txt");
      expect(file.getSize()).toBe(100);
    });

    test("should throw error when trying to add child", () => {
      const anotherFile = new File("another.txt", 50);
      expect(() => file.add(anotherFile)).toThrow(
        "Cannot add to this component"
      );
    });

    test("should throw error when trying to remove child", () => {
      const anotherFile = new File("another.txt", 50);
      // Need to add first to attempt removal, which itself will fail
      // So we test the method directly
      expect(() => file.remove(anotherFile)).toThrow(
        "Cannot remove from this component"
      );
    });

    test("should throw error when trying to get child", () => {
      expect(() => file.getChild(0)).toThrow(
        "Cannot get child from this component"
      );
    });
  });

  describe("Directory (Composite)", () => {
    let directory: Directory;
    let file1: File;
    let file2: File;

    beforeEach(() => {
      directory = new Directory("docs");
      file1 = new File("file1.txt", 50);
      file2 = new File("file2.txt", 150);
    });

    test("should have correct name and initial size 0", () => {
      expect(directory.getName()).toBe("docs");
      expect(directory.getSize()).toBe(0);
    });

    test("should add and get children correctly", () => {
      directory.add(file1);
      directory.add(file2);
      expect(directory.getChild(0)).toBe(file1);
      expect(directory.getChild(1)).toBe(file2);
      expect(() => directory.getChild(2)).toThrow(RangeError); // Check bounds
    });

    test("should remove children correctly", () => {
      directory.add(file1);
      directory.add(file2);
      expect(directory.getChild(0)).toBe(file1);
      expect(directory.getChild(1)).toBe(file2);

      directory.remove(file1); // Remove first child
      expect(directory.getChild(0)).toBe(file2); // Second child is now first
      expect(() => directory.getChild(1)).toThrow(RangeError);

      // Test removing non-existent child
      expect(() => directory.remove(file1)).toThrow(
        "Component 'file1.txt' not found in directory 'docs'"
      );
    });

    test("should calculate size correctly with only files", () => {
      directory.add(file1);
      directory.add(file2);
      expect(directory.getSize()).toBe(50 + 150);
    });

    test("should calculate size correctly with nested directories", () => {
      const root = new Directory("root");
      const docs = new Directory("docs");
      const pics = new Directory("pics");
      const privateDir = new Directory("private");

      const file_r1 = new File("root_file.log", 10);
      const file_d1 = new File("doc1.txt", 100);
      const file_d2 = new File("doc2.pdf", 200);
      const file_p1 = new File("pic1.jpg", 500);
      const file_pr1 = new File("secret.dat", 1000);

      root.add(file_r1);
      root.add(docs);
      root.add(pics);

      docs.add(file_d1);
      docs.add(file_d2);
      docs.add(privateDir);

      pics.add(file_p1);

      privateDir.add(file_pr1);

      // Check individual sizes
      expect(file_r1.getSize()).toBe(10);
      expect(file_d1.getSize()).toBe(100);
      expect(file_d2.getSize()).toBe(200);
      expect(file_p1.getSize()).toBe(500);
      expect(file_pr1.getSize()).toBe(1000);

      // Check directory sizes
      expect(privateDir.getSize()).toBe(1000);
      expect(pics.getSize()).toBe(500);
      expect(docs.getSize()).toBe(100 + 200 + 1000);
      expect(root.getSize()).toBe(10 + (100 + 200 + 1000) + 500);
    });

    test("display method should run without errors", () => {
      // Mock console.log to prevent output during tests (optional)
      const consoleSpy = jest
        .spyOn(console, "log")
        .mockImplementation(() => {});

      const root = new Directory("root");
      const docs = new Directory("docs");
      const file1 = new File("file1.txt", 50);
      root.add(docs);
      docs.add(file1);

      expect(() => root.display()).not.toThrow();

      // Restore console.log
      consoleSpy.mockRestore();
    });
  });
});
