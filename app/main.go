package main

import (
	"bufio"
	"fmt"
	"os"
)

var _ = fmt.Fprint
var _ = os.Stdout

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		fmt.Println(command[:len(command)-1] + ": command not found")
	}

}
