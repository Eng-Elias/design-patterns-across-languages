from data_exporter import CsvExporter, JsonExporter

def main():
    print("--- Python Template Method Demo: Data Export ---")

    print("\nExporting data using CSV Exporter:")
    csv_exporter = CsvExporter()
    csv_result = csv_exporter.export_data()
    print(f"\nCSV Export Final Status: {csv_result}")

    print("\n" + "="*30 + "\n")

    print("Exporting data using JSON Exporter:")
    json_exporter = JsonExporter()
    json_result = json_exporter.export_data()
    print(f"\nJSON Export Final Status: {json_result}")

    print("\n--- Demo Finished ---")

if __name__ == "__main__":
    main()
