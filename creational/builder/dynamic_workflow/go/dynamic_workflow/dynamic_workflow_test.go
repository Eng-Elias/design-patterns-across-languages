package dynamic_workflow

import (
	"bytes" // To capture stdout
	"io"
	"log"
	"os"
	"reflect" // For deep comparison
	"strings"
	"testing"
)

// Helper function to capture stdout
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	// Redirect standard output temporarily
	r, w, _ := os.Pipe()
	origStdout := os.Stdout
	os.Stdout = w

	f() // Execute the function whose output we want to capture

	// Restore standard output
	w.Close()
	os.Stdout = origStdout
	log.SetOutput(os.Stderr) // Restore log output

	out, _ := io.ReadAll(r)
	return string(out)
}

func TestDynamicWorkflowBuilder(t *testing.T) {
	t.Run("BuildEmptyWorkflow", func(t *testing.T) {
		builder := NewDynamicWorkflowBuilder()
		workflow := builder.Build()
		if workflow == nil {
			t.Fatal("Build() returned nil")
		}
		if len(workflow.steps) != 0 {
			t.Errorf("Expected 0 steps, got %d", len(workflow.steps))
		}
	})

	t.Run("AddSingleStep", func(t *testing.T) {
		builder := NewDynamicWorkflowBuilder()
		params := map[string]interface{}{"to": "test@example.com"}
		workflow := builder.AddStep("send_email", params).Build()
		if len(workflow.steps) != 1 {
			t.Fatalf("Expected 1 step, got %d", len(workflow.steps))
		}
		if workflow.steps[0].Type != "send_email" {
			t.Errorf("Expected step type 'send_email', got '%s'", workflow.steps[0].Type)
		}
		if !reflect.DeepEqual(workflow.steps[0].Params, params) {
			t.Errorf("Expected params %v, got %v", params, workflow.steps[0].Params)
		}
	})

	t.Run("AddMultipleStepsChaining", func(t *testing.T) {
		builder := NewDynamicWorkflowBuilder()
		params1 := map[string]interface{}{"path": "/script1"}
		params2 := map[string]interface{}{"channel": "#test"}
		workflow := builder.
			AddStep("run_script", params1).
			AddStep("notify_slack", params2).
			Build()

		if len(workflow.steps) != 2 {
			t.Fatalf("Expected 2 steps, got %d", len(workflow.steps))
		}
		if workflow.steps[0].Type != "run_script" {
			t.Errorf("Expected step 1 type 'run_script', got '%s'", workflow.steps[0].Type)
		}
		if !reflect.DeepEqual(workflow.steps[0].Params, params1) {
			t.Errorf("Expected step 1 params %v, got %v", params1, workflow.steps[0].Params)
		}
		if workflow.steps[1].Type != "notify_slack" {
			t.Errorf("Expected step 2 type 'notify_slack', got '%s'", workflow.steps[1].Type)
		}
		if !reflect.DeepEqual(workflow.steps[1].Params, params2) {
			t.Errorf("Expected step 2 params %v, got %v", params2, workflow.steps[1].Params)
		}
	})

	t.Run("AddStepReturnsSelf", func(t *testing.T) {
		builder := NewDynamicWorkflowBuilder()
		ret := builder.AddStep("test", nil)
		if ret != builder {
			t.Error("AddStep did not return the builder instance")
		}
	})
}

func TestDynamicWorkflowExecution(t *testing.T) {

	t.Run("ExecuteEmptyWorkflow", func(t *testing.T) {
		workflow := NewDynamicWorkflow([]WorkflowStep{}) // Empty steps
		output := captureOutput(workflow.Execute)

		if !strings.Contains(output, "Workflow has no steps") {
			t.Errorf("Expected output to contain 'Workflow has no steps', got '%s'", output)
		}
	})

	t.Run("ExecuteKnownSteps", func(t *testing.T) {
		steps := []WorkflowStep{
			{Type: "send_email", Params: map[string]interface{}{"to": "exec@test.com", "subject": "Exec Test"}},
			{Type: "run_script", Params: map[string]interface{}{"path": "/exec/test.py"}},
			{Type: "notify_slack", Params: map[string]interface{}{"channel": "#exec", "message": "Done"}},
		}
		workflow := NewDynamicWorkflow(steps)
		output := captureOutput(workflow.Execute)

		if !strings.Contains(output, "Simulating sending email to 'exec@test.com'") {
			t.Error("Output missing expected email simulation")
		}
		if !strings.Contains(output, "Simulating running script: '/exec/test.py'") {
			t.Error("Output missing expected script simulation")
		}
		if !strings.Contains(output, "Simulating notifying Slack channel '#exec'") {
			t.Error("Output missing expected Slack simulation")
		}
		if strings.Contains(output, "Unknown step type") {
			t.Error("Output unexpectedly contains 'Unknown step type'")
		}
		if strings.Contains(output, "Error executing step") {
			t.Error("Output unexpectedly contains 'Error executing step'")
		}
	})

	t.Run("ExecuteUnknownStep", func(t *testing.T) {
		steps := []WorkflowStep{
			{Type: "send_email", Params: map[string]interface{}{}},
			{Type: "warp_drive_engage", Params: map[string]interface{}{"speed": 9}},
			{Type: "run_script", Params: map[string]interface{}{}},
		}
		workflow := NewDynamicWorkflow(steps)
		output := captureOutput(workflow.Execute)

		if !strings.Contains(output, "Simulating sending email") {
			t.Error("Output missing expected email simulation")
		}
		if !strings.Contains(output, "⚠️ Unknown step type: 'warp_drive_engage'. Skipping.") {
			t.Error("Output missing expected unknown step warning")
		}
		if !strings.Contains(output, "Simulating running script") {
			t.Error("Output missing expected script simulation")
		}
		if strings.Contains(output, "Error executing step") {
			t.Error("Output unexpectedly contains 'Error executing step'")
		}
	})

	t.Run("HandlerParametersUsed", func(t *testing.T) {
		steps := []WorkflowStep{
			{Type: "notify_slack", Params: map[string]interface{}{"channel": "#specific-go", "message": "Go Specific Message!"}},
		}
		workflow := NewDynamicWorkflow(steps)
		output := captureOutput(workflow.Execute)

		if !strings.Contains(output, "Simulating notifying Slack channel '#specific-go': 'Go Specific Message!...'" ) {
			t.Errorf("Output missing expected Slack simulation with specific params. Got: %s", output)
		}
	})
}
