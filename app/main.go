package main

import (
	"fmt"
	"os"
)

var _ = fmt.Fprint
var _ = os.Stdout

func main() {

	fmt.Fprint(os.Stdout, "$ ")
	var cmd string
	_, err := fmt.Scanln(&cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	switch cmd {
	default:
		fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd)
	}
}
