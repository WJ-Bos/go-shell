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
	builtIns := map[string]bool{
		"echo": true,
		"type": true,
		"exit": true,
	}

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

		if !builtIns[command] {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
		}

		switch command {
		case "echo":
			{
				if len(parts) > 1 {
					toPrint := parts[1:]
					fmt.Println(strings.Join(toPrint, " "))
				} else {
					fmt.Println(" ")
				}
			}
		case "type":
			{
				if len(parts) > 1 {
					if !builtIns[parts[1]] {
						fmt.Fprintf(os.Stderr, "%s: not found\n", parts[1])
					} else {
						fmt.Println(parts[1] + " is a shell builtin")
					}
				} else {
					fmt.Println("To use type command please provide 'type' followed by a shell builtin.")
				}
			}

		}
	}

}
