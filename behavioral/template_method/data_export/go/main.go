package main

import (
	"fmt"
	"log"

	// Use the correct relative path based on your go.mod module path
	"template_method_pattern_data_export_go/data_exporter"
)

func main() {
	fmt.Println("--- Go Template Method Demo: Data Export ---")

	// Create concrete implementation instances
	csvImpl := &data_exporter.CsvExporter{}
	jsonImpl := &data_exporter.JsonExporter{}

	// Create DataExporter instances using the implementations
	csvExporter := data_exporter.NewDataExporter(csvImpl)
	jsonExporter := data_exporter.NewDataExporter(jsonImpl)

	fmt.Println("Exporting data using CSV Exporter:")
	csvResult, err := csvExporter.ExportData()
	if err != nil {
		log.Fatalf("CSV Export failed: %v", err)
	}
	fmt.Printf("CSV Export Final Status: %s\n", csvResult)

	fmt.Println("\n" + "=============================") // Separator
	fmt.Println("Exporting data using JSON Exporter:")
	jsonResult, err := jsonExporter.ExportData()
	if err != nil {
		log.Fatalf("JSON Export failed: %v", err)
	}
	fmt.Printf("JSON Export Final Status: %s\n", jsonResult)

	fmt.Println("\n--- Demo Finished ---")
}
