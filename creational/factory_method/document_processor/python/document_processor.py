from abc import ABC, abstractmethod
import json
import os
from typing import List, Dict, Any

# --- Product Interface ---
class Document(ABC):
    """Abstract Product: Defines the interface for documents."""
    def __init__(self, title: str, content: List[str]):
        self.title = title
        self.content = content
        print(f"Document created: Title='{self.title}', Content preview='{content[0] if content else ''}...'")

    @abstractmethod
    def save(self, path: str):
        """Saves the document to the specified path in its format."""
        pass

# --- Concrete Products ---
class TextDocument(Document):
    """Concrete Product: Implements a plain text document."""
    def save(self, path: str):
        full_path = f"{path}.txt"
        print(f"Saving Text document to: {full_path}")
        try:
            with open(full_path, 'w', encoding='utf-8') as f:
                f.write(f"Title: {self.title}\n")
                f.write("=" * len(self.title) + "==\n\n")
                for line in self.content:
                    f.write(f"{line}\n")
            print("Text document saved successfully.")
        except IOError as e:
            print(f"Error saving text document {full_path}: {e}")

class JSONDocument(Document):
    """Concrete Product: Implements a JSON document."""
    def save(self, path: str):
        full_path = f"{path}.json"
        print(f"Saving JSON document to: {full_path}")
        data: Dict[str, Any] = {
            "title": self.title,
            "content": self.content
        }
        try:
            with open(full_path, 'w', encoding='utf-8') as f:
                json.dump(data, f, indent=4)
            print("JSON document saved successfully.")
        except IOError as e:
            print(f"Error saving JSON document {full_path}: {e}")
        except TypeError as e:
            print(f"Error serializing data to JSON for {full_path}: {e}")

class HTMLDocument(Document):
    """Concrete Product: Implements a simple HTML document."""
    def save(self, path: str):
        full_path = f"{path}.html"
        print(f"Saving HTML document to: {full_path}")
        html_content = f"""<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{self.title}</title>
    <style>
        body {{ font-family: sans-serif; line-height: 1.6; padding: 20px; }}
        h1 {{ color: #333; }}
        p {{ margin-bottom: 10px; }}
    </style>
</head>
<body>
    <h1>{self.title}</h1>
"""
        for line in self.content:
            # Basic paragraph wrapping, escape HTML chars if needed in real use
            html_content += f"    <p>{line}</p>\n"
        html_content += """</body>
</html>
"""
        try:
            with open(full_path, 'w', encoding='utf-8') as f:
                f.write(html_content)
            print("HTML document saved successfully.")
        except IOError as e:
            print(f"Error saving HTML document {full_path}: {e}")

# --- Creator Interface ---
class DocumentProcessor(ABC):
    """Abstract Creator: Declares the factory method."""

    @abstractmethod
    def create_document(self, title: str, content: List[str]) -> Document:
        """The Factory Method."""
        pass

    def process_and_save(self, title: str, content: List[str], output_dir: str, filename_base: str):
        """Processes data and uses the factory method to save the document."""
        # Ensure output directory exists
        try:
            os.makedirs(output_dir, exist_ok=True)
        except OSError as e:
            print(f"Error creating directory {output_dir}: {e}")
            return

        # Create the document using the factory method
        doc = self.create_document(title=title, content=content)

        # Construct the full path without extension
        full_path_base = os.path.join(output_dir, filename_base)

        # Save the document (save method adds the extension)
        print(f"\nProcessing with {self.__class__.__name__}:")
        doc.save(full_path_base)

# --- Concrete Creators ---
class TextProcessor(DocumentProcessor):
    """Concrete Creator: Overrides the factory method for Text documents."""
    def create_document(self, title: str, content: List[str]) -> TextDocument:
        print("TextProcessor: Creating TextDocument.")
        return TextDocument(title, content)

class JSONProcessor(DocumentProcessor):
    """Concrete Creator: Overrides the factory method for JSON documents."""
    def create_document(self, title: str, content: List[str]) -> JSONDocument:
        print("JSONProcessor: Creating JSONDocument.")
        return JSONDocument(title, content)

class HTMLProcessor(DocumentProcessor):
    """Concrete Creator: Overrides the factory method for HTML documents."""
    def create_document(self, title: str, content: List[str]) -> HTMLDocument:
        print("HTMLProcessor: Creating HTMLDocument.")
        return HTMLDocument(title, content)
