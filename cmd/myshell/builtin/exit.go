package builtin

import (
	"fmt"
	"os"
	"strconv"
)

const errDefaultExitCode string = "Using default exit code 1"

func Exit[T string | int](code T) {
	exitCode := 1
	switch argVal := any(code).(type) {

	case string:
		if argCode, err := strconv.Atoi(argVal); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid syntax. %s\n", errDefaultExitCode)
		} else {
			exitCode = argCode
		}
	case int:
		os.Exit(argVal)
	}
	os.Exit(exitCode)
}
