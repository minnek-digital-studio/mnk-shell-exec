package mnkShellExec_test

import (
	"strings"
	"testing"

	mnkShellExec "github.com/minnek-digital-studio/mnk-shell-exec"
)

func TestOut(t *testing.T) {
	stdout, stderr, err := mnkShellExec.Out("echo 'Hello, World!'")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.Contains(stdout, "Hello, World!") {
		t.Errorf("Expected 'Hello, World!' in stdout, got: %s", stdout)
	}

	if stderr != "" {
		t.Errorf("Expected empty stderr, got: %s", stderr)
	}
}

func TestOutLive(t *testing.T) {
	outStr, errStr, err := mnkShellExec.OutLive("echo 'Hello, World!'")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.Contains(outStr, "Hello, World!") {
		t.Errorf("Expected 'Hello, World!' in outStr, got: %s", outStr)
	}

	if errStr != "" {
		t.Errorf("Expected empty errStr, got: %s", errStr)
	}
}

func TestExecuteCommand(t *testing.T) {
	output := mnkShellExec.ExecuteCommand("echo 'Hello, World!'", false)

	if !strings.Contains(output, "Hello, World!") {
		t.Errorf("Expected 'Hello, World!' in output, got: %s", output)
	}
}
