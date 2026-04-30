# note-cli

A small command-line tool that can be used to save, list, and delete notes

## Status

This repository currently contains the Go CLI project structure and the `add`
and `list` commands. Delete is not implemented yet.

## Requirements

- Go 1.22 or newer

## Project Structure

- `cmd/note`: application entry point
- `internal/cli`: CLI application wiring

`cmd/` contains executable entry points. The `cmd/note` package is the
small `main` package that reads process inputs, wires stdout/stderr, runs the
CLI app, and exits with the right status code.

`internal/` contains private application packages. Go prevents packages inside
`internal` from being imported by code outside this module, so `internal/cli` is
where note-cli can keep command wiring and future app behavior without exposing
it as a public library API.

## Development

```sh
make test
make build
make run
make run add "remember to write tests"
make run list
```

The build output is written to `bin/note`.

## Usage

```sh
note add "remember to write tests"
note list
```

Added notes are stored under `.notes/db`. Each note filename is the SHA-256 hash
of the note text.
