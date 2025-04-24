from document_processor import DocumentProcessor, TextProcessor, JSONProcessor, HTMLProcessor

# --- Client Code ---
def client_code(processor: DocumentProcessor, doc_title: str, doc_content: list[str], output_dir: str, filename: str):
    """Client code that uses the document processor (creator)."""
    processor.process_and_save(doc_title, doc_content, output_dir, filename)

if __name__ == "__main__":
    print("--- Factory Method - Document Processor Demo (Python) ---")

    # Define sample data
    report_title = "Quarterly Report Q1 2025"
    report_content = [
        "This report summarizes the key activities and results for the first quarter.",
        "Sales Performance: Met targets, with significant growth in the North region.",
        "Marketing Campaigns: Launched 'Spring Forward' initiative, results pending.",
        "Product Development: Version 2.1 of the flagship product entered beta testing.",
        "Financial Overview: Stable revenue, slight increase in operational costs."
    ]
    output_directory = "output_files"
    base_filename = "quarterly_report_q1"

    # --- Use different processors to create and save documents ---
    print("\n--- Using TextProcessor ---   ")
    client_code(TextProcessor(), report_title, report_content, output_directory, base_filename)

    print("\n--- Using JSONProcessor ---   ")
    client_code(JSONProcessor(), report_title, report_content, output_directory, base_filename)

    print("\n--- Using HTMLProcessor ---   ")
    client_code(HTMLProcessor(), report_title, report_content, output_directory, base_filename)

    print(f"\n--- Demo Complete ---   ")
    print(f"Check the '{output_directory}' directory for the generated files:")
    print(f"- {base_filename}.txt")
    print(f"- {base_filename}.json")
    print(f"- {base_filename}.html")
