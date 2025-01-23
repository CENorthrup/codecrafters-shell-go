package builtin

import (
	"fmt"
	"os"
)

func Cd(path string) {
	cdPath := path
	if path == "~" {
		if cdPath = os.Getenv("HOME"); cdPath == "" {
			fmt.Fprint(os.Stderr, "Error: $HOME not set.\n")
			return
		}
	}
	err := os.Chdir(cdPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", path)
		// fmt.Fprintf(os.Stderr, "Error: Unable to change directory %s\n", err)
	}
}
