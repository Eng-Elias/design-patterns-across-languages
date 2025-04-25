// Define the structure for a single workflow step
interface WorkflowStep {
  type: string;
  params: Record<string, any>; // Using Record for key-value pairs
}

// Define the handler function signature
type StepHandler = (params: Record<string, any>) => void;

// Represents the complex object (Product) being built
export class DynamicWorkflow {
  private readonly steps: WorkflowStep[];
  // Map step types to their handler methods
  private readonly stepHandlers: Map<string, StepHandler> = new Map();

  constructor(steps: WorkflowStep[]) {
    // Create a defensive copy of the steps
    this.steps = [...steps];

    // Initialize handlers
    this.stepHandlers.set("send_email", this._handleSendEmail.bind(this));
    this.stepHandlers.set("run_script", this._handleRunScript.bind(this));
    this.stepHandlers.set("notify_slack", this._handleNotifySlack.bind(this));
    // Add more handlers here
  }

  private _handleSendEmail(params: Record<string, any>): void {
    const to = params.to || "<default_email>";
    const subject = params.subject || "<default_subject>";
    const body = params.body || "";
    console.log(
      `📧 Simulating sending email to '${to}' with subject '${subject}': '${body.substring(
        0,
        30
      )}...'`
    );
  }

  private _handleRunScript(params: Record<string, any>): void {
    const path = params.path || "<default_script_path>";
    console.log(`⚙️ Simulating running script: '${path}'`);
  }

  private _handleNotifySlack(params: Record<string, any>): void {
    const channel = params.channel || "#general";
    const message = params.message || "";
    console.log(
      `💬 Simulating notifying Slack channel '${channel}': '${message.substring(
        0,
        30
      )}...'`
    );
  }

  public execute(): void {
    console.log("--- Executing Workflow ---");
    if (this.steps.length === 0) {
      console.log("Workflow has no steps.");
      return;
    }

    this.steps.forEach((step, index) => {
      const stepType = step.type;
      const params = step.params || {};
      console.log(
        `\nStep ${index + 1}: Type='${stepType}', Params=${JSON.stringify(
          params
        )}`
      );

      const handler = this.stepHandlers.get(stepType);
      if (handler) {
        try {
          handler(params);
          console.log(`Step ${index + 1} executed successfully.`);
        } catch (e) {
          const error = e instanceof Error ? e.message : String(e);
          console.error(
            `❌ Error executing step ${index + 1} ('${stepType}'): ${error}`
          );
          // Decide whether to stop or continue on error
          // throw new Error(`Workflow stopped due to error in step ${index + 1}`); // Uncomment to stop
        }
      } else {
        console.warn(`⚠️ Unknown step type: '${stepType}'. Skipping.`);
      }
    });
    console.log("\n--- Workflow Execution Complete ---");
  }

  // Getter for tests or inspection
  public getSteps(): WorkflowStep[] {
    return [...this.steps]; // Return a copy
  }
}

// The Builder class for constructing DynamicWorkflow objects
export class DynamicWorkflowBuilder {
  private steps: WorkflowStep[] = [];

  public addStep(stepType: string, params: Record<string, any> = {}): this {
    console.log(
      `Builder: Adding step '${stepType}' with params ${JSON.stringify(params)}`
    );
    this.steps.push({ type: stepType, params });
    return this; // Return this for chaining
  }

  public build(): DynamicWorkflow {
    console.log("Builder: Building the workflow...");
    if (this.steps.length === 0) {
      console.warn("Builder Warning: Building an empty workflow.");
    }
    // Pass a copy of the steps to the workflow constructor
    const workflow = new DynamicWorkflow([...this.steps]);
    console.log("Builder: Workflow built.");
    return workflow;
  }
}
