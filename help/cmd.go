package help

import (
	"fmt"
	"io"
)

func Show(out io.Writer) {
	fmt.Fprintln(out, helpText())
}

func ShowHowToWait(out io.Writer) {
	fmt.Fprintln(out, waitInstruction())
}

func helpText() string {
	return fmt.Sprintf(`
Usage: bepatient <command> [args]

The available commands for execution are listed below.

Available Commands:
  help         Showing usage
  version      Print the version information
  wait         Wait for all pods to be ready
`)
}

func waitInstruction() string {
	return fmt.Sprintf(`
Wait for all pods to be ready

Usage:
  bepatient wait [flags] NAMESPACE

Flags:
  -n, --namespace string      The namespace to be wait
`)
}
