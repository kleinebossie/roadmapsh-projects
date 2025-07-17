package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var tasks []Task

const dataFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	loadTasks()

	switch os.Args[1] {
	case "help":
		printHelp()
	case "add":
		cmdAdd()
	case "update":
		cmdUpdate()
	case "delete":
		cmdDelete()
	case "mark-todo":
		cmdMarkTodo()
	case "mark-in-progress":
		cmdMarkInProgress()
	case "mark-done":
		cmdMarkDone()
	case "list":
		cmdList()
	}

	saveTasks()
}

func printHelp() {
	fmt.Println(`Task CLI Usage:
  add <description>           Add a new task
  update <id> <description>   Update a task
  delete <id>                 Delete a task
  mark-todo <id>              Mark task as todo
  mark-in-progress <id>       Mark task as in progress
  mark-done <id>              Mark task as done
  list [filter]               List tasks (optional filter: todo, in-progress, done)`)
}

func loadTasks() {
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		tasks = []Task{}
		return
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatalf("Failed to read json file: %v\n", err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		if len(data) == 0 {
			tasks = []Task{}
			return
		}
		log.Fatalf("Failed to parse json file: %v\n", err)
	}
}

func cmdAdd() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: task-cli add <description>")
	}

	description := strings.Join(os.Args[2:], " ")
	if description == "" {
		log.Fatalln("Usage: task-cli add <description>")
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	now := time.Now()

	newTask := Task{
		id,
		description,
		"todo",
		now,
		now,
	}

	tasks = append(tasks, newTask)
	fmt.Printf("Task added succesfully (ID: %v)\n", newTask.ID)
}

func cmdUpdate() {
	if len(os.Args) < 4 {
		log.Fatalln("Usage: task-cli update <id> <description>")
	}

	idStr := os.Args[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Failed to convert task ID to integer: %v\n", err)
	}

	newDescription := strings.Join(os.Args[3:], " ")

	for i := range tasks {
		if tasks[i].ID != id {
			continue
		}

		tasks[i].Description = newDescription
		fmt.Printf("Task with ID: %v updated succesfully\n", id)
		return
	}

	fmt.Printf("Failed to update task with ID: %v\n", id)
}

func cmdDelete() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: task-cli delete <id>")
	}

	idStr := os.Args[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Failed to convert task ID to integer: %v\n", err)
	}

	for i := range tasks {
		if tasks[i].ID != id {
			continue
		}

		var confirm string
		fmt.Printf("Are you sure you want to delete the task with ID: %v? (y/n) ", id)
		_, err = fmt.Scanln(&confirm)
		if err != nil {
			log.Fatalf("Failed to register answer: %v", err)
		}

		confirm = strings.ToLower(confirm)

		if confirm != "y" && confirm != "yes" {
			continue
		}

		tasks = append(tasks[:i], tasks[i+1:]...)

		fmt.Printf("Task with ID: %v deleted succesfully\n", id)
		return
	}

	fmt.Printf("Failed to delete task with ID: %v\n", id)
}

func cmdMarkTodo() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: task-cli mark-todo <id>")
	}

	idStr := os.Args[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Failed to convert task ID to integer: %v\n", err)
	}

	for i := range tasks {
		if tasks[i].ID != id {
			continue
		}

		tasks[i].Status = "todo"
		fmt.Printf("Task with ID: %v succesfully marked as todo\n", id)
		return
	}

	fmt.Printf("Failed to mark task with ID: %v as todo\n", id)
}


func cmdMarkInProgress() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: task-cli mark-in-progress <id>")
	}

	idStr := os.Args[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Failed to convert task ID to integer: %v\n", err)
	}

	for i := range tasks {
		if tasks[i].ID != id {
			continue
		}

		tasks[i].Status = "in-progress"
		fmt.Printf("Task with ID: %v succesfully marked as in progress\n", id)
		return
	}

	fmt.Printf("Failed to mark task with ID: %v as in progress\n", id)
}

func cmdMarkDone() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: task-cli mark-done <id>")
	}

	idStr := os.Args[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Failed to convert task ID to integer: %v\n", err)
	}

	for i := range tasks {
		if tasks[i].ID != id {
			continue
		}

		tasks[i].Status = "done"
		fmt.Printf("Task with ID: %v succesfully marked as done\n", id)
		return
	}

	fmt.Printf("Failed to mark task with ID: %v as done\n", id)
}

func cmdList() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli list [optional: filter (todo/in-progress/done)]")
	}

	var filter string
	if len(os.Args) == 2 {
		filter = ""
	} else {
		filter = os.Args[2]
	}
	printTasks(filter)
}


func printTasks(filter string) {
	if len(tasks) == 0 {
		return
	}
	fmt.Println("ALL TASKS")
	for i, v := range tasks {
		if v.Status == filter || filter == "" {
			fmt.Printf(`Task #%v:
  ID:          %v
  Description: %v
  Status:      %v
  Created at:  %v
  Updated at:  %v
%v`, i + 1, v.ID, v.Description, v.Status, v.CreatedAt.Format(time.ANSIC), v.UpdatedAt.Format(time.ANSIC), "\n")
		}
	}
}

func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		log.Fatalf("Failed to convert tasks to json: %v\n", err)
	}

	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		log.Fatalf("Failed to update json: %v\n", err)
	}
}
