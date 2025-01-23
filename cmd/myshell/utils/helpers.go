package utils

import (
	"os/exec"
	"strings"
)

func ParseCommand(input string) string {
	inputParts := strings.Fields(input)
	return inputParts[0]
}

func ParseArgs(cmd string, input string) (args string, argsExist bool) {
	rawArgs, argsFound := strings.CutPrefix(input, cmd)
	return strings.TrimSpace(rawArgs), argsFound
}

func CheckCmdType(cmd string) (cmdType string, cmdArgs string) {
	if CheckBuiltin(cmd) {
		return "builtin", ""
	}
	if cmdPath := CheckExecutable(cmd); cmdPath != "" {
		return "executable", cmdPath
	}
	return "", ""
}

func CheckBuiltin(cmd string) bool {
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

func CheckExecutable(args string) string {
	cmdPath, err := exec.LookPath(args)
	if err != nil {
		return ""
	}
	return cmdPath
}
