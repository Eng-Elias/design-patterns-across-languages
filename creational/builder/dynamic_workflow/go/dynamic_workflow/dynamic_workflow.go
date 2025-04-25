package dynamic_workflow

import (
	"encoding/json" // Used for cleaner printing of params map
	"fmt"
)

// WorkflowStep defines the structure for a single step in the workflow.
// Using interface{} for params value to allow different data types.
type WorkflowStep struct {
	Type   string
	Params map[string]interface{}
}

// StepHandler defines the function signature for step handlers.
type StepHandler func(params map[string]interface{}) error

// DynamicWorkflow represents the complex object (Product) being built.
// It holds the sequence of steps and the logic to execute them.
type DynamicWorkflow struct {
	steps        []WorkflowStep
	stepHandlers map[string]StepHandler
}

// NewDynamicWorkflow creates a new workflow instance.
// It initializes the handlers map.
func NewDynamicWorkflow(steps []WorkflowStep) *DynamicWorkflow {
	w := &DynamicWorkflow{
		// Create a defensive copy of the steps slice
		steps:        append([]WorkflowStep(nil), steps...),
		stepHandlers: make(map[string]StepHandler),
	}
	// Register handlers
	w.stepHandlers["send_email"] = w.handleSendEmail
	w.stepHandlers["run_script"] = w.handleRunScript
	w.stepHandlers["notify_slack"] = w.handleNotifySlack
	// Add more handlers here
	return w
}

// --- Private Handler Methods ---

func (w *DynamicWorkflow) handleSendEmail(params map[string]interface{}) error {
	to, _ := params["to"].(string)       // Use type assertion, ignore error for simplicity
	subject, _ := params["subject"].(string)
	body, _ := params["body"].(string)
	if to == "" {
		to = "<default_email>"
	}
	if subject == "" {
		subject = "<default_subject>"
	}
	bodyDisplay := body
	if len(bodyDisplay) > 30 {
		bodyDisplay = bodyDisplay[:30] + "..."
	}
	fmt.Printf("📧 Simulating sending email to '%s' with subject '%s': '%s'\n", to, subject, bodyDisplay)
	return nil // Indicate success
}

func (w *DynamicWorkflow) handleRunScript(params map[string]interface{}) error {
	path, _ := params["path"].(string)
	if path == "" {
		path = "<default_script_path>"
	}
	fmt.Printf("⚙️ Simulating running script: '%s'\n", path)
	return nil // Indicate success
}

func (w *DynamicWorkflow) handleNotifySlack(params map[string]interface{}) error {
	channel, _ := params["channel"].(string)
	message, _ := params["message"].(string)
	if channel == "" {
		channel = "#general"
	}
	messageDisplay := message
	if len(messageDisplay) > 30 {
		messageDisplay = messageDisplay[:30] + "..."
	}
	fmt.Printf("💬 Simulating notifying Slack channel '%s': '%s'\n", channel, messageDisplay)
	return nil // Indicate success
}

// --- Public Execution Method ---

// Execute runs the steps defined in the workflow sequentially.
func (w *DynamicWorkflow) Execute() {
	fmt.Println("--- Executing Workflow ---")
	if len(w.steps) == 0 {
		fmt.Println("Workflow has no steps.")
		return
	}

	for i, step := range w.steps {
		paramsJSON, _ := json.Marshal(step.Params) // For prettier printing
		fmt.Printf("\nStep %d: Type='%s', Params=%s\n", i+1, step.Type, string(paramsJSON))

		handler, exists := w.stepHandlers[step.Type]
		if exists {
			err := handler(step.Params)
			if err != nil {
				fmt.Printf("❌ Error executing step %d ('%s'): %v\n", i+1, step.Type, err)
				// Decide whether to stop or continue on error
				// return // Uncomment to stop on first error
			} else {
				fmt.Printf("Step %d executed successfully.\n", i+1)
			}
		} else {
			fmt.Printf("⚠️ Unknown step type: '%s'. Skipping.\n", step.Type)
		}
	}
	fmt.Println("\n--- Workflow Execution Complete ---")
}

// GetSteps returns a copy of the workflow steps (for testing/inspection).
func (w *DynamicWorkflow) GetSteps() []WorkflowStep {
    // Return a copy to prevent external modification
    stepsCopy := make([]WorkflowStep, len(w.steps))
    copy(stepsCopy, w.steps)
    return stepsCopy
}


// --- Builder Implementation ---

// DynamicWorkflowBuilder is the builder for creating DynamicWorkflow instances.
type DynamicWorkflowBuilder struct {
	steps []WorkflowStep
}

// NewDynamicWorkflowBuilder creates a new builder instance.
func NewDynamicWorkflowBuilder() *DynamicWorkflowBuilder {
	return &DynamicWorkflowBuilder{
		steps: []WorkflowStep{},
	}
}

// AddStep adds a new step to the workflow being built.
// It returns the builder pointer to allow method chaining.
func (b *DynamicWorkflowBuilder) AddStep(stepType string, params map[string]interface{}) *DynamicWorkflowBuilder {
	paramsJSON, _ := json.Marshal(params)
	fmt.Printf("Builder: Adding step '%s' with params %s\n", stepType, string(paramsJSON))
	b.steps = append(b.steps, WorkflowStep{Type: stepType, Params: params})
	return b // Return pointer for chaining
}

// Build constructs the final DynamicWorkflow product.
func (b *DynamicWorkflowBuilder) Build() *DynamicWorkflow {
	fmt.Println("Builder: Building the workflow...")
	if len(b.steps) == 0 {
		fmt.Println("Builder Warning: Building an empty workflow.")
	}
	// Pass a copy of the steps slice to the workflow constructor
	workflow := NewDynamicWorkflow(b.steps)
	fmt.Println("Builder: Workflow built.")
	return workflow
}
