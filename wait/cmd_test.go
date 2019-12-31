package wait

import (
	"bytes"
	"flag"
	"io"
	"os"
	"testing"

	"gotest.tools/assert"
)

var (
	out io.Writer = os.Stdout
)

func TestRunWithoutFlag(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	fs := &flag.FlagSet{}
	Run(out, fs)

	assert.Equal(t, `
Wait for all pods to be ready

Usage:
  bepatient wait [flags] NAMESPACE

Flags:
  -n, --namespace string      The namespace to be wait

`, out.(*bytes.Buffer).String())
}
