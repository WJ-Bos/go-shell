package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var _ = fmt.Fprint
var _ = os.Stdout

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		if parts[0] == "exit" && len(parts) == 2 {
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid exit code:", parts[1])
				continue
			}
			os.Exit(num)
		}

		command := parts[0]
		fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
	}

}
