import unittest
from unittest.mock import patch
import io # To capture print output

from dynamic_workflow import DynamicWorkflowBuilder, DynamicWorkflow

class TestDynamicWorkflowBuilder(unittest.TestCase):

    def test_build_empty_workflow(self):
        builder = DynamicWorkflowBuilder()
        workflow = builder.build()
        self.assertIsInstance(workflow, DynamicWorkflow)
        self.assertEqual(len(workflow.steps), 0)

    def test_add_single_step(self):
        builder = DynamicWorkflowBuilder()
        workflow = builder.add_step("send_email", to="test@example.com").build()
        self.assertEqual(len(workflow.steps), 1)
        self.assertEqual(workflow.steps[0]["type"], "send_email")
        self.assertEqual(workflow.steps[0]["params"], {"to": "test@example.com"})

    def test_add_multiple_steps_chaining(self):
        builder = DynamicWorkflowBuilder()
        workflow = (
            builder
            .add_step("run_script", path="/script1")
            .add_step("notify_slack", channel="#test")
            .build()
        )
        self.assertEqual(len(workflow.steps), 2)
        self.assertEqual(workflow.steps[0]["type"], "run_script")
        self.assertEqual(workflow.steps[1]["type"], "notify_slack")
        self.assertEqual(workflow.steps[1]["params"], {"channel": "#test"})

    def test_builder_returns_self(self):
        builder = DynamicWorkflowBuilder()
        ret = builder.add_step("test")
        self.assertIs(ret, builder) # Check if add_step returns the builder itself

    def test_build_returns_workflow_instance(self):
        builder = DynamicWorkflowBuilder()
        workflow = builder.build()
        self.assertIsInstance(workflow, DynamicWorkflow)

class TestDynamicWorkflowExecution(unittest.TestCase):

    def test_execute_empty_workflow(self):
        workflow = DynamicWorkflow([])
        with patch('sys.stdout', new_callable=io.StringIO) as mock_stdout:
            workflow.execute()
            output = mock_stdout.getvalue()
            self.assertIn("Workflow has no steps", output)

    def test_execute_known_steps(self):
        steps = [
            {"type": "send_email", "params": {"to": "exec@test.com", "subject": "Exec Test"}},
            {"type": "run_script", "params": {"path": "/exec/test.py"}},
            {"type": "notify_slack", "params": {"channel": "#exec", "message": "Done"}}
        ]
        workflow = DynamicWorkflow(steps)
        with patch('sys.stdout', new_callable=io.StringIO) as mock_stdout:
            workflow.execute()
            output = mock_stdout.getvalue()
            self.assertIn("Simulating sending email to 'exec@test.com'", output)
            self.assertIn("Simulating running script: '/exec/test.py'", output)
            self.assertIn("Simulating notifying Slack channel '#exec'", output)
            self.assertNotIn("Unknown step type", output)
            self.assertNotIn("Error executing step", output)

    def test_execute_unknown_step(self):
        steps = [
            {"type": "send_email", "params": {}},
            {"type": "do_magic", "params": {"power": 11}},
            {"type": "run_script", "params": {}}
        ]
        workflow = DynamicWorkflow(steps)
        with patch('sys.stdout', new_callable=io.StringIO) as mock_stdout:
            workflow.execute()
            output = mock_stdout.getvalue()
            self.assertIn("Simulating sending email", output)
            self.assertIn("⚠️ Unknown step type: 'do_magic'. Skipping.", output)
            self.assertIn("Simulating running script", output)

    def test_handler_parameters_used(self):
        # Focus on one handler to check parameter passing
        steps = [
            {"type": "send_email", "params": {"to": "param_user@test.org", "subject": "Param Subject"}}
        ]
        workflow = DynamicWorkflow(steps)
        with patch('sys.stdout', new_callable=io.StringIO) as mock_stdout:
            workflow.execute()
            output = mock_stdout.getvalue()
            self.assertIn("email to 'param_user@test.org'", output)
            self.assertIn("subject 'Param Subject'", output)

if __name__ == '__main__':
    unittest.main()
