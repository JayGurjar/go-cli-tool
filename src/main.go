package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseAndExecute(input string, todos *Todos) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := parts[0]

	switch command {
	case "list":
		todos.printTodos()
	case "add":
		if len(parts) < 2 {
			fmt.Println("Usage: add <task>")
			return
		}
		task := strings.Join(parts[1:], " ")
		todos.add(task)
		fmt.Println("Task added!")
	case "delete":
		if len(parts) < 2 {
			fmt.Println("Usage: delete <index>")
			return
		}
		index, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid index")
			return
		}
		todos.remove(index)
		fmt.Println("Task deleted!")
	case "toggle":
		if len(parts) < 2 {
			fmt.Println("Usage: toggle <index>")
			return
		}
		index, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid index")
			return
		}
		todos.edit(index, "completed", "")
		fmt.Println("Task toggled!")
	case "edit":
		if len(parts) < 3 {
			fmt.Println("Usage: edit <index> <new_title>")
			return
		}
		index, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid index")
			return
		}
		newTitle := strings.Join(parts[2:], " ")
		todos.edit(index, "title", newTitle)
		fmt.Println("Task updated!")
	default:
		fmt.Println("Unknown command. Available: add, list, delete, toggle, edit, quit")
	}
}

func main() {
	todos := Todos{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command (add, list, delete, toggle, edit, quit): ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "quit" {
			break
		}

		// Parse the input manually and execute commands
		parseAndExecute(input, &todos)
	}
}
