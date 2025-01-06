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

		if cmd == "exit" {
			code := 1
			if len(args) > 0 {
				code, err = strconv.Atoi(args[0])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error parsing exit code: invalid syntax. Using default exit code 1")
				}
			} else {
				fmt.Fprintln(os.Stderr, "Error parsing exit code: no code provided. Using default exit code 1")
			}
			os.Exit(code)
		}
		fmt.Fprintln(os.Stderr, cmd+": command not found")
	}
}
