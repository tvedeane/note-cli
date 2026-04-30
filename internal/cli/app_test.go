package cli

import (
	"bytes"
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

	if got, want := out.String(), "note-cli test\n"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRunRejectsUnknownCommand(t *testing.T) {
	app := New(Config{})

	if err := app.Run([]string{"missing"}); err == nil {
		t.Fatal("expected error")
	}
}
