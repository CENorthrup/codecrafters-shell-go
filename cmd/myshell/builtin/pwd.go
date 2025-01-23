package builtin

import (
	"fmt"
	"os"
)

func Pwd() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Unable to get present working directory %s\n", err)
		return
	}
	fmt.Fprintln(os.Stdout, pwd)
}
