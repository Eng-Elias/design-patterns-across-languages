package data_exporter

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// DataRecord represents the structure of our data (similar to TS/Python)
type DataRecord struct {
	ID    int    `json:"id"`    // Use struct tags for JSON marshalling
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ExporterImplementation defines the interface for the varying steps
// that concrete exporters must implement.
type ExporterImplementation interface {
	FormatData(data []DataRecord) (string, error)
	SaveData(formattedData string) (string, error)
	PreSaveHook(formattedData string)
	PostSaveHook(resultMessage string)
	GetName() string // To get the type name for logging
}

// DataExporter is the 'template' struct. It holds the implementation
// and defines the overall algorithm.
type DataExporter struct {
	Impl      ExporterImplementation
	fetchFunc func() ([]DataRecord, error)
}

// NewDataExporter creates a new DataExporter with a specific implementation.
func NewDataExporter(impl ExporterImplementation) *DataExporter {
	if impl == nil {
		// Or handle this error more gracefully depending on requirements
		log.Fatal("ExporterImplementation cannot be nil")
	}
	de := &DataExporter{Impl: impl}
	de.fetchFunc = de.defaultFetchData
	return de
}

// ExportData is the template method.
func (de *DataExporter) ExportData() (string, error) {
	data, err := de.fetchFunc()
	if err != nil {
		return "", fmt.Errorf("%s: error fetching data: %w", de.Impl.GetName(), err)
	}

	formattedData, err := de.Impl.FormatData(data)
	if err != nil {
		return "", fmt.Errorf("%s: error formatting data: %w", de.Impl.GetName(), err)
	}

	// Optional hook before saving
	de.Impl.PreSaveHook(formattedData)

	resultMessage, err := de.Impl.SaveData(formattedData)
	if err != nil {
		return "", fmt.Errorf("%s: error saving data: %w", de.Impl.GetName(), err)
	}

	// Optional hook after saving
	de.Impl.PostSaveHook(resultMessage)

	// Return class name along with message for clarity
	return fmt.Sprintf("%s: %s", de.Impl.GetName(), resultMessage), nil
}

// defaultFetchData - the original fetchData method, renamed
func (de *DataExporter) defaultFetchData() ([]DataRecord, error) {
	fmt.Printf("%s: Fetching data...\n", de.Impl.GetName())
	// Simulate fetching data
	return []DataRecord{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}, nil // Simulate success
}

// --- Concrete Implementations ---

// BaseImplementation provides default empty hooks
type BaseImplementation struct{}

func (b *BaseImplementation) PreSaveHook(formattedData string) {}
func (b *BaseImplementation) PostSaveHook(resultMessage string) {}

// CsvExporter implements ExporterImplementation for CSV format.
type CsvExporter struct {
	BaseImplementation // Embed base for default hooks
}

func (c *CsvExporter) GetName() string {
	return "CsvExporter"
}

func (c *CsvExporter) FormatData(data []DataRecord) (string, error) {
	fmt.Printf("%s: Formatting data into CSV...\n", c.GetName())
	if len(data) == 0 {
		return "", nil // Return empty string for no data
	}

	var buffer bytes.Buffer                // Use a buffer to write CSV data
	writer := csv.NewWriter(&buffer)

	// Write header dynamically based on the first record's fields
	// (Requires reflection or manual definition if struct fields vary)
	// For simplicity, assuming fixed fields: ID, Name, Email
	header := []string{"ID", "Name", "Email"}
	if err := writer.Write(header); err != nil {
		return "", fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Write rows
	for _, record := range data {
		row := []string{fmt.Sprint(record.ID), record.Name, record.Email}
		if err := writer.Write(row); err != nil {
			// Log intermediate errors?
			continue // Skip bad rows or handle differently
		}
	}

	writer.Flush() // Ensure all data is written to the buffer

	if err := writer.Error(); err != nil {
		return "", fmt.Errorf("error during CSV writing: %w", err)
	}

	// Return the buffer contents as a string, removing trailing newline if any
	return strings.TrimSpace(buffer.String()), nil
}

func (c *CsvExporter) SaveData(formattedData string) (string, error) {
	fmt.Printf("%s: Saving data as CSV:\n", c.GetName())
	fmt.Println("--- CSV START ---")
	fmt.Println(formattedData)
	fmt.Println("--- CSV END ---")
	// Simulate saving to a file
	return "Data successfully saved to output.csv", nil
}

// JsonExporter implements ExporterImplementation for JSON format.
type JsonExporter struct {
	BaseImplementation // Embed base
	formatFunc func(data []DataRecord) (string, error) // Add function field
}

func (j *JsonExporter) GetName() string {
	return "JsonExporter"
}

// FormatData implements the interface method by calling the function field.
func (j *JsonExporter) FormatData(data []DataRecord) (string, error) {
	if j.formatFunc == nil { // Lazy initialization
		j.formatFunc = j.defaultFormatData
	}
	return j.formatFunc(data) // Call the field
}

// defaultFormatData contains the original formatting logic.
func (j *JsonExporter) defaultFormatData(data []DataRecord) (string, error) {
	fmt.Printf("%s: Formatting data into JSON...\n", j.GetName())
	jsonData, err := json.MarshalIndent(data, "", "  ") // Use MarshalIndent for pretty printing
	if err != nil {
		return "", fmt.Errorf("failed to marshal data to JSON: %w", err)
	}
	return string(jsonData), nil
}

// SaveData simulates saving JSON data.
func (j *JsonExporter) SaveData(formattedData string) (string, error) {
	fmt.Printf("%s: Saving data as JSON:\n", j.GetName())
	fmt.Println("--- JSON START ---")
	fmt.Println(formattedData)
	fmt.Println("--- JSON END ---")
	// Simulate writing to file or DB
	return "Data successfully saved to output.json", nil
}

// PreSaveHook provides a specific pre-save step for JSON.
func (j *JsonExporter) PreSaveHook(formattedData string) {
	fmt.Printf("%s: (Pre-save hook) Validating JSON structure before saving...\n", j.GetName())
	var js json.RawMessage
	if err := json.Unmarshal([]byte(formattedData), &js); err != nil {
		fmt.Printf("%s: (Pre-save hook) Invalid JSON detected: %v\n", j.GetName(), err)
		// In a real app, might return error or handle differently
	} else {
		fmt.Printf("%s: (Pre-save hook) JSON is valid.\n", j.GetName())
	}
}

// PostSaveHook is inherited from BaseImplementation for JSON.
