package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			exitCommand("0")
		}

		input = strings.TrimSpace(input)
		inputParts := strings.Fields(input)
		if len(inputParts) == 0 {
			continue
		}

		cmd := inputParts[0]
		rawArgs, argsFound := strings.CutPrefix(input, cmd)
		cleanArgs := strings.TrimSpace(rawArgs)

		switch cmd {
		case "exit":
			os.Exit(0)
			if !argsFound {
				cleanArgs = "1"
				fmt.Fprintln(os.Stderr, "Error: No exit code provided. Using default exit code 1")
			}
			exitCommand(cleanArgs)
		case "echo":
			if !argsFound {
				cleanArgs = ""
			}
			echoCommand(cleanArgs)
		case "type":
			if !argsFound {
				cleanArgs = ""
			}
			typeCommand(cleanArgs)
		default:
			fmt.Fprintln(os.Stderr, cmd+": command not found")
		}
	}
}

func exitCommand(args string) {
	exitCode, err := strconv.Atoi(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid syntax. Using default exit code 1")
		exitCode = 1
	}
	os.Exit(exitCode)
}

func echoCommand(args string) {
	fmt.Fprintln(os.Stdout, args)
}

func typeCommand(args string) {
	commandTypes := map[string]string{
		"echo": "builtin",
		"exit": "builtin",
		"type": "builtin",
	}

	inputType, exists := commandTypes[args]
	var inputDescription string

	if exists {
		switch inputType {
		case "builtin":
			inputDescription = " is a shell builtin"
		}
	} else {
		inputDescription = ": not found"
	}
	fmt.Fprintf(os.Stdout, "%s%s\n", args, inputDescription)
}
