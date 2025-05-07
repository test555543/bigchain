package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const todoFile = "todo.json"

type Todo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func loadTodos() []Todo {
	data, err := ioutil.ReadFile(todoFile)
	if err != nil {
		return []Todo{}
	}
	var todos []Todo
	json.Unmarshal(data, &todos)
	return todos
}

func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	ioutil.WriteFile(todoFile, data, 0644)
}

func listTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No todo items.")
		return
	}
	for i, t := range todos {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, i+1, t.Task)
	}
}

func addTodo(task string) {
	todos := loadTodos()
	todos = append(todos, Todo{Task: task})
	saveTodos(todos)
	fmt.Println("Task added:", task)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [list | add <task>]")
		return
	}

	command := os.Args[1]
	switch command {
	case "list":
		listTodos(loadTodos())
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		task := os.Args[2]
		addTodo(task)
	default:
		fmt.Println("Unknown command:", command)
	}
}
