package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/commands"
)

var _ = fmt.Fprint
var _ = os.Stdout

type Shell struct {
	builtIns map[string]func([]string) error
}

func NewShell() *Shell {
	s := &Shell{
		builtIns: make(map[string]func([]string) error),
	}

	s.builtIns["echo"] = commands.Echo
	s.builtIns["exit"] = commands.Exit
	s.builtIns["type"] = func(args []string) error {
		return commands.Type(s.getBuiltInNames(), args)
	}

	return s
}

func (s *Shell) getBuiltInNames() map[string]bool {
	names := make(map[string]bool)
	for name := range s.builtIns {
		names[name] = true
	}
	return names
}

func (s *Shell) executeCommand(command string, args []string) error {
	if handler, exists := s.builtIns[command]; exists {
		return handler(args)
	}
	return fmt.Errorf("%s: command not found", command)
}

func (s *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		command, args := parts[0], parts[1:]

		if err := s.executeCommand(command, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func main() {
	shell := NewShell()
	shell.Run()
}
