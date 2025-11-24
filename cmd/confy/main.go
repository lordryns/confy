package main

import (
	"context"
	"os"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v3"
)


func main() {
	var cmd = cli.Command{
		Name: "confy", 
		Usage: "Sort and manage configs easily",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
