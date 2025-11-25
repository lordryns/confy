package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v3" 

	"confy/internal/globals"
	"confy/internal/core"
)


func main() {
	var confErr error
	globals.CONFIG_HOME_PATH, confErr = core.CheckForGlobalConfigPath()
	if confErr != nil {
		log.Warn("Could not detect a global config path!")
	}

	core.CreateConfyModuleIfNotExist()
	var cmd = cli.Command{
		Name: "confy", 
		Usage: "Sort and manage configs easily",
		Commands: []*cli.Command{commandGetConfigPath(), commandSetConfig(), commandGetConfig()},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func commandGetConfig() *cli.Command {
	return &cli.Command{
		Name: "get", 
		Usage: "Get module/config details",
		Flags: []cli.Flag{
			
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			core.GetModuleDetails(c.Args().Get(0))
			return nil
		},
	}
}

func commandSetConfig() *cli.Command {
	return &cli.Command{
		Name: "set", 
		Usage: "Replaces module/config with new one (and backs up the existing one)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "no-backup",
				Value: false,
				Usage: "Prevent backup during set command",
				
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return nil
		},
	}
} 

func commandGetConfigPath() *cli.Command {
	return &cli.Command{
		Name: "path",
		Usage: "See current config home  path",
		Flags: []cli.Flag{
            &cli.StringFlag{
                Name:  "set",
				Aliases: []string{"s"},
                Usage: "Set new config path",
            },
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			var res = c.String("set")

			if  len(res) > 0 {
				log.Infof("Setting new config path to: %v\n", res)
				return nil
			}


			fmt.Println(globals.CONFIG_HOME_PATH)

			return nil
		},
	}
}
