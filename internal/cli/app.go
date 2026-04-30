package cli

import (
	"fmt"
	"io"
)

const appName = "note-cli"

type Config struct {
	Out     io.Writer
	Err     io.Writer
	Version string
}

type App struct {
	out     io.Writer
	err     io.Writer
	version string
}

func New(config Config) *App {
	if config.Out == nil {
		config.Out = io.Discard
	}
	if config.Err == nil {
		config.Err = io.Discard
	}
	if config.Version == "" {
		config.Version = "dev"
	}

	return &App{
		out:     config.Out,
		err:     config.Err,
		version: config.Version,
	}
}

func (app *App) Run(args []string) error {
	if len(args) == 0 {
		app.printHelp()
		return nil
	}

	switch args[0] {
	case "-h", "--help", "help":
		app.printHelp()
		return nil
	case "-v", "--version", "version":
		fmt.Fprintf(app.out, "%s %s\n", appName, app.version)
		return nil
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func (app *App) printHelp() {
	fmt.Fprintf(app.out, `%s is a small command-line tool for notes.

Usage:
  %s <command> [arguments]

Commands:
  help       Show this help text
  version    Show the application version

Note commands will be added later.
`, appName, appName)
}
