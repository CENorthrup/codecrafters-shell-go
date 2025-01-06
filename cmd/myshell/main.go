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
			os.Exit(1)
		}

		input = strings.TrimSpace(input)
		inputParts := strings.Fields(input)
		if len(inputParts) == 0 {
			continue
		}

		cmd := inputParts[0]
		var args []string
		if len(inputParts) > 1 {
			args = inputParts[1:]
		}

		switch cmd {
		case "exit":
			if len(args) == 0 {
				args[0] = "1"
				fmt.Fprintln(os.Stderr, "Error: No exit code provided. Using default exit code 1")
			}
			exitCommand(args)

		default:
			fmt.Fprintln(os.Stderr, cmd+": command not found")
		}
	}
}

func exitCommand(args []string) {
	code, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid syntax. Using default exit code 1")
		code = 1
	}
	fmt.Fprintf(os.Stdout, "Exiting with code %d", code)
	os.Exit(code)
}
