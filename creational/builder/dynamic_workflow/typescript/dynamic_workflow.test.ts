import { DynamicWorkflowBuilder, DynamicWorkflow } from './dynamic_workflow';

// Mock console methods to capture output
const consoleLogSpy = jest.spyOn(console, 'log').mockImplementation();
const consoleWarnSpy = jest.spyOn(console, 'warn').mockImplementation();
const consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation();

describe('DynamicWorkflowBuilder', () => {

    beforeEach(() => {
        // Clear mocks before each test
        consoleLogSpy.mockClear();
        consoleWarnSpy.mockClear();
        consoleErrorSpy.mockClear();
    });

    test('should build an empty workflow', () => {
        const builder = new DynamicWorkflowBuilder();
        const workflow = builder.build();
        expect(workflow).toBeInstanceOf(DynamicWorkflow);
        expect(workflow.getSteps()).toHaveLength(0);
        expect(consoleWarnSpy).toHaveBeenCalledWith("Builder Warning: Building an empty workflow.");
    });

    test('should add a single step', () => {
        const builder = new DynamicWorkflowBuilder();
        const workflow = builder.addStep("send_email", { to: "test@example.com" }).build();
        const steps = workflow.getSteps();
        expect(steps).toHaveLength(1);
        expect(steps[0].type).toBe("send_email");
        expect(steps[0].params).toEqual({ to: "test@example.com" });
    });

    test('should add multiple steps with chaining', () => {
        const builder = new DynamicWorkflowBuilder();
        const workflow = builder
            .addStep("run_script", { path: "/script1" })
            .addStep("notify_slack", { channel: "#test" })
            .build();
        const steps = workflow.getSteps();
        expect(steps).toHaveLength(2);
        expect(steps[0].type).toBe("run_script");
        expect(steps[1].type).toBe("notify_slack");
        expect(steps[1].params).toEqual({ channel: "#test" });
    });

    test('addStep should return the builder instance (this)', () => {
        const builder = new DynamicWorkflowBuilder();
        const returnValue = builder.addStep("test_step");
        expect(returnValue).toBe(builder);
    });
});

describe('DynamicWorkflow Execution', () => {

     beforeEach(() => {
        // Clear mocks before each test
        consoleLogSpy.mockClear();
        consoleWarnSpy.mockClear();
        consoleErrorSpy.mockClear();
    });

    test('should execute an empty workflow', () => {
        const workflow = new DynamicWorkflow([]);
        workflow.execute();
        // Check if the specific message for empty workflow was logged
        expect(consoleLogSpy).toHaveBeenCalledWith("Workflow has no steps.");
    });

    test('should execute known steps successfully', () => {
        const steps = [
            { type: "send_email", params: { to: "exec@test.com", subject: "Exec Test" } },
            { type: "run_script", params: { path: "/exec/test.py" } },
            { type: "notify_slack", params: { channel: "#exec", message: "Done" } }
        ];
        const workflow = new DynamicWorkflow(steps);
        workflow.execute();

        // Check if log messages contain expected simulation output
        const logOutput = consoleLogSpy.mock.calls.flat().join('\n');
        expect(logOutput).toContain("Simulating sending email to 'exec@test.com'");
        expect(logOutput).toContain("Simulating running script: '/exec/test.py'");
        expect(logOutput).toContain("Simulating notifying Slack channel '#exec'");

        // Check that no warnings or errors were logged
        expect(consoleWarnSpy).not.toHaveBeenCalled();
        expect(consoleErrorSpy).not.toHaveBeenCalled();
    });

    test('should handle unknown step types with a warning', () => {
        const steps = [
            { type: "send_email", params: {} },
            { type: "fly_to_moon", params: { speed: "warp" } }, // Unknown step
            { type: "run_script", params: {} }
        ];
        const workflow = new DynamicWorkflow(steps);
        workflow.execute();

        // Check if the warning for the unknown step was logged
        expect(consoleWarnSpy).toHaveBeenCalledWith("⚠️ Unknown step type: 'fly_to_moon'. Skipping.");

        // Check that known steps still executed (based on log messages)
        const logOutput = consoleLogSpy.mock.calls.flat().join('\n');
        expect(logOutput).toContain("Simulating sending email");
        expect(logOutput).toContain("Simulating running script");

        // Check that no errors were logged
        expect(consoleErrorSpy).not.toHaveBeenCalled();
    });

     test('should use parameters passed to handlers', () => {
        const steps = [
            { type: "notify_slack", params: { channel: "#specific-channel", message: "Specific message!" } }
        ];
        const workflow = new DynamicWorkflow(steps);
        workflow.execute();

        const logOutput = consoleLogSpy.mock.calls.flat().join('\n');
        expect(logOutput).toContain("Simulating notifying Slack channel '#specific-channel': 'Specific message!...'" );
    });
});
