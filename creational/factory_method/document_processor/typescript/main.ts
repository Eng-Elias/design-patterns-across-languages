import {
  DocumentProcessor,
  TextProcessor,
  JSONProcessor,
  HTMLProcessor,
} from "./document_processor";

// --- Client Code ---
async function clientCode(
  processor: DocumentProcessor,
  docTitle: string,
  docContent: string[],
  outputDir: string,
  filename: string
): Promise<void> {
  console.log(
    `Client: Using ${processor.constructor.name} to process a document.`
  );
  await processor.processAndSave(docTitle, docContent, outputDir, filename);
}

// --- Main Execution (Async) ---
async function main() {
  console.log("--- Factory Method - Document Processor Demo (TypeScript) ---");

  // Define sample data
  const reportTitle = "Quarterly Report Q1 2025";
  const reportContent = [
    "This report summarizes the key activities and results for the first quarter.",
    "Sales Performance: Met targets, with significant growth in the North region.",
    "Marketing Campaigns: Launched 'Spring Forward' initiative, results pending.",
    "Product Development: Version 2.1 of the flagship product entered beta testing.",
    "Financial Overview: Stable revenue, slight increase in operational costs.",
  ];
  const outputDirectory = "output_files";
  const baseFilename = "quarterly_report_q1";

  // --- Use different processors ---
  // Note: processAndSave now handles directory creation and saving

  console.log("\n--- Using TextProcessor ---   ");
  await clientCode(
    new TextProcessor(),
    reportTitle,
    reportContent,
    outputDirectory,
    baseFilename
  );

  console.log("\n--- Using JSONProcessor ---   ");
  await clientCode(
    new JSONProcessor(),
    reportTitle,
    reportContent,
    outputDirectory,
    baseFilename
  );

  console.log("\n--- Using HTMLProcessor ---   ");
  await clientCode(
    new HTMLProcessor(),
    reportTitle,
    reportContent,
    outputDirectory,
    baseFilename
  );

  console.log("\n--- Demo Complete ---   ");
  console.log(
    `Check the '${outputDirectory}' directory for the generated files:`
  );
  console.log(`- ${baseFilename}.txt`);
  console.log(`- ${baseFilename}.json`);
  console.log(`- ${baseFilename}.html`);
}

// Run the async main function
main().catch((error) => {
  console.error("An error occurred during execution:", error);
});
