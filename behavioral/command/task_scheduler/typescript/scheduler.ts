import { Command } from "./commands";

export class TaskScheduler {
  private tasks: Command[] = [];

  public addTask(command: Command): void {
    // Attempt to get a description for logging, default if method doesn't exist
    const description = (command as any).getDescription
      ? (command as any).getDescription()
      : command.constructor.name;
    console.log(`Adding task: ${description}`);
    this.tasks.push(command);
  }

  public async runPendingTasks(): Promise<void> {
    console.log("\n--- Running Scheduled Tasks ---");
    if (this.tasks.length === 0) {
      console.log("No tasks to run.");
      return;
    }

    // Execute tasks sequentially in FIFO order
    // Could also run in parallel with Promise.all if tasks are independent
    for (const task of this.tasks) {
      const description = (task as any).getDescription
        ? (task as any).getDescription()
        : task.constructor.name;
      console.log(`\nExecuting task: ${description}...`);
      try {
        await task.execute(); // Await completion if execute is async
        console.log(`Task ${description} completed.`);
      } catch (error) {
        console.error(`Error executing task ${description}:`, error);
      }
    }

    // Clear the tasks queue after execution
    this.tasks = [];
    console.log("--- All tasks processed ---");
  }

  // Helper for testing or inspection
  public getTaskCount(): number {
    return this.tasks.length;
  }
}
