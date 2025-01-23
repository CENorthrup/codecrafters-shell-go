package builtin

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

func Type(cmd string) {
	switch cmdType, cmdArgs := utils.CheckCmdType(cmd); cmdType {
	case "builtin":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
		// TODO: add alias case
	case "executable":
		fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, cmdArgs)
	default:
		fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
	}
}
