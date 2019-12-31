package version

import (
	"fmt"
	"io"
)

var (
	Version string = ""
)

func Show(out io.Writer) {
	fmt.Fprintln(out, versionText())
}

func versionText() string {
	return fmt.Sprintf("bepatient version: %s", Version)
}
