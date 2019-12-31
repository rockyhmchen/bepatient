package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/rockyhmchen/bepatient/help"
	"github.com/rockyhmchen/bepatient/version"
	"github.com/rockyhmchen/bepatient/wait"
)

var (
	out io.Writer = os.Stdout
)

func main() {
	waitCmd := flag.NewFlagSet("wait", flag.ExitOnError)
	flag.NewFlagSet("help", flag.ExitOnError)
	flag.NewFlagSet("version", flag.ExitOnError)

	if len(os.Args) < 2 {
		help.Show(out)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "wait":
		wait.Run(out, waitCmd)
	case "help":
		help.Show(out)
	case "version":
		version.Show(out)
	default:
		fmt.Println("Unrecognised command")
		flag.PrintDefaults()
		os.Exit(0)
	}
}
