from typing import List, Dict, Any, Callable


class DynamicWorkflow:
    """Represents the complex object (Product) being built."""

    def __init__(self, steps: List[Dict[str, Any]]):
        self.steps = steps
        # Map step types to their handler methods
        self.step_handlers: Dict[str, Callable[[Dict[str, Any]], None]] = {
            "send_email": self._handle_send_email,
            "run_script": self._handle_run_script,
            "notify_slack": self._handle_notify_slack,
            # Add more handlers here as needed
        }

    def _handle_send_email(self, params: Dict[str, Any]):
        to = params.get("to", "<default_email>")
        subject = params.get("subject", "<default_subject>")
        body = params.get("body", "")
        print(f"📧 Simulating sending email to '{to}' with subject '{subject}': '{body[:30]}...'" )

    def _handle_run_script(self, params: Dict[str, Any]):
        path = params.get("path", "<default_script_path>")
        print(f"⚙️ Simulating running script: '{path}'")

    def _handle_notify_slack(self, params: Dict[str, Any]):
        channel = params.get("channel", "#general")
        message = params.get("message", "")
        print(f"💬 Simulating notifying Slack channel '{channel}': '{message[:30]}...'" )

    def execute(self):
        print("--- Executing Workflow ---")
        if not self.steps:
            print("Workflow has no steps.")
            return

        for i, step in enumerate(self.steps):
            step_type = step.get("type")
            params = step.get("params", {})
            print(f"\nStep {i+1}: Type='{step_type}', Params={params}")

            handler = self.step_handlers.get(step_type)
            if handler:
                try:
                    handler(params)
                    print(f"Step {i+1} executed successfully.")
                except Exception as e:
                    print(f"❌ Error executing step {i+1} ('{step_type}'): {e}")
                    # Decide whether to stop or continue on error
                    # break # Uncomment to stop on first error
            else:
                print(f"⚠️ Unknown step type: '{step_type}'. Skipping.")
        print("\n--- Workflow Execution Complete ---")


class DynamicWorkflowBuilder:
    """The Builder class for constructing DynamicWorkflow objects."""

    def __init__(self):
        self._steps: List[Dict[str, Any]] = []

    def add_step(self, step_type: str, **params: Any):
        """Adds a step to the workflow being built."""
        print(f"Builder: Adding step '{step_type}' with params {params}")
        self._steps.append({"type": step_type, "params": params})
        return self  # Return self to allow method chaining

    def build(self) -> DynamicWorkflow:
        """Constructs and returns the final DynamicWorkflow object."""
        print("Builder: Building the workflow...")
        if not self._steps:
            print("Builder Warning: Building an empty workflow.")
        workflow = DynamicWorkflow(list(self._steps)) # Pass a copy
        print("Builder: Workflow built.")
        return workflow
