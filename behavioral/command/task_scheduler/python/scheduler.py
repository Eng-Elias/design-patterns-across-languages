from typing import List, Deque
from collections import deque
from commands import Command # Relative import within the package

class TaskScheduler:
    """
    The Invoker class. It holds commands and triggers their execution.
    Doesn't know anything about concrete commands or receivers.
    """
    def __init__(self):
        # Using deque for potential efficiency if tasks were added/removed frequently from both ends
        # but a list would also work fine here.
        self._tasks: Deque[Command] = deque()

    def add_task(self, command: Command) -> None:
        """Adds a command (task) to the queue."""
        print(f"Adding task: {command.__class__.__name__}")
        self._tasks.append(command)

    def run_pending_tasks(self) -> None:
        """Executes all tasks currently in the queue."""
        print("\n--- Running Scheduled Tasks ---")
        if not self._tasks:
            print("No tasks to run.")
            return

        # Execute tasks in FIFO order
        while self._tasks:
            task = self._tasks.popleft()
            print(f"\nExecuting task: {task.__class__.__name__}...")
            try:
                task.execute()
                print(f"Task {task.__class__.__name__} completed.")
            except Exception as e:
                print(f"Error executing task {task.__class__.__name__}: {e}")
        print("--- All tasks processed ---")
