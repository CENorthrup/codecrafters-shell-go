package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		cmd := parseCommand(input)
		args, argsExist := parseArgs(cmd, input)

		switch cmdType, _ := checkCmdType(cmd); cmdType {
		case "builtin":
			runBuiltinCommand(cmd, args, argsExist)
			// TODO: add alias case
		case "executable":

			runExternalProgram(cmd, args)
		default:
			fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd)
		}
	}
}

func parseCommand(input string) string {
	inputParts := strings.Fields(input)
	return inputParts[0]
}

func parseArgs(cmd string, input string) (args string, argsExist bool) {
	rawArgs, argsFound := strings.CutPrefix(input, cmd)
	return strings.TrimSpace(rawArgs), argsFound
}

func checkCmdType(cmd string) (cmdType string, cmdArgs string) {
	builtin := map[string]struct{}{
		"echo": {},
		"exit": {},
		"type": {},
	}
	if _, exists := builtin[cmd]; exists {
		return "builtin", ""
	}
	if cmdPath := checkExecutable(cmd); cmdPath != "" {
		return "executable", cmdPath
	}
	return "", ""
}

func checkExecutable(args string) string {
	if cmdPath, err := exec.LookPath(args); err != nil {
		fmt.Fprintf(os.Stderr, "%s: not found\n", args)
		return ""
	} else {
		return cmdPath
	}
}

func runExternalProgram(cmd string, args string) {
	argSlice := strings.Fields(args)
	fmt.Fprintf(os.Stderr, "Program was passed %d args (including program name).\n", len(argSlice)+1)
	fmt.Fprintf(os.Stderr, "Arg #0 (program name): %s\n", cmd)
	for i, arg := range argSlice {
		fmt.Fprintf(os.Stderr, "Arg #%d: %s\n", i+1, arg)
	}
	cmdPath := checkExecutable(cmd)
	if cmdPath == "" {
		fmt.Fprintf(os.Stderr, "Error: Path to command: %s not found", cmd)
		return
	}
	command := exec.Command(cmdPath, argSlice...)

	out, err := command.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Unable to run program: %s\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "Program Signature: %s\n", strings.TrimSpace(string(out)))
}

func runBuiltinCommand(cmd string, args string, argsExist bool) {
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
	}
}

// TODO: add aliasCommand()
func echoCommand(args string) {
	fmt.Fprintln(os.Stdout, args)
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

func typeCommand(cmd string) {
	switch cmdType, cmdArgs := checkCmdType(cmd); cmdType {
	case "builtin":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
		// TODO: add alias case
	case "executable":
		fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, cmdArgs)
	}
}
