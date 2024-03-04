package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define constants for task statuses
const (
	IncompleteStatus = "Incomplete"
	CompleteStatus   = "Complete"
)

// Define custom errors
var (
	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskAlreadyExists = errors.New("task with the same ID already exists")
)

// Define a struct for the task
type Task struct {
	ID          int
	Description string
	Status      string
}

// TaskManager interface defines methods for managing tasks
type TaskManager interface {
	CompleteTask(taskID int) error
	AddTask(newTask Task) error
	GetTaskList() map[int]Task
}

// TaskList implements TaskManager interface
type TaskList struct {
	tasks map[int]Task
	idGen func() int // Closure to generate unique IDs
}

// NewTaskList creates a new instance of TaskList with a unique ID generator
func NewTaskList() *TaskList {
	return &TaskList{
		tasks: make(map[int]Task),
		idGen: generateIDGenerator(),
	}
}

// generateIDGenerator returns a closure that generates auto-incrementing IDs
func generateIDGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// CompleteTask marks a task as complete
func (tl *TaskList) CompleteTask(taskID int) error {
	task, ok := tl.tasks[taskID]
	if !ok {
		return ErrTaskNotFound
	}
	task.Status = CompleteStatus
	tl.tasks[taskID] = task
	fmt.Printf("Task marked as complete! '%s'\n", task.Description)
	return nil
}

// AddTask adds a new task to the list
func (tl *TaskList) AddTask(newTask Task) error {
	newTask.ID = tl.idGen()
	if _, ok := tl.tasks[newTask.ID]; ok {
		return ErrTaskAlreadyExists
	}
	tl.tasks[newTask.ID] = newTask
	fmt.Printf("New task added! '%s' with ID %d added to the list.\n", newTask.Description, newTask.ID)
	return nil
}

// GetTaskList returns the current task list
func (tl *TaskList) GetTaskList() map[int]Task {
	return tl.tasks
}

// Function to display the current task list
func displayTaskList(tasks map[int]Task) {
	fmt.Println("\nCurrent Task List:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: '%s', Status: %s\n", task.ID, task.Description, task.Status)
	}
}

func main() {
	taskManager := NewTaskList()
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTask Management Menu:")
		fmt.Println("1. Mark Task as Complete")
		fmt.Println("2. Add New Task")
		fmt.Println("3. View Task List")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		if !reader.Scan() {
			fmt.Println("Error reading input:", reader.Err())
			continue
		}
		choiceStr := reader.Text()
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter Task ID to mark as complete: ")
			if !reader.Scan() {
				fmt.Println("Error reading input:", reader.Err())
				continue
			}
			taskIDStr := reader.Text()
			taskID, err := strconv.Atoi(taskIDStr)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if err := taskManager.CompleteTask(taskID); err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			var newTask Task
			fmt.Print("Enter Task Description: ")
			if !reader.Scan() {
				fmt.Println("Error reading input:", reader.Err())
				continue
			}
			newTask.Description = strings.TrimSpace(reader.Text())
			newTask.Status = IncompleteStatus
			if err := taskManager.AddTask(newTask); err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			displayTaskList(taskManager.GetTaskList())
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid menu selection. Please try again.")
		}
	}
}
