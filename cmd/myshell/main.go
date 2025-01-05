package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	reader := bufio.NewReader(os.Stdin)

	for {

		cmd, err := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if err != nil {
			fmt.Fprintln(os.Stdout, "Error reading input:", err)
			os.Exit(1)
		}

		fmt.Println(cmd + ": command not found")
		fmt.Fprint(os.Stdout, "$ ")
	}
}
