package main

import (
	"fmt"
	"os"

	"github.com/evercoinx/bitcoin/cmd/cli/commands"
	"github.com/evercoinx/bitcoin/cmd/cli/flags"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "bitcoin",
		Usage:    "toolkit for operations with bitcoin blockchain",
		Commands: commands.New(),
		Flags:    flags.New(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
