import { DynamicWorkflowBuilder } from "./dynamic_workflow";

function main() {
  console.log("--- Builder Pattern - Dynamic Workflow Demo (TypeScript) ---");

  // Create a workflow using the builder with method chaining
  const builder = new DynamicWorkflowBuilder();

  const workflow = builder
    .addStep("send_email", {
      to: "dev@example.com",
      subject: "Build Started",
      body: "Starting the main build process...",
    })
    .addStep("run_script", { path: "/scripts/compile.sh" })
    .addStep("run_script", { path: "/scripts/test.sh" })
    .addStep("notify_slack", {
      channel: "#builds",
      message: "Build and Tests successful!",
    })
    .addStep("unknown_step", { parameter: "some_value" }) // Test unknown step
    .addStep("send_email", {
      to: "qa@example.com",
      subject: "Build Ready for QA",
    }) // Missing body
    .build();

  console.log("\n--- Workflow Constructed ---");

  // Execute the constructed workflow
  workflow.execute();

  console.log("\n--- Building and executing a second, simpler workflow --- ");
  const workflow2 = new DynamicWorkflowBuilder() // Create a new builder instance
    .addStep("notify_slack", {
      channel: "#monitoring",
      message: "System check initiated.",
    })
    .build();

  workflow2.execute();

  console.log("\n--- Demo Complete ---");
}

// Run the main function
main();
