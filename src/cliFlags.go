package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type CLIFlags struct {
	// Help string
	Add string
	List bool
	Edit string
	Delete int
	Toggle int
}

type NewCmdFlags []CLIFlags

func cmdFlags () *CLIFlags {
	cliFlags := CLIFlags{}

	// flag.StringVar(&cliFlags.Help, "help", "", "Available commands")
	flag.StringVar(&cliFlags.Add, "add", "", "Todo Task")
	flag.IntVar(&cliFlags.Delete, "delete", -1, "Index to of todo to delete")
	flag.BoolVar(&cliFlags.List, "list", true, "List All Todos")
	flag.IntVar(&cliFlags.Toggle, "toggle", -1, "Index to toggle")
	flag.StringVar(&cliFlags.Edit, "edit", "", "Edit a todo with index specified. Ex: index:todoString")

	flag.Parse()

	return &cliFlags
}

func (cf *CLIFlags) execute (todos *Todos) {
	switch {
		case cf.List:
			todos.printTodos()
		case cf.Add != "":
			todos.add(cf.Add)
		case cf.Edit != "":
			parts := strings.Split(cf.Edit, ":")
			if len(parts) != 2 {
				fmt.Println("Error: Invalid edit syantax. Use index:todoTitle")
				os.Exit(1)
			}
			index, error := strconv.Atoi(parts[0])

			if error != nil {
				fmt.Print("Invalid index for edit")
				os.Exit(1)
			}

			todos.edit(index, "title", parts[1])
		case cf.Toggle != -1:
			todos.edit(cf.Toggle, "completed", "")
		case cf.Delete != -1:
			todos.remove(cf.Delete)
		default:
			fmt.Println("Invalid command")
	}
}
