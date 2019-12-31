package help

import (
	"bytes"
	"io"
	"os"
	"testing"

	"gotest.tools/assert"
)

var (
	out io.Writer = os.Stdout
)

func TestShow(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	Show(out)

	assert.Equal(t, `
Usage: bepatient <command> [args]

The available commands for execution are listed below.

Available Commands:
  help         Showing usage
  version      Print the version information
  wait         Wait for all pods to be ready

`, out.(*bytes.Buffer).String())
}

func TestShowHowToWait(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	ShowHowToWait(out)

	assert.Equal(t, `
Wait for all pods to be ready

Usage:
  bepatient wait [flags] NAMESPACE

Flags:
  -n, --namespace string      The namespace to be wait

`, out.(*bytes.Buffer).String())
}
