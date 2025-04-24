import unittest
from io import StringIO
import sys
import os
import shutil
import json
from document_processor import (
    TextProcessor, JSONProcessor, HTMLProcessor
)

class TestDocumentProcessorFactoryMethod(unittest.TestCase):

    TEST_OUTPUT_DIR = "temp_test_output"
    TEST_FILENAME_BASE = "test_doc"
    TEST_TITLE = "Test Document Title"
    TEST_CONTENT = ["Line 1 of content.", "Line 2: Another line."]

    def setUp(self):
        """Create a temporary directory for test output files."""
        os.makedirs(self.TEST_OUTPUT_DIR, exist_ok=True)

    def tearDown(self):
        """Remove the temporary directory and its contents after tests."""
        if os.path.exists(self.TEST_OUTPUT_DIR):
            shutil.rmtree(self.TEST_OUTPUT_DIR)

    def test_text_processor(self):
        """Test TextProcessor creates and saves a correct .txt file."""
        processor = TextProcessor()
        processor.process_and_save(self.TEST_TITLE, self.TEST_CONTENT, self.TEST_OUTPUT_DIR, self.TEST_FILENAME_BASE)

        expected_filepath = os.path.join(self.TEST_OUTPUT_DIR, f"{self.TEST_FILENAME_BASE}.txt")
        self.assertTrue(os.path.exists(expected_filepath), f"File not found: {expected_filepath}")

        # Read and verify content
        try:
            with open(expected_filepath, 'r', encoding='utf-8') as f:
                lines = f.readlines()
            self.assertIn(f"Title: {self.TEST_TITLE}\n", lines[0])
            self.assertIn(self.TEST_CONTENT[0], lines[3]) # Check first line of content
            self.assertIn(self.TEST_CONTENT[1], lines[4]) # Check second line of content
        except IOError as e:
            self.fail(f"Failed to read or verify text file {expected_filepath}: {e}")

    def test_json_processor(self):
        """Test JSONProcessor creates and saves a correct .json file."""
        processor = JSONProcessor()
        processor.process_and_save(self.TEST_TITLE, self.TEST_CONTENT, self.TEST_OUTPUT_DIR, self.TEST_FILENAME_BASE)

        expected_filepath = os.path.join(self.TEST_OUTPUT_DIR, f"{self.TEST_FILENAME_BASE}.json")
        self.assertTrue(os.path.exists(expected_filepath), f"File not found: {expected_filepath}")

        # Read and verify content
        try:
            with open(expected_filepath, 'r', encoding='utf-8') as f:
                data = json.load(f)
            self.assertIsInstance(data, dict)
            self.assertEqual(data.get("title"), self.TEST_TITLE)
            self.assertEqual(data.get("content"), self.TEST_CONTENT)
        except (IOError, json.JSONDecodeError) as e:
            self.fail(f"Failed to read or verify JSON file {expected_filepath}: {e}")

    def test_html_processor(self):
        """Test HTMLProcessor creates and saves a correct .html file."""
        processor = HTMLProcessor()
        processor.process_and_save(self.TEST_TITLE, self.TEST_CONTENT, self.TEST_OUTPUT_DIR, self.TEST_FILENAME_BASE)

        expected_filepath = os.path.join(self.TEST_OUTPUT_DIR, f"{self.TEST_FILENAME_BASE}.html")
        self.assertTrue(os.path.exists(expected_filepath), f"File not found: {expected_filepath}")

        # Read and verify content (basic checks)
        try:
            with open(expected_filepath, 'r', encoding='utf-8') as f:
                html_content = f.read()
            self.assertIn(f"<title>{self.TEST_TITLE}</title>", html_content)
            self.assertIn(f"<h1>{self.TEST_TITLE}</h1>", html_content)
            self.assertIn(f"<p>{self.TEST_CONTENT[0]}</p>", html_content)
            self.assertIn(f"<p>{self.TEST_CONTENT[1]}</p>", html_content)
        except IOError as e:
            self.fail(f"Failed to read or verify HTML file {expected_filepath}: {e}")

if __name__ == '__main__':
    unittest.main()
