// Task manager app that persist tasks in a json file:
// 1. Add task
// 2. Get task by id
// 3. Add bulk tasks
// 4. Update a task
// 5. Delete a task
// 6. Mark task as complete
// 7. Get all tasks

package tasker

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"is_complete"`
}

var tasks = []Task{}
var nextId = 1

func Tasker() {
	fmt.Println("Welcome to Tasker!!")

	// todo
	firstTodo := Task{
		ID:          1,
		Title:       "Building a website",
		Description: "Building an e-commerce website using HTML and CSS",
		IsComplete:  false,
	}
	secondTodo := Task{
		ID:          2,
		Title:       "Chopping the money",
		Description: "A go chop the money today",
		IsComplete:  true,
	}

	thirdTodo := Task{
		ID:          3,
		Title:       "Collecting debts",
		Description: "Collecting my business debts today for review",
		IsComplete:  true,
	}
	addTask(firstTodo)
	addTask(secondTodo)
	addTask(thirdTodo)
	fmt.Println("Your tasks:", getAllTasks())
	secondTask, _ := getTask(2)
	fmt.Println("Second task:", secondTask)
	// deleteTask(3)
	// markComplete(1)
}

// complete method
func (task *Task) Complete() {
	task.IsComplete = true
}

// reopen method
func (task *Task) Reopen() {
	task.IsComplete = false
}

// check if is completed
func (task Task) IsCompleted() bool {
	return task.IsComplete
}

// add task
func addTask(task Task) {
	task.ID = nextId
	nextId++
	tasks = append(tasks, task)
}

// get task by ID
func getTask(id int) (Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("todo with id %d not found", id)
}

// get all tasks
func getAllTasks() []Task {
	return tasks
}

// update a task
func updateTask(id int, title, description string) error {
	for i := range tasks { // we use i instead of for _, task := range tasks because we want to modify what is at that index
		if tasks[i].ID == id {
			tasks[i].Title = title
			tasks[i].Description = description
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}

// delete task
func deleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}

// mark task as complete
func markComplete(id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].IsComplete = true
			return nil
		}
	}
	return fmt.Errorf("task with id %v not found", id)
}

// working with JSON
func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

// load tasks
func loadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}
	return nil
}
