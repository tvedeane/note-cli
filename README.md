# note-cli

A small command-line tool that can be used to save, list, and delete notes

## Requirements

- Go 1.22 or newer

## Project Structure

- `cmd/note-cli`: application entry point
- `internal/cli`: CLI application wiring

`cmd/` contains executable entry points. The `cmd/note-cli` package is the
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
```

The build output is written to `bin/note-cli`.
