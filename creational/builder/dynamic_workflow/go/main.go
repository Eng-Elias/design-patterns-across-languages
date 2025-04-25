package main

import (
	"builder_pattern_dynamic_workflow_go/dynamic_workflow"
	"fmt"
)

func main() {
	fmt.Println("--- Builder Pattern - Dynamic Workflow Demo (Go) ---")

	// Create a workflow using the builder with method chaining
	builder := dynamic_workflow.NewDynamicWorkflowBuilder()

	workflow := builder.
		AddStep("send_email", map[string]interface{}{
			"to":      "dev@example.com",
			"subject": "Build Started",
			"body":    "Starting the main build process...",
		}).
		AddStep("run_script", map[string]interface{}{
			"path": "/scripts/compile.sh",
		}).
		AddStep("run_script", map[string]interface{}{
			"path": "/scripts/test.sh",
		}).
		AddStep("notify_slack", map[string]interface{}{
			"channel": "#builds",
			"message": "Build and Tests successful!",
		}).
		AddStep("unknown_step", map[string]interface{}{
			"parameter": "some_value",
		}). // Test unknown step
		AddStep("send_email", map[string]interface{}{
			"to":      "qa@example.com",
			"subject": "Build Ready for QA",
		}). // Missing body
		Build()

	fmt.Println("\n--- Workflow Constructed ---")

	// Execute the constructed workflow
	workflow.Execute()

	fmt.Println("\n--- Building and executing a second, simpler workflow --- ")
	workflow2 := dynamic_workflow.NewDynamicWorkflowBuilder().
		AddStep("notify_slack", map[string]interface{}{
			"channel": "#monitoring",
			"message": "System check initiated.",
		}).
		Build()

	workflow2.Execute()

	fmt.Println("\n--- Demo Complete ---")
}
