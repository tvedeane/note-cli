package cli

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunPrintsHelpByDefault(t *testing.T) {
	var out bytes.Buffer
	app := New(Config{Out: &out})

	if err := app.Run(nil); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if !strings.Contains(out.String(), "Usage:") {
		t.Fatalf("expected help output, got %q", out.String())
	}
}

func TestRunPrintsVersion(t *testing.T) {
	var out bytes.Buffer
	app := New(Config{Out: &out, Version: "test"})

	if err := app.Run([]string{"version"}); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if got, want := out.String(), "note test\n"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunRejectsUnknownCommand(t *testing.T) {
	app := New(Config{})

	if err := app.Run([]string{"missing"}); err == nil {
		t.Fatal("expected error")
	}
}

func TestRunAddsNote(t *testing.T) {
	var out bytes.Buffer
	notesDir := t.TempDir()
	app := New(Config{Out: &out, NotesDir: notesDir})

	if err := app.Run([]string{"add", "hello", "world"}); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	hash := noteHash("hello world")
	path := filepath.Join(notesDir, hash)
	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected note file: %v", err)
	}

	if got, want := string(contents), "hello world\n"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}

	if got, want := out.String(), "added note "+hash+"\n"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunRejectsEmptyNote(t *testing.T) {
	app := New(Config{NotesDir: t.TempDir()})

	if err := app.Run([]string{"add"}); err == nil {
		t.Fatal("expected error")
	}
}

func TestRunDeletesNote(t *testing.T) {
	var out bytes.Buffer
	notesDir := t.TempDir()
	hash := noteHash("hello world")
	path := filepath.Join(notesDir, hash)
	if err := os.WriteFile(path, []byte("hello world\n"), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	app := New(Config{Out: &out, NotesDir: notesDir})
	if err := app.Run([]string{"delete", hash}); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected note to be deleted, got err %v", err)
	}

	if got, want := out.String(), "deleted note "+hash+"\n"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunRejectsDeleteWithoutHash(t *testing.T) {
	app := New(Config{NotesDir: t.TempDir()})

	if err := app.Run([]string{"delete"}); err == nil {
		t.Fatal("expected error")
	}
}

func TestRunRejectsDeleteForMissingNote(t *testing.T) {
	app := New(Config{NotesDir: t.TempDir()})

	if err := app.Run([]string{"delete", noteHash("missing")}); err == nil {
		t.Fatal("expected error")
	}
}

func noteHash(note string) string {
	sum := sha256.Sum256([]byte(note))
	return hex.EncodeToString(sum[:])
}
