package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const appName = "note"
const defaultNotesDir = ".notes/db"

type Config struct {
	Out      io.Writer
	Err      io.Writer
	NotesDir string
	Version  string
}

type App struct {
	out      io.Writer
	err      io.Writer
	notesDir string
	version  string
}

func New(config Config) *App {
	if config.Out == nil {
		config.Out = io.Discard
	}
	if config.Err == nil {
		config.Err = io.Discard
	}
	if config.NotesDir == "" {
		config.NotesDir = defaultNotesDir
	}
	if config.Version == "" {
		config.Version = "dev"
	}

	return &App{
		out:      config.Out,
		err:      config.Err,
		notesDir: config.NotesDir,
		version:  config.Version,
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
	case "add":
		return app.addNote(args[1:])
	case "delete":
		return app.deleteNote(args[1:])
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func (app *App) addNote(args []string) error {
	note := strings.TrimSpace(strings.Join(args, " "))
	if note == "" {
		return fmt.Errorf("usage: %s add <note>", appName)
	}

	hash := hashNote(note)
	if err := os.MkdirAll(app.notesDir, 0o755); err != nil {
		return fmt.Errorf("create notes directory: %w", err)
	}

	path := filepath.Join(app.notesDir, hash)
	if err := os.WriteFile(path, []byte(note+"\n"), 0o644); err != nil {
		return fmt.Errorf("write note: %w", err)
	}

	fmt.Fprintf(app.out, "added note %s\n", hash)
	return nil
}

func (app *App) deleteNote(args []string) error {
	if len(args) != 1 || strings.TrimSpace(args[0]) == "" {
		return fmt.Errorf("usage: %s delete <hash>", appName)
	}

	hash := strings.TrimSpace(args[0])
	path := filepath.Join(app.notesDir, hash)
	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("note %s not found", hash)
		}
		return fmt.Errorf("delete note: %w", err)
	}

	fmt.Fprintf(app.out, "deleted note %s\n", hash)
	return nil
}

func hashNote(note string) string {
	sum := sha256.Sum256([]byte(note))
	return hex.EncodeToString(sum[:])
}

func (app *App) printHelp() {
	fmt.Fprintf(app.out, `%s is a small command-line tool for notes.

Usage:
  %s <command> [arguments]

Commands:
  add <note>      Add a note
  delete <hash>   Delete a note
  help            Show this help text
  version         Show the application version

Notes are stored in .notes/db with the note hash as the filename.
`, appName, appName)
}
