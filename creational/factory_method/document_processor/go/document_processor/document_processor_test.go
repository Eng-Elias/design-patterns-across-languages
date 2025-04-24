package document_processor

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTextProcessor(t *testing.T) {
	processor := &TextProcessor{}
	title := "Go Text Test"
	content := []string{"Line 1", "Line 2 for text"}
	filenameBase := "test_text_doc"
	outputDir := t.TempDir() // Create a temporary directory for test output

	// Process and save
	doc, err := processor.ProcessDocument(title, content, outputDir, filenameBase)

	// Assertions
	if err != nil {
		t.Fatalf("ProcessDocument failed: %v", err)
	}
	if doc == nil {
		t.Fatal("ProcessDocument returned nil document")
	}

	// Check document type (optional, more for interface satisfaction check)
	if _, ok := doc.(*TextDocument); !ok {
		t.Errorf("Expected document type *TextDocument, got %T", doc)
	}

	// Check file existence
	expectedFilePath := filepath.Join(outputDir, filenameBase+".txt")
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected file %s was not created", expectedFilePath)
	}

	// Check file content
	fileBytes, err := os.ReadFile(expectedFilePath)
	if err != nil {
		t.Fatalf("Failed to read created file %s: %v", expectedFilePath, err)
	}
	fileContent := string(fileBytes)

	expectedTitleLine := "Title: " + title
	if !strings.Contains(fileContent, expectedTitleLine) {
		t.Errorf("File content missing expected title line '%s'. Got:\n%s", expectedTitleLine, fileContent)
	}
	if !strings.Contains(fileContent, content[0]) {
		t.Errorf("File content missing expected content '%s'. Got:\n%s", content[0], fileContent)
	}
	if !strings.Contains(fileContent, content[1]) {
		t.Errorf("File content missing expected content '%s'. Got:\n%s", content[1], fileContent)
	}
}

func TestJSONProcessor(t *testing.T) {
	processor := &JSONProcessor{}
	title := "Go JSON Test"
	content := []string{"JSON Line 1", "JSON Line 2"}
	filenameBase := "test_json_doc"
	outputDir := t.TempDir()

	doc, err := processor.ProcessDocument(title, content, outputDir, filenameBase)

	if err != nil {
		t.Fatalf("ProcessDocument failed: %v", err)
	}
	if _, ok := doc.(*JSONDocument); !ok {
		t.Errorf("Expected document type *JSONDocument, got %T", doc)
	}

	expectedFilePath := filepath.Join(outputDir, filenameBase+".json")
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected file %s was not created", expectedFilePath)
	}

	fileBytes, err := os.ReadFile(expectedFilePath)
	if err != nil {
		t.Fatalf("Failed to read created file %s: %v", expectedFilePath, err)
	}

	var decodedDoc JSONDocument
	err = json.Unmarshal(fileBytes, &decodedDoc)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON content from %s: %v", expectedFilePath, err)
	}

	if decodedDoc.Title != title {
		t.Errorf("Expected JSON title '%s', got '%s'", title, decodedDoc.Title)
	}
	if len(decodedDoc.Content) != len(content) {
		t.Errorf("Expected JSON content length %d, got %d", len(content), len(decodedDoc.Content))
	} else {
		for i, line := range content {
			if decodedDoc.Content[i] != line {
				t.Errorf("Expected JSON content line %d to be '%s', got '%s'", i, line, decodedDoc.Content[i])
			}
		}
	}
}

func TestHTMLProcessor(t *testing.T) {
	processor := &HTMLProcessor{}
	title := "Go HTML Test"
	content := []string{"HTML First Line <br>", "Second line & stuff"}
	filenameBase := "test_html_doc"
	outputDir := t.TempDir()

	doc, err := processor.ProcessDocument(title, content, outputDir, filenameBase)

	if err != nil {
		t.Fatalf("ProcessDocument failed: %v", err)
	}
	if _, ok := doc.(*HTMLDocument); !ok {
		t.Errorf("Expected document type *HTMLDocument, got %T", doc)
	}

	expectedFilePath := filepath.Join(outputDir, filenameBase+".html")
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected file %s was not created", expectedFilePath)
	}

	fileBytes, err := os.ReadFile(expectedFilePath)
	if err != nil {
		t.Fatalf("Failed to read created file %s: %v", expectedFilePath, err)
	}
	fileContent := string(fileBytes)

	expectedTitleTag := "<title>" + title + "</title>"
	expectedH1Tag := "<h1>" + title + "</h1>"
	expectedP1Tag := "<p>" + content[0] + "</p>"
	expectedP2Tag := "<p>" + content[1] + "</p>"

	if !strings.Contains(fileContent, expectedTitleTag) {
		t.Errorf("File content missing expected tag '%s'. Got:\n%s", expectedTitleTag, fileContent)
	}
	if !strings.Contains(fileContent, expectedH1Tag) {
		t.Errorf("File content missing expected tag '%s'. Got:\n%s", expectedH1Tag, fileContent)
	}
	if !strings.Contains(fileContent, expectedP1Tag) {
		t.Errorf("File content missing expected tag '%s'. Got:\n%s", expectedP1Tag, fileContent)
	}
	if !strings.Contains(fileContent, expectedP2Tag) {
		t.Errorf("File content missing expected tag '%s'. Got:\n%s", expectedP2Tag, fileContent)
	}
}
