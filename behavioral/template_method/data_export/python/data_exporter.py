from abc import ABC, abstractmethod
import json

class DataExporter(ABC):
    """
    Abstract base class defining the template method for exporting data.
    """

    def export_data(self) -> str:
        """
        The template method defining the skeleton of the export algorithm.
        It calls the primitive operations in a specific order.
        Returns the final output/status message.
        """
        data = self.fetch_data()
        formatted_data = self.format_data(data)
        # Optional hook before saving
        self.pre_save_hook(formatted_data)
        result_message = self.save_data(formatted_data)
        # Optional hook after saving
        self.post_save_hook(result_message)
        # Return class name along with message for clarity
        return f"{self.__class__.__name__}: {result_message}"

    def fetch_data(self) -> list[dict]:
        """
        A concrete method common to all subclasses.
        In a real scenario, this might involve DB queries, API calls, etc.
        """
        print(f"{self.__class__.__name__}: Fetching data...")
        # Simulate fetching some data
        return [
            {"id": 1, "name": "Alice", "email": "alice@example.com"},
            {"id": 2, "name": "Bob", "email": "bob@example.com"}
        ]

    @abstractmethod
    def format_data(self, data: list[dict]) -> str:
        """
        Abstract method for formatting data. Must be implemented by subclasses.
        """
        pass

    @abstractmethod
    def save_data(self, formatted_data: str) -> str:
        """
        Abstract method for saving data. Must be implemented by subclasses.
        Returns a status message.
        """
        pass

    # --- Hooks (optional steps with default implementation) ---
    def pre_save_hook(self, formatted_data: str):
        """Hook before saving data. Default implementation does nothing."""
        # print(f"{self.__class__.__name__}: (Pre-save hook) Data ready for saving.")
        pass

    def post_save_hook(self, result_message: str):
        """Hook after saving data. Default implementation does nothing."""
        # print(f"{self.__class__.__name__}: (Post-save hook) Save operation completed with status: {result_message}")
        pass


class CsvExporter(DataExporter):
    """
    Concrete subclass for exporting data to CSV format.
    """
    def format_data(self, data: list[dict]) -> str:
        print(f"{self.__class__.__name__}: Formatting data into CSV...")
        if not data:
            return ""
        # Ensure consistent order of columns based on the first item's keys
        header = ",".join(data[0].keys())
        rows = [",".join(map(str, [item[key] for key in data[0].keys()])) for item in data]
        return header + "\n" + "\n".join(rows)

    def save_data(self, formatted_data: str) -> str:
        print(f"{self.__class__.__name__}: Saving data as CSV:")
        print("--- CSV START ---")
        print(formatted_data)
        print("--- CSV END ---")
        # Simulate saving to a file
        return "Data successfully saved to output.csv"

class JsonExporter(DataExporter):
    """
    Concrete subclass for exporting data to JSON format.
    """
    def format_data(self, data: list[dict]) -> str:
        print(f"{self.__class__.__name__}: Formatting data into JSON...")
        return json.dumps(data, indent=2)

    def save_data(self, formatted_data: str) -> str:
        print(f"{self.__class__.__name__}: Saving data as JSON:")
        print("--- JSON START ---")
        print(formatted_data)
        print("--- JSON END ---")
        # Simulate saving to a file
        return "Data successfully saved to output.json"

    # Example of overriding a hook
    def pre_save_hook(self, formatted_data: str):
        print(f"{self.__class__.__name__}: (Pre-save hook) Validating JSON structure before saving...")
        try:
            json.loads(formatted_data)
            print(f"{self.__class__.__name__}: (Pre-save hook) JSON is valid.")
        except json.JSONDecodeError as e:
            print(f"{self.__class__.__name__}: (Pre-save hook) Invalid JSON detected: {e}")
            # Optionally raise an error or handle it
