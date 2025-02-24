package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func UnexpectedCommand(ctx *cli.Context) error {
	commandName := ctx.Args().First()
	return cli.Exit(fmt.Sprintf("Unrecognized command: %s. Please refer to 'make --help'", commandName), 1)
}

func MissingCommand(ctx *cli.Context) error {
	return cli.Exit("No command provided to make. Please refer to 'make --help'", 1)
}

func NotFoundMakeFile(ctx *cli.Context) error {
	return cli.Exit("No Makefile.bat file provided to make", 1)
}
