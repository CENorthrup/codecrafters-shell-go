package builtin

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

func Echo(args string) {
	s := utils.StripQuotes(args)
	fmt.Fprintln(os.Stdout, s)
}
