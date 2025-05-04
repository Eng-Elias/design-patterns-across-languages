import { CsvExporter, JsonExporter } from "./data_exporter";

function main() {
  console.log("--- TypeScript Template Method Demo: Data Export ---");

  console.log("\nExporting data using CSV Exporter:");
  const csvExporter = new CsvExporter();
  const csvResult = csvExporter.exportData();
  console.log(`\nCSV Export Final Status: ${csvResult}`);

  console.log("\n" + "=".repeat(30) + "\n");

  console.log("Exporting data using JSON Exporter:");
  const jsonExporter = new JsonExporter();
  const jsonResult = jsonExporter.exportData();
  console.log(`\nJSON Export Final Status: ${jsonResult}`);

  console.log("\n--- Demo Finished ---");
}

main();
