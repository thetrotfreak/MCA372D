package main

import (
	"bufio"
	"encoding/json"
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
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// TaskList implements TaskManager interface
type TaskList struct {
	tasks map[int]Task
	idGen func() int // Closure to generate unique IDs
}

// TaskManager interface defines methods for managing tasks
type TaskManager interface {
	CompleteTask(taskID int) error
	AddTask(newTask Task) error
	EditTask(taskID int) error
	GetTaskList() map[int]Task
	ToJSON() (string, error)
	FromJSON() error
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

func (t Task) ToJSON() (string, error) {
	payload, err := json.Marshal(t)

	if err != nil {
		return "", err
	}

	return string(payload), err
}

func (t *Task) FromJSON(payload string) error {
	fmt.Println("In FromJSON(), got", *t)
	err := json.Unmarshal([]byte(payload), t)
	return err
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

// EditTask edits a given task
func (tl *TaskList) EditTask(taskID int, editedTask Task) error {
	task, ok := tl.tasks[taskID]
	if !ok {
		return ErrTaskNotFound
	}
	editedTask.ID = taskID
	tl.tasks[taskID] = editedTask
	fmt.Println("Task was:\n", task)
	fmt.Printf("Task updated! '%s' with ID %d was edited.\n", editedTask.Description, editedTask.ID)
	return nil
}

// GetTaskList returns the current task list
func (tl TaskList) GetTaskList() map[int]Task {
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
		fmt.Println("3. Edit Task")
		fmt.Println("4. View Task List")
		fmt.Println("5. View As JSON")
		fmt.Println("6. View As Type")
		fmt.Println("7. Exit")
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
			fmt.Print("Enter Task ID to edit: ")
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

			var editedTask Task
			fmt.Print("Enter Updated Task Description: ")
			if !reader.Scan() {
				fmt.Println("Error reading input:", reader.Err())
				continue
			}

			editedTask.Description = strings.TrimSpace(reader.Text())
			editedTask.Status = taskManager.tasks[taskID].Status

			if err := taskManager.EditTask(taskID, editedTask); err != nil {
				fmt.Println("Error:", err)
			}
		case 4:
			displayTaskList(taskManager.GetTaskList())
		case 5:
			for _, task := range taskManager.GetTaskList() {
				fmt.Println(task.ToJSON())
			}
		case 6:
			var t Task

			for _, task := range taskManager.GetTaskList() {
				payload, err := task.ToJSON()

				if err == nil {
					err = (&t).FromJSON(payload)
					if err == nil {
						fmt.Println(t)
					}
				}
			}
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid menu selection. Please try again.")
		}
	}
}
