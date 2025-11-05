package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Echo(args []string) error {
	if len(args) > 0 {
		fmt.Println(strings.Join(args, " "))
	} else {
		fmt.Println()
	}
	return nil
}

func Type(builtIns map[string]bool, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("type: missing argument")
	}

	if builtIns[args[0]] {
		fmt.Printf("%s is a shell builtin\n", args[0])
	} else {
		fmt.Fprintf(os.Stderr, "%s: not found\n", args[0])
	}
	return nil
}

func Exit(args []string) error {
	if len(args) == 0 {
		os.Exit(0)
	}

	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid exit code: %s", args[0])
	}
	os.Exit(exitCode)
	return nil
}
