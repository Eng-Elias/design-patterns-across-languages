import unittest
from io import StringIO
import sys
import json
from data_exporter import CsvExporter, JsonExporter, DataExporter

# Mock DataExporter to test template method structure without abstract errors
class MockExporter(DataExporter):
    def __init__(self):
        self.steps_called = []
        # Store data passed between steps for potential verification
        self.fetched_data = None
        self.formatted_data_for_save = None
        self.save_result_for_post_hook = None

    def fetch_data(self) -> list[dict]:
        self.steps_called.append("fetch")
        self.fetched_data = [{"mock_id": 1, "value": "mock_data"}] # Return non-empty list
        return self.fetched_data

    def format_data(self, data: list[dict]) -> str:
        self.steps_called.append("format")
        return "formatted_mock_data"

    def save_data(self, formatted_data: str) -> str:
        self.steps_called.append("save")
        self.formatted_data_for_save = formatted_data # Store for pre_save_hook check
        self.save_result_for_post_hook = "mock_saved_status" # Store for post_save_hook check
        return self.save_result_for_post_hook

    def pre_save_hook(self, formatted_data: str):
         self.steps_called.append("pre_save")

    def post_save_hook(self, result_message: str):
         self.steps_called.append("post_save")

class TestDataExporter(unittest.TestCase):

    def setUp(self):
        # Redirect stdout to capture print statements
        self.held_stdout = sys.stdout
        sys.stdout = StringIO()

    def tearDown(self):
        # Restore stdout
        sys.stdout = self.held_stdout

    def test_template_method_execution_order(self):
        """Tests if the steps in the template method are called in the correct order."""
        mock_exporter = MockExporter()
        final_status = mock_exporter.export_data()
        expected_order = ["fetch", "format", "pre_save", "save", "post_save"]
        self.assertEqual(mock_exporter.steps_called, expected_order,
                         "Template method steps were not executed in the expected order.")
        # Check the final status returned by the template method
        self.assertEqual(final_status, f"MockExporter: {mock_exporter.save_result_for_post_hook}",
                         "Final status returned by export_data is incorrect.")


    def test_csv_exporter_output_and_result(self):
        """Tests the CsvExporter concrete implementation's output and return value."""
        csv_exporter = CsvExporter()
        result = csv_exporter.export_data()
        output = sys.stdout.getvalue() # Capture printed output

        # Check key messages are printed
        self.assertIn("CsvExporter: Fetching data...", output)
        self.assertIn("CsvExporter: Formatting data into CSV...", output)
        self.assertIn("CsvExporter: Saving data as CSV:", output)

        # Check actual CSV content in output
        expected_csv_header = "id,name,email"
        expected_csv_row1 = "1,Alice,alice@example.com"
        expected_csv_row2 = "2,Bob,bob@example.com"
        self.assertIn(expected_csv_header, output)
        self.assertIn(expected_csv_row1, output)
        self.assertIn(expected_csv_row2, output)

        # Check final status message returned
        self.assertEqual(result, "CsvExporter: Data successfully saved to output.csv")

    def test_json_exporter_output_and_result(self):
        """Tests the JsonExporter concrete implementation's output and return value."""
        json_exporter = JsonExporter()
        result = json_exporter.export_data()
        output = sys.stdout.getvalue() # Capture printed output

        # Check key messages are printed
        self.assertIn("JsonExporter: Fetching data...", output)
        self.assertIn("JsonExporter: Formatting data into JSON...", output)
        self.assertIn("JsonExporter: (Pre-save hook) Validating JSON structure", output) # Check hook override
        self.assertIn("JsonExporter: (Pre-save hook) JSON is valid.", output)
        self.assertIn("JsonExporter: Saving data as JSON:", output)

        # Check actual JSON content in output by parsing it
        try:
            # Find the JSON block in the output robustly
            json_start_marker = "--- JSON START ---"
            json_end_marker = "--- JSON END ---"
            start_index = output.find(json_start_marker)
            end_index = output.find(json_end_marker)
            self.assertNotEqual(start_index, -1, "JSON start marker not found in output")
            self.assertNotEqual(end_index, -1, "JSON end marker not found in output")

            json_str = output[start_index + len(json_start_marker):end_index].strip()
            parsed_json = json.loads(json_str)

            # Verify structure and content of parsed JSON
            expected_json = [
                {"id": 1, "name": "Alice", "email": "alice@example.com"},
                {"id": 2, "name": "Bob", "email": "bob@example.com"}
            ]
            self.assertEqual(parsed_json, expected_json, "Parsed JSON content does not match expected structure.")

        except json.JSONDecodeError as e:
            self.fail(f"Failed to parse JSON from output: {e}\nOutput:\n{output}")
        except Exception as e:
             self.fail(f"An unexpected error occurred during JSON verification: {e}\nOutput:\n{output}")

        # Check final status message returned
        self.assertEqual(result, "JsonExporter: Data successfully saved to output.json")

if __name__ == '__main__':
    unittest.main()
