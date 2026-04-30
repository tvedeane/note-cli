# note-cli

`note-cli` is a small command-line tool for saving and listing personal notes
from the terminal.

## Status

The project currently supports adding and listing notes. Deleting notes is not
implemented yet.

## Requirements

- Go 1.22 or newer

## Install From Source

Clone the repository and build the binary:

```sh
make build
```

The compiled binary is written to `bin/note`.

You can also run the tool without building it first:

```sh
make run
make run add "remember to write tests"
make run list
```

## Usage

```sh
note add "remember to write tests"
note list
note help
note version
```

## Commands

| Command | Description |
| --- | --- |
| `note add <note>` | Add a note. |
| `note list` | List all saved notes. |
| `note help` | Show help text. |
| `note version` | Show the application version. |

## Storage

Notes are stored under `.notes/db`. Each note is written to a file named with
the SHA-256 hash of the note text.

## Development

```sh
make test
make build
make run
```

## Repository Structure

- `cmd/note`: application entry point
- `internal/cli`: CLI application wiring and behavior
