package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"factory_method_document_processor_go/document_processor"
)

func main() {
	fmt.Println("--- Running Go Factory Method (Document Processor - Refactored) ---")

	// Define sample data
	title := "Quarterly Report Q1 2025"
	content := []string{
		"This report summarizes the key activities and results for the first quarter.",
		"Sales Performance: Met targets, with significant growth in the North region.",
		"Marketing Campaigns: Launched 'Spring Forward' initiative, results pending.",
		"Product Development: Version 2.1 of the flagship product entered beta testing.",
		"Financial Overview: Stable revenue, slight increase in operational costs.",
	}
	outputDir := "output_files"
	filenameBase := "quarterly_report_q1"

	// Ensure output directory exists
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create output directory '%s': %v", outputDir, err)
	}
	fmt.Printf("Ensured output directory exists: %s\n", outputDir)

	// Create processor instances
	processors := []document_processor.DocumentProcessor{
		&document_processor.TextProcessor{},
		&document_processor.JSONProcessor{},
		&document_processor.HTMLProcessor{},
	}

	// Process using each processor
	for i, processor := range processors {
		// Determine the type name for logging
		var processorType string
		switch processor.(type) {
		case *document_processor.TextProcessor:
			processorType = "TextProcessor"
		case *document_processor.JSONProcessor:
			processorType = "JSONProcessor"
		case *document_processor.HTMLProcessor:
			processorType = "HTMLProcessor"
		default:
			processorType = "Unknown Processor"
		}

		fmt.Printf("\n--- Using %s ---\n", processorType)
		// Adjust filename base slightly for uniqueness per type
		currentFilenameBase := fmt.Sprintf("%s_%d", filenameBase, i)
		
		doc, err := processor.ProcessDocument(title, content, outputDir, currentFilenameBase)
		if err != nil {
			log.Printf("ERROR processing with %s: %v", processorType, err)
			continue // Move to the next processor
		}

		// Construct expected path to report success
		var expectedExt string
		switch doc.(type) {
		case *document_processor.TextDocument:
			expectedExt = ".txt"
		case *document_processor.JSONDocument:
			expectedExt = ".json"
		case *document_processor.HTMLDocument:
			expectedExt = ".html"
		}
		expectedPath := filepath.Join(outputDir, currentFilenameBase+expectedExt)
		fmt.Printf("Successfully processed and saved document: %s\n", expectedPath)
	}

	fmt.Println("\n--- Go Demo Complete ---")
}
