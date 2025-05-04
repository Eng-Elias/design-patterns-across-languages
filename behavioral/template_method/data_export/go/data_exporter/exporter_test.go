package data_exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	// Import testify for assertions if you prefer that style, e.g.:
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
	// Remember to run `go mod tidy` if you add imports
)

// --- Mock Implementation for Testing Template Method Order ---
// Simplified: Only records steps, no internal assertions.
type MockExporterImplementation struct {
	StepsCalled       []string
	SaveResultMessage string // Store the message SaveData should return
	BaseImplementation
}

func (m *MockExporterImplementation) GetName() string {
	return "MockExporter"
}

func (m *MockExporterImplementation) FormatData(data []DataRecord) (string, error) {
	m.StepsCalled = append(m.StepsCalled, "format")
	// Return a fixed string, assuming data passed from fetchData is implicitly correct for this test
	return "formatted_mock_data", nil
}

func (m *MockExporterImplementation) SaveData(formattedData string) (string, error) {
	m.StepsCalled = append(m.StepsCalled, "save")
	m.SaveResultMessage = "mock_saved_status" // Set the expected return message
	return m.SaveResultMessage, nil
}

func (m *MockExporterImplementation) PreSaveHook(formattedData string) {
	m.StepsCalled = append(m.StepsCalled, "pre_save")
}

func (m *MockExporterImplementation) PostSaveHook(resultMessage string) {
	m.StepsCalled = append(m.StepsCalled, "post_save")
}

// --- Helper to capture stdout ---
func captureOutput(f func()) string {
	var buf bytes.Buffer
	// Capture both standard logger and direct fmt prints to stdout
	log.SetOutput(&buf)
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // Execute the function that prints

	// Restore
	w.Close()
	os.Stdout = oldStdout
	log.SetOutput(os.Stderr) // Restore default logger output

	// Read captured stdout more reliably
	var stdoutBytes bytes.Buffer
	_, err := stdoutBytes.ReadFrom(r)
	if err != nil {
		// Log or handle the read error if necessary, but don't stop the test
		fmt.Fprintf(os.Stderr, "Error reading from pipe: %v\n", err)
	}
	r.Close()

	// Combine stdout and log buffer
	return stdoutBytes.String() + buf.String()
}

// --- Test Cases ---

// Test the order of operations using the mock
func TestDataExporter_TemplateMethodOrder(t *testing.T) {
	mockImpl := &MockExporterImplementation{}
	exporter := NewDataExporter(mockImpl) // Constructor now sets default fetchFunc

	// Assign the test function directly to the fetchFunc field
	// This is the fix for the "cannot assign" error
	exporter.fetchFunc = func() ([]DataRecord, error) {
		// Simulate fetch step being called conceptually for ordering
		mockImpl.StepsCalled = append(mockImpl.StepsCalled, "fetch")
		return []DataRecord{{ID: 99, Name: "Mock", Email: "mock@test.com"}}, nil // Return dummy data
	}
	// No need for defer as the change is scoped to this instance 'exporter'

	finalStatus, err := exporter.ExportData()

	if err != nil {
		t.Fatalf("ExportData with mock failed: %v", err)
	}

	// Verify the sequence of calls recorded by the mock
	expectedOrder := []string{"fetch", "format", "pre_save", "save", "post_save"}
	if len(mockImpl.StepsCalled) != len(expectedOrder) {
		t.Fatalf("Expected %d steps, got %d. Steps: %v", len(expectedOrder), len(mockImpl.StepsCalled), mockImpl.StepsCalled)
	}
	for i, step := range expectedOrder {
		if mockImpl.StepsCalled[i] != step {
			t.Errorf("Step %d mismatch: expected %q, got %q. Steps: %v", i, step, mockImpl.StepsCalled[i], mockImpl.StepsCalled)
		}
	}

	// Verify the final status message uses the mock's name and result
	expectedStatus := fmt.Sprintf("%s: %s", mockImpl.GetName(), mockImpl.SaveResultMessage)
	if finalStatus != expectedStatus {
		t.Errorf("Final status mismatch: expected %q, got %q", expectedStatus, finalStatus)
	}
}

// Test the real CsvExporter's output and result
func TestCsvExporter(t *testing.T) {
	csvImpl := &CsvExporter{}
	exporter := NewDataExporter(csvImpl) // Uses defaultFetchData via constructor

	var finalStatus string
	var err error

	// Capture printed output
	output := captureOutput(func() {
		finalStatus, err = exporter.ExportData()
	})

	if err != nil {
		t.Fatalf("CsvExporter.ExportData failed: %v", err)
	}

	// Check key messages printed to output (similar to Python's assertIn)
	expectedMessages := []string{
		"CsvExporter: Fetching data...", // This comes from defaultFetchData now
		"CsvExporter: Formatting data into CSV...",
		"CsvExporter: Saving data as CSV:",
		"--- CSV START ---",
		"ID,Name,Email", // Check header
		"1,Alice,alice@example.com", // Check data row 1
		"2,Bob,bob@example.com",   // Check data row 2
		"--- CSV END ---",
	}
	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Expected output to contain %q, but it didn't.\nOutput:\n%s", msg, output)
		}
	}

	// Check final status message returned (similar to Python's assertEqual)
	expectedStatus := "CsvExporter: Data successfully saved to output.csv"
	if finalStatus != expectedStatus {
		t.Errorf("Final status mismatch: expected %q, got %q", expectedStatus, finalStatus)
	}
}

// Test the real JsonExporter's output and result
func TestJsonExporter(t *testing.T) {
	jsonImpl := &JsonExporter{}
	exporter := NewDataExporter(jsonImpl) // Uses defaultFetchData

	var finalStatus string
	var err error

	// Capture printed output
	output := captureOutput(func() {
		finalStatus, err = exporter.ExportData()
	})

	if err != nil {
		t.Fatalf("JsonExporter.ExportData failed: %v", err)
	}

	// Check key messages printed to output
	expectedMessages := []string{
		"JsonExporter: Fetching data...", // From defaultFetchData
		"JsonExporter: Formatting data into JSON...",
		"JsonExporter: (Pre-save hook) Validating JSON structure before saving...", // Hook message
		"JsonExporter: (Pre-save hook) JSON is valid.",                           // Hook message
		"JsonExporter: Saving data as JSON:",
		"--- JSON START ---",
		// Actual JSON content checked below
		"--- JSON END ---",
	}
	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Expected output to contain %q, but it didn't.\nOutput:\n%s", msg, output)
		}
	}

	// Extract and parse JSON content
	startMarker := "--- JSON START ---"
	endMarker := "--- JSON END ---"
	startIndex := strings.Index(output, startMarker)
	endIndex := strings.Index(output, endMarker)
	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		t.Fatalf("Could not find valid JSON START/END markers in output:\n%s", output)
	}

	capturedJsonStr := strings.TrimSpace(output[startIndex+len(startMarker) : endIndex])
	var parsedJson []DataRecord
	if err := json.Unmarshal([]byte(capturedJsonStr), &parsedJson); err != nil {
		t.Fatalf("Failed to parse captured JSON output: %v\nOutput JSON string:\n%s", err, capturedJsonStr)
	}

	// Verify parsed JSON structure and content
	expectedData := []DataRecord{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	// Simple equality check for slices of structs (works if order is guaranteed)
	if len(parsedJson) != len(expectedData) {
        t.Fatalf("Parsed JSON length mismatch: expected %d, got %d.\nParsed JSON:\n%+v", len(expectedData), len(parsedJson), parsedJson)
    }
    for i := range expectedData {
        if parsedJson[i] != expectedData[i] {
            t.Errorf("Parsed JSON element mismatch at index %d: expected %+v, got %+v", i, expectedData[i], parsedJson[i])
        }
    }
    // Or using testify/require for better diffs: require.Equal(t, expectedData, parsedJson)

	// Check final status message returned
	expectedStatus := "JsonExporter: Data successfully saved to output.json"
	if finalStatus != expectedStatus {
		t.Errorf("Final status mismatch: expected %q, got %q", expectedStatus, finalStatus)
	}
}

// Test the JsonExporter's hook when FormatData produces invalid JSON
func TestJsonExporter_InvalidJsonHook(t *testing.T) {
	jsonImpl := &JsonExporter{}
	exporter := NewDataExporter(jsonImpl) // Uses defaultFetchData

	// Temporarily replace FormatData on the *implementation* using the field
	// This is the fix for the "cannot assign" error
	jsonImpl.formatFunc = func(data []DataRecord) (string, error) {
		// We don't need the log message from format here, just the bad data
		return `{"id": 1, "name": "Test", "email": "test@example.com`, nil // Invalid JSON
	}
	// No need for defer here as we are modifying a field on a test-specific instance.
	// If the instance were reused, a defer might be needed.

	// Capture output to check the hook's error message
	output := captureOutput(func() {
		// We don't care about the final result/error here, just the logged message
		_, _ = exporter.ExportData() // Run the export; hook should print error
	})

	// Check that the hook logged the invalid JSON error
	expectedHookErrorMsg := "JsonExporter: (Pre-save hook) Invalid JSON detected:"
	if !strings.Contains(output, expectedHookErrorMsg) {
		t.Errorf("Expected hook error message %q not found in output.\nOutput:\n%s", expectedHookErrorMsg, output)
	}
}
