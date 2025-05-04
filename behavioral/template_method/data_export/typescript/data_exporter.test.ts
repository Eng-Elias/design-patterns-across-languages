import { DataExporter, CsvExporter, JsonExporter } from "./data_exporter";

// Define DataRecord type here as well or import if shared
type DataRecord = {
  id: number;
  name: string;
  email: string;
};

// Mock implementation to test the template method structure and hooks
class MockExporter extends DataExporter {
  public stepsCalled: string[] = [];
  public fetchedData: DataRecord[] | null = null;
  public formattedDataForSave: string | null = null;
  public saveResultForPostHook: string | null = null;

  protected override fetchData(): DataRecord[] {
    this.stepsCalled.push("fetch");
    this.fetchedData = [{ id: 99, name: "Mock", email: "mock@test.com" }];
    return this.fetchedData;
  }

  protected formatData(data: DataRecord[]): string {
    expect(data).toEqual(this.fetchedData); // Verify data passed correctly
    this.stepsCalled.push("format");
    return "formatted_mock_data";
  }

  protected saveData(formattedData: string): string {
    expect(formattedData).toBe("formatted_mock_data"); // Verify data passed correctly
    this.stepsCalled.push("save");
    this.formattedDataForSave = formattedData;
    this.saveResultForPostHook = "mock_saved_status";
    return this.saveResultForPostHook;
  }

  protected override preSaveHook(formattedData: string): void {
    expect(formattedData).toBe("formatted_mock_data"); // Verify data passed correctly
    this.stepsCalled.push("pre_save");
  }

  protected override postSaveHook(resultMessage: string): void {
    expect(resultMessage).toBe(this.saveResultForPostHook); // Verify data passed correctly
    this.stepsCalled.push("post_save");
  }
}

describe("DataExporter Template Method", () => {
  let consoleSpy: jest.SpyInstance;

  beforeEach(() => {
    // Spy on console.log before each test
    consoleSpy = jest.spyOn(console, "log").mockImplementation();
    // Spy on console.error for the JSON hook test
    jest.spyOn(console, "error").mockImplementation();
  });

  afterEach(() => {
    // Restore original console.log and console.error after each test
    consoleSpy.mockRestore();
    jest.spyOn(console, "error").mockRestore(); // Ensure error spy is also restored
  });

  test("should execute steps in the correct order", () => {
    const mockExporter = new MockExporter();
    const finalStatus = mockExporter.exportData();
    const expectedOrder = ["fetch", "format", "pre_save", "save", "post_save"];

    expect(mockExporter.stepsCalled).toEqual(expectedOrder);
    expect(finalStatus).toBe(
      `MockExporter: ${mockExporter.saveResultForPostHook}`
    );
  });

  test("CsvExporter should format and save data as CSV", () => {
    const csvExporter = new CsvExporter();
    const result = csvExporter.exportData();
    const output = consoleSpy.mock.calls.map((call) => call[0]).join("\n"); // Get console output

    // Check key messages
    expect(output).toContain("CsvExporter: Fetching data...");
    expect(output).toContain("CsvExporter: Formatting data into CSV...");
    expect(output).toContain("CsvExporter: Saving data as CSV:");

    // Check CSV content (basic structure)
    expect(output).toContain("id,name,email");
    expect(output).toContain("1,Alice,alice@example.com");
    expect(output).toContain("2,Bob,bob@example.com");

    // Check final status
    expect(result).toBe("CsvExporter: Data successfully saved to output.csv");
  });

  test("JsonExporter should format and save data as JSON and run hooks", () => {
    const jsonExporter = new JsonExporter();
    const result = jsonExporter.exportData();
    const output = consoleSpy.mock.calls.map((call) => call[0]).join("\n"); // Get console output

    // Check key messages
    expect(output).toContain("JsonExporter: Fetching data...");
    expect(output).toContain("JsonExporter: Formatting data into JSON...");
    // Check overridden hook message
    expect(output).toContain(
      "JsonExporter: (Pre-save hook) Validating JSON structure before saving..."
    );
    expect(output).toContain("JsonExporter: (Pre-save hook) JSON is valid.");
    expect(output).toContain("JsonExporter: Saving data as JSON:");

    // Check JSON content (more robustly than just includes)
    const jsonStartIndex = output.indexOf("--- JSON START ---");
    const jsonEndIndex = output.indexOf("--- JSON END ---");
    expect(jsonStartIndex).toBeGreaterThan(-1);
    expect(jsonEndIndex).toBeGreaterThan(-1);
    const jsonString = output
      .substring(jsonStartIndex + "--- JSON START ---".length, jsonEndIndex)
      .trim();

    let parsedJson;
    try {
      parsedJson = JSON.parse(jsonString);
    } catch (e) {
      // Fail test if JSON parsing fails
      throw new Error(
        `Failed to parse JSON output: ${e}\nOutput:\n${jsonString}`
      );
    }

    const expectedJson = [
      { id: 1, name: "Alice", email: "alice@example.com" },
      { id: 2, name: "Bob", email: "bob@example.com" },
    ];
    expect(parsedJson).toEqual(expectedJson);

    // Check final status
    expect(result).toBe("JsonExporter: Data successfully saved to output.json");
  });

  test("JsonExporter preSaveHook should log error for invalid JSON", () => {
    const jsonExporter = new JsonExporter();
    // Mock formatData to return invalid JSON to test the hook
    jest
      .spyOn(JsonExporter.prototype as any, "formatData")
      .mockReturnValue("invalid json {");

    jsonExporter.exportData(); // Run the export
    const errorOutput = jest
      .spyOn(console, "error")
      .mock.calls.map((call) => call[0])
      .join("\n");

    // Check that the hook logged an error
    expect(errorOutput).toContain(
      "JsonExporter: (Pre-save hook) Invalid JSON detected:"
    );

    // Restore the original formatData method
    jest.spyOn(JsonExporter.prototype as any, "formatData").mockRestore();
  });
});
