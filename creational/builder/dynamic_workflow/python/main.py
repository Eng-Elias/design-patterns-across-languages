from dynamic_workflow import DynamicWorkflowBuilder

def main():
    print("--- Builder Pattern - Dynamic Workflow Demo (Python) ---")

    # Create a workflow using the builder with method chaining
    builder = DynamicWorkflowBuilder()

    workflow = (
        builder
        .add_step("send_email", to="dev@example.com", subject="Build Started", body="Starting the main build process...")
        .add_step("run_script", path="/scripts/compile.sh")
        .add_step("run_script", path="/scripts/test.sh")
        .add_step("notify_slack", channel="#builds", message="Build and Tests successful!")
        .add_step("unknown_step", parameter="some_value") # Test unknown step
        .add_step("send_email", to="qa@example.com", subject="Build Ready for QA") # Missing body
        .build()
    )

    print("\n--- Workflow Constructed ---")

    # Execute the constructed workflow
    workflow.execute()

    print("\n--- Building and executing a second, simpler workflow --- ")
    workflow2 = (
        DynamicWorkflowBuilder() # Create a new builder instance
        .add_step("notify_slack", channel="#monitoring", message="System check initiated.")
        .build()
    )
    workflow2.execute()

    print("\n--- Demo Complete ---")

if __name__ == "__main__":
    main()
