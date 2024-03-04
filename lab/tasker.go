package main

import (
	"fmt"
)

// Define constants for task statuses
const (
	IncompleteStatus = "Incomplete"
	CompleteStatus   = "Complete"
)

// Define a struct for the task
type Task struct {
	ID          int
	Description string
	Status      string
}

func main() {
	// Initializing a map to store tasks using their IDs as keys
	taskMap := make(map[int]Task)
	// Initializing tasks in the map
	taskMap[1] = Task{1, "Brainstorm Go domain", IncompleteStatus}
	taskMap[2] = Task{2, "Google SPD model", IncompleteStatus}
	taskMap[3] = Task{3, "Prepare SPD paper LaTex", IncompleteStatus}
	// Simulate marking a task as complete
	completeTask(1, taskMap)
	// Simulate adding a new task

	newTask := Task{4, "Eat the grapes before they rot", IncompleteStatus}
	addTask(newTask, taskMap)
	// Display the updated task list
	displayTaskList(taskMap)
}

// Function to mark a task as complete
func completeTask(taskID int, tasks map[int]Task) {
	if task, ok := tasks[taskID]; ok {
		tasks[taskID] = Task{task.ID, task.Description, CompleteStatus}
		fmt.Printf("Task marked as complete! '%s'\n", task.Description)
	} else {
		fmt.Println("Task not found in the list.")
	}
}

// Function to add a new task to the list
func addTask(newTask Task, tasks map[int]Task) {
	if _, ok := tasks[newTask.ID]; !ok { // Ignore the task value with the blank

		tasks[newTask.ID] = newTask
		fmt.Printf("New task added! '%s' with ID %d added to the list.\n",

			newTask.Description, newTask.ID)
	} else {
		fmt.Println("Task with the same ID already exists.")
		fmt.Println("Please choose a different ID.")
	}
}

// Function to display the current task list
func displayTaskList(tasks map[int]Task) {
	fmt.Println("\nCurrent Task List:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: '%s', Status: %s\n", task.ID,

			task.Description, task.Status)
	}
}
