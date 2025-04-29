import unittest
from file_system import File, Directory, FileSystemComponent

class TestFileSystemComposite(unittest.TestCase):

    def test_file_properties(self):
        """Test basic properties of a File (Leaf)."""
        file = File("test.txt", 100)
        self.assertEqual(file.get_name(), "test.txt")
        self.assertEqual(file.get_size(), 100)

    def test_file_cannot_have_children(self):
        """Verify Leaf nodes raise errors for child management methods."""
        file = File("test.txt", 100)
        another_file = File("another.txt", 50)
        with self.assertRaises(NotImplementedError):
            file.add(another_file)
        with self.assertRaises(NotImplementedError):
            file.remove(another_file)
        with self.assertRaises(NotImplementedError):
            file.get_child(0)

    def test_directory_properties(self):
        """Test basic properties of a Directory (Composite)."""
        directory = Directory("docs")
        self.assertEqual(directory.get_name(), "docs")
        # Initial size should be 0
        self.assertEqual(directory.get_size(), 0)

    def test_directory_add_and_get_child(self):
        """Test adding and retrieving children from a Directory."""
        directory = Directory("docs")
        file1 = File("file1.txt", 50)
        file2 = File("file2.txt", 150)

        directory.add(file1)
        directory.add(file2)

        self.assertEqual(directory.get_child(0), file1)
        self.assertEqual(directory.get_child(1), file2)
        with self.assertRaises(IndexError): # Check bounds
            directory.get_child(2)

    def test_directory_remove_child(self):
        """Test removing children from a Directory."""
        directory = Directory("docs")
        file1 = File("file1.txt", 50)
        file2 = File("file2.txt", 150)

        directory.add(file1)
        directory.add(file2)

        self.assertEqual(directory.get_child(0), file1)
        self.assertEqual(directory.get_child(1), file2)

        directory.remove(file1)
        # file2 should now be at index 0
        self.assertEqual(directory.get_child(0), file2)
        with self.assertRaises(IndexError):
            directory.get_child(1)

        # Test removing non-existent child
        with self.assertRaises(ValueError):
             directory.remove(file1)

    def test_directory_calculate_size_simple(self):
        """Test directory size calculation with only files."""
        directory = Directory("docs")
        file1 = File("file1.txt", 50)
        file2 = File("file2.txt", 150)

        directory.add(file1)
        directory.add(file2)

        self.assertEqual(directory.get_size(), 50 + 150)

    def test_directory_calculate_size_nested(self):
        """Test directory size calculation with nested directories."""
        root = Directory("root")
        docs = Directory("docs")
        pics = Directory("pics")
        private = Directory("private")

        file_r1 = File("root_file.log", 10)
        file_d1 = File("doc1.txt", 100)
        file_d2 = File("doc2.pdf", 200)
        file_p1 = File("pic1.jpg", 500)
        file_pr1 = File("secret.dat", 1000)

        root.add(file_r1)
        root.add(docs)
        root.add(pics)

        docs.add(file_d1)
        docs.add(file_d2)
        docs.add(private)

        pics.add(file_p1)

        private.add(file_pr1)

        # Check individual sizes
        self.assertEqual(file_r1.get_size(), 10)
        self.assertEqual(file_d1.get_size(), 100)
        self.assertEqual(file_d2.get_size(), 200)
        self.assertEqual(file_p1.get_size(), 500)
        self.assertEqual(file_pr1.get_size(), 1000)

        # Check directory sizes
        self.assertEqual(private.get_size(), 1000)
        self.assertEqual(pics.get_size(), 500)
        self.assertEqual(docs.get_size(), 100 + 200 + 1000)
        self.assertEqual(root.get_size(), 10 + (100 + 200 + 1000) + 500)

    # Display method is harder to test assertively for exact output format,
    # but we can ensure it runs without errors.
    def test_display_runs_without_error(self):
        """Ensure the display method runs without raising exceptions."""
        root = Directory("root")
        docs = Directory("docs")
        file1 = File("file1.txt", 50)
        root.add(docs)
        docs.add(file1)
        try:
            root.display()
            # If it reaches here without error, it's a basic pass.
            # For more robust testing, you might capture stdout.
        except Exception as e:
            self.fail(f"display() raised {type(e).__name__} unexpectedly: {e}")

if __name__ == '__main__':
    unittest.main()
