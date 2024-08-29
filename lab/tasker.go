package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Define constants for task statuses
const (
	IncompleteStatus = "Incomplete"
	CompleteStatus   = "Complete"
)

// Define custom errors
var (
	ErrTaskNotFound      = errors.New("Task not found")
	ErrTaskAlreadyExists = errors.New("Task with the same ID already exists")
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
	mutex sync.Mutex
}

type taskRequest struct {
	operation string
	taskID    int
	task      Task
	response  chan string
}

var taskRequestChan = make(chan taskRequest)

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
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	task, ok := tl.tasks[taskID]
	if !ok || task == (Task{}) {
		return ErrTaskNotFound
	}
	task.Status = CompleteStatus
	tl.tasks[taskID] = task
	fmt.Printf("Task marked as complete! '%s'\n", task.Description)
	return nil
}

// AddTask adds a new task to the list
func (tl *TaskList) AddTask(newTask Task) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

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
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	task, ok := tl.tasks[taskID]
	fmt.Println(ok, task)
	if !ok || task == (Task{}) {
		return ErrTaskNotFound
	}
	editedTask.ID = taskID
	tl.tasks[taskID] = editedTask
	fmt.Println("Task was:\n", task)
	fmt.Printf("Task updated! '%s' with ID %d was edited.\n", editedTask.Description, editedTask.ID)
	return nil
}

// GetTaskList returns the current task list
// func (tl TaskList) GetTaskList() map[int]Task {
func (tl *TaskList) GetTaskList() map[int]Task {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	taskListCopy := make(map[int]Task)
	for k, v := range tl.tasks {
		taskListCopy[k] = v
	}
	return taskListCopy
	// return tl.tasks
}

// Function to display the current task list
func displayTaskList(tasks map[int]Task) {
	fmt.Println("\nCurrent Task List:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: '%s', Status: %s\n", task.ID, task.Description, task.Status)
	}
}

// Goroutine to handle task management requests
func taskManagerLoop(taskManager *TaskList) {
	for req := range taskRequestChan {
		var response string
		switch req.operation {
		case "complete":
			err := taskManager.CompleteTask(req.taskID)
			if err != nil {
				response = fmt.Sprintf("Error: %v", err)
			} else {
				task, _ := taskManager.GetTaskList()[req.taskID]
				response = fmt.Sprintf("Task marked as complete! '%s'", task.Description)
			}
		case "add":
			err := taskManager.AddTask(req.task)
			if err != nil {
				response = fmt.Sprintf("Error: %v", err)
			} else {
				response = fmt.Sprintf("New task added! '%s' with ID %d added to the list.", req.task.Description, req.task.ID)
			}
		case "edit":
			err := taskManager.EditTask(req.taskID, req.task)
			if err != nil {
				response = fmt.Sprintf("Error: %v", err)
			} else {
				response = fmt.Sprintf("Task updated! '%s' with ID %d was edited.", req.task.Description, req.task.ID)
			}
		}
		req.response <- response
	}
}

// Goroutine to handle console output
func consoleOutputLoop() {
	for {
		req := <-taskRequestChan
		fmt.Println(req.response)
	}
}

func main() {
	taskManager := NewTaskList()
	reader := bufio.NewScanner(os.Stdin)

	go taskManagerLoop(taskManager)
	go consoleOutputLoop()

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
			responseChan := make(chan string)
			taskRequestChan <- taskRequest{operation: "complete", taskID: taskID, response: responseChan}
			<-responseChan
			// if err := taskManager.CompleteTask(taskID); err != nil {
			// 	fmt.Println("Error:", err)
			// }
		case 2:
			var newTask Task
			fmt.Print("Enter Task Description: ")
			if !reader.Scan() {
				fmt.Println("Error reading input:", reader.Err())
				continue
			}
			newTask.Description = strings.TrimSpace(reader.Text())
			newTask.Status = IncompleteStatus
			responseChan := make(chan string)
			taskRequestChan <- taskRequest{operation: "add", task: newTask, response: responseChan}
			<-responseChan
			// if err := taskManager.AddTask(newTask); err != nil {
			// 	fmt.Println("Error:", err)
			// }
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

			responseChan := make(chan string)
			taskRequestChan <- taskRequest{operation: "edit", taskID: taskID, task: editedTask, response: responseChan}
			<-responseChan
			// if err := taskManager.EditTask(taskID, editedTask); err != nil {
			// 	fmt.Println("Error:", err)
			// }
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
