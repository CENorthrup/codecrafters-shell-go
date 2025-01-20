package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const errDefaultExitCode string = "Using default exit code 1"

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
		// TODO: add alias case
		case "builtin":
			runBuiltinCommand(cmd, args, argsExist)
		case "executable":
			runExternalProgram(cmd, args)
		default:
			fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
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
	if checkBuiltin(cmd) {
		return "builtin", ""
	}
	if cmdPath := checkExecutable(cmd); cmdPath != "" {
		return "executable", cmdPath
	}
	return "", ""
}

func checkBuiltin(cmd string) bool {
	builtin := map[string]struct{}{
		"cd":   {},
		"echo": {},
		"exit": {},
		"pwd":  {},
		"type": {},
	}
	_, exists := builtin[cmd]
	return exists
}

// TODO: add alias

func checkExecutable(args string) string {
	cmdPath, err := exec.LookPath(args)
	if err != nil {
		return ""
	}
	return cmdPath
}

func runExternalProgram(cmd string, args string) {
	argSlice := strings.Fields(args)
	cmdPath := checkExecutable(cmd)
	if cmdPath == "" {
		fmt.Fprintf(os.Stderr, "Error: Path to command: %s not found", cmd)
		return
	}
	command := exec.Command(cmd, argSlice...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Command finished with error: %s", err)
	}
}

func runBuiltinCommand(cmd string, args string, argsExist bool) {
	switch cmd {
	case "cd":
		cdCommand(args)
	case "exit":
		if argsExist {
			exitCommand(args)
		} else {
			fmt.Fprintf(os.Stderr, "Error: No exit code provided. %s\n", errDefaultExitCode)
			exitCommand(1)
		}
	case "echo":
		echoCommand(args)
	case "type":
		typeCommand(args)
	case "pwd":
		pwdCommand()
	default:
		return
	}
}

// TODO: add aliasCommand()

func cdCommand(args string) {
	err := os.Chdir(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", args)
		// fmt.Fprintf(os.Stderr, "Error: Unable to change directory %s\n", err)
	}
}

func echoCommand(args string) {
	fmt.Fprintln(os.Stdout, args)
}

func exitCommand[T string | int](arg T) {
	exitCode := 1
	switch argVal := any(arg).(type) {
	case string:
		if argCode, err := strconv.Atoi(argVal); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid syntax. %s\n", errDefaultExitCode)
		} else {
			exitCode = argCode
		}
	case int:
		os.Exit(argVal)
	default:
		fmt.Fprintf(os.Stderr, "Error: Unsupported type. %s\n", errDefaultExitCode)
	}
	os.Exit(exitCode)
}

func pwdCommand() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Unable to get present working directory %s\n", err)
		return
	}
	fmt.Fprintln(os.Stdout, pwd)
}

func typeCommand(cmd string) {
	switch cmdType, cmdArgs := checkCmdType(cmd); cmdType {
	case "builtin":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
		// TODO: add alias case
	case "executable":
		fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, cmdArgs)
	default:
		fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
	}
}
