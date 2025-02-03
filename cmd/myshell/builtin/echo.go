package builtin

import (
	"fmt"
	"os"
)

func Echo(args string) {
	fmt.Fprintln(os.Stdout, args)
}
