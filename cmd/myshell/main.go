package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stdout, "Error reading input:", err)
		os.Exit(1)
	}
	fmt.Println(input[:len(input)-1] + ": command not found")
}
