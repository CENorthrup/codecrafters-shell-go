package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Unable to read input:%v\n", err)
			exitCommand(1)
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		cmd, args, argsExist := parseInput(input)

		switch cmd {
		case "exit":
			if argsExist {
				exitCommand(args)
			} else {
				fmt.Fprintf(os.Stderr, "Error: No exit code provided. Using default exit code 1\n")
				exitCommand(1)
			}
		case "echo":
			echoCommand(args)
		case "type":
			typeCommand(args)
		default:
			if cmdPath, cmdExists := checkExecutables(args); cmdExists {
				execProgram(cmdPath, args)
			} else {
				fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd)
			}
		}
	}
}

func parseInput(input string) (cmd string, args string, argsExist bool) {
	inputParts := strings.Fields(input)
	cmd = inputParts[0]
	rawArgs, argsFound := strings.CutPrefix(input, cmd)
	return cmd, strings.TrimSpace(rawArgs), argsFound
}

func exitCommand[T string | int](arg T) {
	exitCode := 1
	switch argVal := any(arg).(type) {
	case string:
		if argCode, err := strconv.Atoi(argVal); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid syntax. Using default exit code 1\n")
		} else {
			exitCode = argCode
		}
	case int:
		os.Exit(argVal)
	default:
		fmt.Fprintf(os.Stderr, "Error: Unsupported type. Using default exit code 1\n")
	}
	os.Exit(exitCode)
}

func echoCommand(args string) {
	fmt.Fprintln(os.Stdout, args)
}

func typeCommand(args string) {
	builtin := map[string]struct{}{
		"echo": {},
		"exit": {},
		"type": {},
	}
	if _, exists := builtin[args]; exists {
		fmt.Printf("%s is a shell builtin\n", args)
		return
	}
	if fullPath, found := checkExecutables(args); found {
		fmt.Printf("%s is %s\n", args, fullPath)
		return
	}
	fmt.Fprintf(os.Stderr, "%s: not found\n", args)
}

func checkExecutables(args string) (string, bool) {
	// cmdPath, err := exec.LookPath(args)
	// if err != nil {
	//   fmt.Fprintf(os.Stderr, "Error: Unable to locate executable on PATH: %v\n", err)
	//   return "", false
	// }
	// fullPath := filepath.Join(cmdPath, args)
	// return fullPath, true

	pathEnv := os.Getenv("PATH")
	pathEntries := strings.Split(pathEnv, ":")

	for _, dir := range pathEntries {
		dirContents, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			fmt.Fprintf(os.Stderr, "Error: Unable to access or read PATH directory %s: %v\n", dir, err)
			continue
		}

		for _, dirEntry := range dirContents {
			if dirEntry.Name() == args {
				fullPath := filepath.Join(dir, dirEntry.Name())
				dirEntryInfo, err := dirEntry.Info()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: Unable to retrieve file info for %s: %v\n", dirEntry.Name(), err)
					return "", false
				}
				if dirEntryInfo.Mode()&0111 != 0 {
					return fullPath, true
				}
			}
		}
	}
	return "", false
}

func execProgram(cmdPath string, args string) {
	exec.Command(cmdPath, args)
}
