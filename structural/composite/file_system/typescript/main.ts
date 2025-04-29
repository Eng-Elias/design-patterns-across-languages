import { File, Directory } from "./file_system";

function main(): void {
  /**
   * Demonstrates the Composite pattern using a file system example.
   */

  // Create some files (Leaf objects)
  const file1 = new File("document.txt", 1024); // 1 KB
  const file2 = new File("image.jpg", 5120); // 5 KB
  const file3 = new File("archive.zip", 10240); // 10 KB
  const file4 = new File("report.pdf", 2048); // 2 KB

  // Create some directories (Composite objects)
  const root = new Directory("root");
  const documentsDir = new Directory("Documents");
  const picturesDir = new Directory("Pictures");
  const privateDir = new Directory("Private");

  // Build the file system tree structure
  root.add(documentsDir);
  root.add(picturesDir);

  documentsDir.add(file1);
  documentsDir.add(file4);
  documentsDir.add(privateDir); // Add a subdirectory

  picturesDir.add(file2);

  privateDir.add(file3); // Add a file to the subdirectory

  // --- Demonstrate Uniform Treatment --- //

  console.log("--- Displaying the entire file system structure ---");
  root.display();

  console.log("\n--- Calculating sizes ---");

  // Calculate size of the entire root directory
  console.log(`Total size of '${root.getName()}': ${root.getSize()} bytes`);

  // Calculate size of a subdirectory
  console.log(
    `Total size of '${documentsDir.getName()}': ${documentsDir.getSize()} bytes`
  );

  // Calculate size of an individual file (Leaf)
  console.log(`Size of '${file1.getName()}': ${file1.getSize()} bytes`);

  // Display a specific subdirectory
  console.log("\n--- Displaying the 'Documents' directory ---");
  documentsDir.display();
}

main();
