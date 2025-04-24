import * as fs from "fs";
import * as path from "path";

// --- Product Interface ---
interface Document {
  title: string;
  content: string[];
  save(filePathBase: string): Promise<void>; // Changed signature, now async
}

// --- Concrete Products ---
class TextDocument implements Document {
  constructor(public title: string, public content: string[]) {
    console.log(
      `TextDocument created: Title='${this.title}', Content preview='${
        content[0] ?? ""
      }...'`
    );
  }

  async save(filePathBase: string): Promise<void> {
    const fullPath = `${filePathBase}.txt`;
    console.log(`Saving Text document to: ${fullPath}`);
    const fileContent = `Title: ${this.title}\n${"=".repeat(
      this.title.length + 2
    )}\n\n${this.content.join("\n")}`;
    try {
      await fs.promises.writeFile(fullPath, fileContent, "utf-8");
      console.log("Text document saved successfully.");
    } catch (err) {
      console.error(`Error saving text document ${fullPath}:`, err);
      throw err; // Re-throw error for testability/handling
    }
  }
}

class JSONDocument implements Document {
  constructor(public title: string, public content: string[]) {
    console.log(
      `JSONDocument created: Title='${this.title}', Content preview='${
        content[0] ?? ""
      }...'`
    );
  }

  async save(filePathBase: string): Promise<void> {
    const fullPath = `${filePathBase}.json`;
    console.log(`Saving JSON document to: ${fullPath}`);
    const data = {
      title: this.title,
      content: this.content,
    };
    try {
      await fs.promises.writeFile(
        fullPath,
        JSON.stringify(data, null, 4),
        "utf-8"
      );
      console.log("JSON document saved successfully.");
    } catch (err) {
      console.error(`Error saving JSON document ${fullPath}:`, err);
      throw err;
    }
  }
}

class HTMLDocument implements Document {
  constructor(public title: string, public content: string[]) {
    console.log(
      `HTMLDocument created: Title='${this.title}', Content preview='${
        content[0] ?? ""
      }...'`
    );
  }

  async save(filePathBase: string): Promise<void> {
    const fullPath = `${filePathBase}.html`;
    console.log(`Saving HTML document to: ${fullPath}`);
    let htmlContent = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>${this.title}</title>
    <style>
        body { font-family: sans-serif; line-height: 1.6; padding: 20px; }
        h1 { color: #333; }
        p { margin-bottom: 10px; }
    </style>
</head>
<body>
    <h1>${this.title}</h1>
`;
    this.content.forEach((line) => {
      // Basic paragraph wrapping, consider escaping HTML chars in real use
      htmlContent += `    <p>${line}</p>\n`;
    });
    htmlContent += `</body>\n</html>\n`;

    try {
      await fs.promises.writeFile(fullPath, htmlContent, "utf-8");
      console.log("HTML document saved successfully.");
    } catch (err) {
      console.error(`Error saving HTML document ${fullPath}:`, err);
      throw err;
    }
  }
}

// --- Creator Interface ---
abstract class DocumentProcessor {
  // The Factory Method
  public abstract createDocument(title: string, content: string[]): Document;

  // Core logic using the factory method
  public async processAndSave(
    title: string,
    content: string[],
    outputDir: string,
    filenameBase: string
  ): Promise<void> {
    // Ensure output directory exists
    try {
      await fs.promises.mkdir(outputDir, { recursive: true });
    } catch (err) {
      console.error(`Error creating directory ${outputDir}:`, err);
      return; // Exit if directory creation fails
    }

    // Create the document using the factory method
    const doc = this.createDocument(title, content);

    // Construct the full path without extension
    const fullPathBase = path.join(outputDir, filenameBase);

    // Save the document
    console.log(`\nProcessing with ${this.constructor.name}:`);
    try {
      await doc.save(fullPathBase);
    } catch (err) {
      console.error(
        `Failed to save document ${filenameBase} using ${this.constructor.name}.`
      );
      // Decide if processing should stop or continue
    }
  }
}

// --- Concrete Creators ---
class TextProcessor extends DocumentProcessor {
  public createDocument(title: string, content: string[]): TextDocument {
    console.log("TextProcessor: Creating TextDocument.");
    return new TextDocument(title, content);
  }
}

class JSONProcessor extends DocumentProcessor {
  public createDocument(title: string, content: string[]): JSONDocument {
    console.log("JSONProcessor: Creating JSONDocument.");
    return new JSONDocument(title, content);
  }
}

class HTMLProcessor extends DocumentProcessor {
  public createDocument(title: string, content: string[]): HTMLDocument {
    console.log("HTMLProcessor: Creating HTMLDocument.");
    return new HTMLDocument(title, content);
  }
}

// Export classes and interfaces
export {
  Document,
  TextDocument,
  JSONDocument,
  HTMLDocument,
  DocumentProcessor,
  TextProcessor,
  JSONProcessor,
  HTMLProcessor,
};
