// Define a type for the data structure
type DataRecord = {
  id: number;
  name: string;
  email: string;
};

export abstract class DataExporter {
  /**
   * The template method defining the skeleton of the export algorithm.
   * It calls the primitive operations in a specific order.
   * Returns the final output/status message.
   */
  public exportData(): string {
    const data = this.fetchData();
    const formattedData = this.formatData(data);
    // Optional hook before saving
    this.preSaveHook(formattedData);
    const resultMessage = this.saveData(formattedData);
    // Optional hook after saving
    this.postSaveHook(resultMessage);
    // Return class name along with message for clarity
    return `${this.constructor.name}: ${resultMessage}`;
  }

  /**
   * A concrete method common to all subclasses.
   * In a real scenario, this might involve DB queries, API calls, etc.
   */
  protected fetchData(): DataRecord[] {
    console.log(`${this.constructor.name}: Fetching data...`);
    // Simulate fetching some data
    return [
      { id: 1, name: "Alice", email: "alice@example.com" },
      { id: 2, name: "Bob", email: "bob@example.com" },
    ];
  }

  /**
   * Abstract method for formatting data. Must be implemented by subclasses.
   */
  protected abstract formatData(data: DataRecord[]): string;

  /**
   * Abstract method for saving data. Must be implemented by subclasses.
   * Returns a status message.
   */
  protected abstract saveData(formattedData: string): string;

  // --- Hooks (optional steps with default implementation) ---

  /**
   * Hook before saving data. Default implementation does nothing.
   */
  protected preSaveHook(formattedData: string): void {
    // console.log(`${this.constructor.name}: (Pre-save hook) Data ready for saving.`);
  }

  /**
   * Hook after saving data. Default implementation does nothing.
   */
  protected postSaveHook(resultMessage: string): void {
    // console.log(`${this.constructor.name}: (Post-save hook) Save operation completed with status: ${resultMessage}`);
  }
}

export class CsvExporter extends DataExporter {
  protected formatData(data: DataRecord[]): string {
    console.log(`${this.constructor.name}: Formatting data into CSV...`);
    if (!data || data.length === 0) {
      return "";
    }
    // Ensure consistent order using keys from the first record
    const headers = Object.keys(data[0]!).join(",");
    const rows = data.map((row) => Object.values(row).join(","));
    return `${headers}\n${rows.join("\n")}`;
  }

  protected saveData(formattedData: string): string {
    console.log(`${this.constructor.name}: Saving data as CSV:`);
    console.log("--- CSV START ---");
    console.log(formattedData);
    console.log("--- CSV END ---");
    // Simulate saving to a file
    return "Data successfully saved to output.csv";
  }
}

export class JsonExporter extends DataExporter {
  protected formatData(data: DataRecord[]): string {
    console.log(`${this.constructor.name}: Formatting data into JSON...`);
    return JSON.stringify(data, null, 2); // Pretty print JSON
  }

  protected saveData(formattedData: string): string {
    console.log(`${this.constructor.name}: Saving data as JSON:`);
    console.log("--- JSON START ---");
    console.log(formattedData);
    console.log("--- JSON END ---");
    // Simulate saving to a file
    return "Data successfully saved to output.json";
  }

  // Example of overriding a hook
  protected override preSaveHook(formattedData: string): void {
    console.log(
      `${this.constructor.name}: (Pre-save hook) Validating JSON structure before saving...`
    );
    try {
      JSON.parse(formattedData); // Attempt to parse
      console.log(`${this.constructor.name}: (Pre-save hook) JSON is valid.`);
    } catch (e) {
      // Log error or handle appropriately
      console.error(
        `${this.constructor.name}: (Pre-save hook) Invalid JSON detected: ${
          e instanceof Error ? e.message : String(e)
        }`
      );
      // Optionally throw error to stop the process if validation fails critically
      // throw new Error("Invalid JSON structure detected during pre-save validation.");
    }
  }
}
