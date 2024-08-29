package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

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

func generateIDGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func (tl *TaskList) ToJSON() ([]byte, error) {
	taskList := tl.GetTaskList()
	payload, err := json.Marshal(taskList)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (tl *TaskList) FromJSON(data []byte) error {
	taskList := make(map[int]Task)
	err := json.Unmarshal(data, &taskList)
	if err != nil {
		return err
	}
	tl.mutex.Lock()
	defer tl.mutex.Unlock()
	tl.tasks = taskList
	return nil
}

func (tl *TaskList) CompleteTask(taskID int) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	task, ok := tl.tasks[taskID]
	if !ok || task == (Task{}) {
		return ErrTaskNotFound
	}
	task.Status = "Completed"
	tl.tasks[taskID] = task
	fmt.Printf("Task marked as complete! '%s'\n", task.Description)
	return nil
}

func (tl *TaskList) AddTask(newTask Task) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	newTask.ID = tl.idGen()
	if _, ok := tl.tasks[newTask.ID]; ok {
		return ErrTaskAlreadyExists
	}
	// Send request to add task concurrently using a channel
	taskRequestChan <- taskRequest{operation: "add", task: newTask}
	return nil
}

func (tl *TaskList) EditTask(taskID int, editedTask Task) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	task, ok := tl.tasks[taskID]
	if !ok || task == (Task{}) {
		return ErrTaskNotFound
	}
	editedTask.ID = taskID
	tl.tasks[taskID] = editedTask
	// Send request to edit task concurrently using a channel
	taskRequestChan <- taskRequest{operation: "edit", taskID: taskID, task: editedTask}
	return nil
}

func (tl *TaskList) DeleteTask(taskID int) error {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()

	_, ok := tl.tasks[taskID]
	if !ok {
		return ErrTaskNotFound
	}
	delete(tl.tasks, taskID)
	// Send request to delete task concurrently using a channel
	taskRequestChan <- taskRequest{operation: "delete", taskID: taskID}
	return nil
}

func (tl *TaskList) GetTaskList() map[int]Task {
	tl.mutex.Lock()
	defer tl.mutex.Unlock()
	taskListCopy := make(map[int]Task)
	for k, v := range tl.tasks {
		taskListCopy[k] = v
	}
	return taskListCopy
}

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
		case "delete":
			err := taskManager.DeleteTask(req.taskID)
			if err != nil {
				response = fmt.Sprintf("Error: %v", err)
			} else {
				response = fmt.Sprintf("Task with ID %d deleted.", req.taskID)
			}
		}
		req.response <- response
	}
}

func displayTaskList(taskList map[int]Task) {
	if len(taskList) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\nTask List:")
	fmt.Println("ID\tDescription\tStatus")
	fmt.Println("---------\t-----------\t-------")
	for _, task := range taskList {
		fmt.Printf("%d\t%s\t%s\n", task.ID, task.Description, task.Status)
	}
}

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

	const filePath = "tasks.json" // Define file path for saving/loading tasks

	for {
		fmt.Println("\nTask Management Menu:")
		fmt.Println("1. Mark Task as Complete")
		fmt.Println("2. Add New Task")
		fmt.Println("3. Edit Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. View Task List")
		fmt.Println("6. View As JSON")
		fmt.Println("7. View As Type")
		fmt.Println("8. Save Tasks to File")
		fmt.Println("9. Load Tasks from File")
		fmt.Println("10. Exit")
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
			// Send request to mark task complete concurrently
			taskRequestChan <- taskRequest{operation: "complete", taskID: taskID}
		case 2:
			var newTask Task
			fmt.Print("Enter Task Description: ")
			if !reader.Scan() {
				fmt.Println("Error reading input:", reader.Err())
				continue
			}
			newTask.Description = strings.TrimSpace(reader.Text())
			newTask.Status = "Pending" // Set default status to Pending
			// Send request to add task concurrently
			taskRequestChan <- taskRequest{operation: "add", task: newTask}
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
			editedTask.Status = taskManager.tasks[taskID].Status // Maintain existing status

			// Send request to edit task concurrently
			taskRequestChan <- taskRequest{operation: "edit", taskID: taskID, task: editedTask}
		case 4:
			fmt.Print("Enter Task ID to delete: ")
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
			taskRequestChan <- taskRequest{operation: "delete", taskID: taskID}
		case 5:
			displayTaskList(taskManager.GetTaskList())
		case 6:
			for _, task := range taskManager.GetTaskList() {
				fmt.Println(task.ToJSON())
			}
		case 7:
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
		case 8:
			payload, err := taskManager.ToJSON()
			if err != nil {
				fmt.Println("Error marshalling tasks to JSON:", err)
				continue
			}
			err = os.WriteFile(filePath, payload, 0644) // Write tasks to file with read/write permissions
			if err != nil {
				fmt.Println("Error saving tasks to file:", err)
				continue
			}
			fmt.Println("Tasks saved successfully to", filePath)
		case 9:
			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading tasks from file:", err)
				continue
			}
			err = taskManager.FromJSON(data)
			if err != nil {
				fmt.Println("Error unmarshalling tasks from JSON:", err)
				continue
			}
			fmt.Println("Tasks loaded successfully from", filePath)
		case 10:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid menu selection. Please try again.")
		}
	}
}
