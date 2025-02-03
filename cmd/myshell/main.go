package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/builtin"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Unable to read input:%v\n", err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		cmd, args, argsExist := utils.TokenizeInput(input)

		switch cmdType, _ := utils.CheckCmdType(cmd); cmdType {
		// TODO: add alias case
		case "builtin":
			runBuiltinCommand(cmd, utils.SliceToString(args), argsExist)
		case "executable":
			runExternalProgram(cmd, args)
		default:
			fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
		}
	}
}

func runBuiltinCommand(cmd string, args string, argsExist bool) {
	switch cmd {
	case "cd":
		builtin.Cd(args)
	case "exit":
		if argsExist {
			builtin.Exit(args)
		} else {
			builtin.Exit(1)
		}
	case "echo":
		builtin.Echo(args)
	case "type":
		builtin.Type(args)
	case "pwd":
		builtin.Pwd()
	default:
		return
	}
}

// TODO: add runAliasCommand()

func runExternalProgram(cmd string, args []string) {
	cmdPath := utils.CheckExecutable(cmd)
	if cmdPath == "" {
		fmt.Fprintf(os.Stderr, "Error: Path to command: %s not found", cmd)
		return
	}

	command := exec.Command(cmd, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Command finished with error: %s\n", err)
	}
}
