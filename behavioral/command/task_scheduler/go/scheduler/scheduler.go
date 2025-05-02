package scheduler

import (
	"fmt"
	// Use the correct import path based on your go.mod file for the command package
	"log"

	"command_pattern_task_scheduler_go/command"
)

// TaskScheduler is the Invoker.
type TaskScheduler struct {
	tasks []command.Command // Slice to hold tasks (Commands)
}

// NewTaskScheduler creates a new scheduler instance.
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]command.Command, 0),
	}
}

// AddTask adds a command to the scheduler's queue.
func (s *TaskScheduler) AddTask(cmd command.Command) {
	// Use fmt.Stringer interface for logging if command implements it
    cmdStr := "Unknown Command"
    if stringer, ok := cmd.(fmt.Stringer); ok {
        cmdStr = stringer.String()
    } else {
         // Fallback using type reflection (less ideal)
         cmdStr = fmt.Sprintf("%T", cmd)
    }

	log.Printf("Adding task: %s\n", cmdStr)
	s.tasks = append(s.tasks, cmd)
}

// RunPendingTasks executes all commands in the queue sequentially.
func (s *TaskScheduler) RunPendingTasks() {
	fmt.Println("\n--- Running Scheduled Tasks ---")
	if len(s.tasks) == 0 {
		fmt.Println("No tasks to run.")
		return
	}

	// Keep track of tasks to remove (or create a new slice)
	executedTasks := 0
	for i, task := range s.tasks {
        cmdStr := "Unknown Command"
        if stringer, ok := task.(fmt.Stringer); ok {
            cmdStr = stringer.String()
        } else {
             cmdStr = fmt.Sprintf("%T", task)
        }

		fmt.Printf("\nExecuting task [%d]: %s...\n", i+1, cmdStr)
		err := task.Execute()
		if err != nil {
			// Log error but continue with other tasks
			log.Printf("Error executing task %s: %v\n", cmdStr, err)
		} else {
			fmt.Printf("Task %s completed.\n", cmdStr)
		}
		executedTasks++
	}

	// Clear the executed tasks (create a new empty slice)
	s.tasks = make([]command.Command, 0)
	fmt.Printf("--- %d tasks processed ---\n", executedTasks)
}

// GetTaskCount returns the number of pending tasks (for testing/inspection).
func (s *TaskScheduler) GetTaskCount() int {
    return len(s.tasks)
}
