package cmd

import (
	"fmt"
	"os"

	"make/config"

	"github.com/urfave/cli/v2"
)

type Cmd interface {
	// Options set within this command
	Options() Options
	// The cli app within this cmd
	App() *cli.App
	// Run executes the command
	Run() error
}

type command struct {
	opts    Options
	app     *cli.App
	service bool
}

func (c *command) App() *cli.App {
	return c.app
}

func (c *command) Options() Options {
	return c.opts
}

// Before is executed before any subcommand
func (c *command) Before(ctx *cli.Context) error {
	// set the config file if specified
	if cf := ctx.String("c"); len(cf) > 0 {
		config.SetConfig(cf)
	}

	command := ctx.Args().First()

	// certain commands don't require loading
	if command == "env" {
		return nil
	}

	return nil
}

func (c *command) Run() error {
	return c.app.Run(os.Args)
}

func action(c *cli.Context) error {
	return build(c, c.Args().Slice())
}

func New(opts ...Option) *command {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	cmd := new(command)
	cmd.opts = options
	cmd.app = cli.NewApp()
	cmd.app.Name = name
	cmd.app.Version = buildVersion()
	cmd.app.Usage = description
	cmd.app.Flags = defaultFlags
	cmd.app.Action = action
	cmd.app.Before = beforeFromContext(options.Context, cmd.Before)

	// if this option has been set, we're running a service
	// and no action needs to be performed. The CMD package
	// is just being used to parse flags and configure micro.
	if serviceFromContext(options.Context) {
		cmd.service = true
		cmd.app.Action = func(ctx *cli.Context) error { return nil }
	}

	//flags to add
	if len(options.Flags) > 0 {
		cmd.app.Flags = append(cmd.app.Flags, options.Flags...)
	}
	//action to replace
	if options.Action != nil {
		cmd.app.Action = options.Action
	}

	return cmd
}

var (
	DefaultCmd Cmd = New()
	// name of the binary
	name = "make"
	// description of the binary
	description = "build and release service"
	// defaultFlags which are used on all commands
	defaultFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "c",
			Usage:   "Set the config file: Defaults to ./.env",
			EnvVars: []string{"CONFIG_FILE"},
		},
		&cli.StringFlag{
			Name:    "name",
			Usage:   "Set the service name",
			EnvVars: []string{"SERVICE_NAME"},
		},
	}
)

// Run the default command
func Run() {
	if err := DefaultCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
