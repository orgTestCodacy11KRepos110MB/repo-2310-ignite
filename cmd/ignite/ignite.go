package main

import (
	"fmt"
	"os"

	"github.com/luxas/ignite/cmd/ignite/cmd"
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// Run runs the main cobra command of this application
func Run() error {
	c := cmd.NewIgniteCommand(os.Stdin, os.Stdout, os.Stderr)
	return c.Execute()
}