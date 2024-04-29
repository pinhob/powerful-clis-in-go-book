package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

var todoFileName = ".todo.json"

func main() {
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed following the book Powerful CLIs in Go\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Tasks can be added using the commands `./todo -add + task name` or `echo task name' | ./ todo -add`\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2024\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")

		flag.PrintDefaults()
	}

	add := flag.Bool("add", false, "add a to-do to the list")
	del := flag.Int("del", 0, "Item to be deleted")
	listTasks := flag.Bool("list", false, "List all tasks")
	verb := flag.Bool("verb", false, "List all tasks with the date")
	listNonComplete := flag.Bool("lnc", false, "List all tasks not completed")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	list := &todo.List{}
	if err := list.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listTasks:
		fmt.Print(list)

	case *complete > 0:
		if err := list.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := list.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		list.Add(t)

		if err := list.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *del > 0:
		if err := list.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := list.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *verb:
		for k, item := range *list {
			prefix := ""

			if item.Done {
				prefix = "X "
			}

			fmt.Printf("%s %d: %s | %s\n", prefix, k+1, item.Task, item.CreatedAt.UTC().Local())
		}

	case *listNonComplete:
		listWithoutComplete := list

		for k, item := range *listWithoutComplete {
			if item.Done {
				listWithoutComplete.Delete(k + 1)
			}
		}

		fmt.Print(listWithoutComplete)

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// getTask function decides where to get the description for a new task from: arguments or STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()

	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}
