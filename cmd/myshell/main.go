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
			// fmt.Fprintf(os.Stdout, "..................%s\n", cmd)
			os.Exit(0)
			if !argsFound {
				// fmt.Println("................... no args found")
				cleanArgs = "1"
				fmt.Fprintln(os.Stderr, "Error: No exit code provided. Using default exit code 1")
			}
			exitCommand(cleanArgs)
		case "echo":
			// fmt.Fprintf(os.Stdout, "..................%s\n", cmd)

			if !argsFound {
				cleanArgs = ""
			}
			echoCommand(cleanArgs)
		default:
			fmt.Fprintln(os.Stderr, cmd+": command not found")
		}
	}
}

func exitCommand(args string) {
	exitCode, err := strconv.Atoi(args)
	// fmt.Printf("........................err: %s", err)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid syntax. Using default exit code 1")
		exitCode = 1
	}
	// fmt.Fprintf(os.Stdout, "Exiting with code %d", exitCode)
	// fmt.Fprint(os.Stdout, "$ ")
	os.Exit(exitCode)
}

func echoCommand(args string) {
	echoString := strings.TrimSpace(args)
	fmt.Fprintln(os.Stdout, echoString)
}
