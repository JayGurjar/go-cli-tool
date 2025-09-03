package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	title string;
	completed bool;
	createdAt time.Time;
	completedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add (title string) {
	todo := Todo {
		title: title,
		completed: false,
		createdAt: time.Now(),
		completedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex (index int) error {

	t := *todos

	if index < 0 || index >= len(t) {
		err := errors.New("Invalid Index")
		fmt.Println("Invalid index")
		return err
	}

	return nil
}

func (todos *Todos) remove (index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index + 1:]...)
	return nil
}

func (todos *Todos) edit (index int, attributeName string, value string) error {

	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	existingTodo := &t[index]

	switch attributeName {
		case "completed":
			completedAt := time.Now()
			existingTodo.completed = true
			existingTodo.completedAt = &completedAt
			break
		case "title":
			existingTodo.title = value
			break
	}

	return nil

}

func (todos *Todos) printTodos() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		var completedAt string = ""
		if t.completed {
			completed = "✅"
			if t.completedAt != nil {
				completedAt = t.completedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.title, completed, t.createdAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
