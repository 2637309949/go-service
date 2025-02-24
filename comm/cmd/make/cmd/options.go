package cmd

import (
	"context"

	"github.com/urfave/cli/v2"
)

type Option func(o *Options)

type Options struct {
	// Name of the application
	Name string
	// Description of the application
	Description string
	// Version of the application
	Version string
	// Action to execute when Run is called and there is no subcommand
	Action func(*cli.Context) error
	// TODO replace with built in command definition
	Commands []*cli.Command
	// TODO replace with built in flags definition
	Flags []cli.Flag
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type beforeKey struct{}
type serviceKey struct{}

func beforeFromContext(ctx context.Context, def cli.BeforeFunc) cli.BeforeFunc {
	if ctx == nil {
		return def
	}

	a, ok := ctx.Value(beforeKey{}).(cli.BeforeFunc)
	if !ok {
		return def
	}

	// perform the before func passed in the context before the default
	return func(ctx *cli.Context) error {
		if err := a(ctx); err != nil {
			return err
		}
		return def(ctx)
	}
}

func serviceFromContext(ctx context.Context) bool {
	if ctx == nil {
		return false
	}

	a, _ := ctx.Value(serviceKey{}).(bool)
	return a
}
