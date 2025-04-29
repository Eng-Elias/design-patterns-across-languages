from file_system import File, Directory

def main():
    """Demonstrates the Composite pattern using a file system example."""

    # Create some files (Leaf objects)
    file1 = File("document.txt", 1024) # 1 KB
    file2 = File("image.jpg", 5120)    # 5 KB
    file3 = File("archive.zip", 10240) # 10 KB
    file4 = File("report.pdf", 2048)   # 2 KB

    # Create some directories (Composite objects)
    root = Directory("root")
    documents_dir = Directory("Documents")
    pictures_dir = Directory("Pictures")
    private_dir = Directory("Private")

    # Build the file system tree structure
    root.add(documents_dir)
    root.add(pictures_dir)

    documents_dir.add(file1)
    documents_dir.add(file4)
    documents_dir.add(private_dir) # Add a subdirectory

    pictures_dir.add(file2)

    private_dir.add(file3) # Add a file to the subdirectory

    # --- Demonstrate Uniform Treatment --- #

    print("--- Displaying the entire file system structure ---")
    root.display()

    print("\n--- Calculating sizes ---")

    # Calculate size of the entire root directory
    print(f"Total size of '{root.get_name()}': {root.get_size()} bytes")

    # Calculate size of a subdirectory
    print(f"Total size of '{documents_dir.get_name()}': {documents_dir.get_size()} bytes")

    # Calculate size of an individual file (Leaf)
    print(f"Size of '{file1.get_name()}': {file1.get_size()} bytes")

    # Display a specific subdirectory
    print("\n--- Displaying the 'Documents' directory ---")
    documents_dir.display()

if __name__ == "__main__":
    main()
