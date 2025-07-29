package main

import (
	"fmt"
	"gocli/api"
	"gocli/config"
	"gocli/tasks"
	"os"
)

func main() {
	// Load config (from default location or user-specified)
	cfg, err := config.LoadConfig("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Example: Start API server in a goroutine
	go api.StartServer(cfg)

	// Example: Run CLI task manager
	tasks.RunCLI(cfg)
}
