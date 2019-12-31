package version

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

	Version = "1.1.31"
	Show(out)

	assert.Equal(t, "bepatient version: 1.1.31\n", out.(*bytes.Buffer).String())
}
