package main

import (
	"fmt"
	"os"

	"github.com/tvedeane/note-cli/internal/cli"
)

var version = "dev"

func main() {
	app := cli.New(cli.Config{
		Out:     os.Stdout,
		Err:     os.Stderr,
		Version: version,
	})

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
