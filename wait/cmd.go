package wait

import (
	"flag"
	"io"
	"os"

	"github.com/rockyhmchen/bepatient/help"
)

func Run(out io.Writer, cmd *flag.FlagSet) {
	ns := getNamespace(cmd)
	if ns == "" {
		help.ShowHowToWait(out)
		return
	}

	WaitFor(ns)
}

func getNamespace(cmd *flag.FlagSet) string {
	fullFlag := cmd.String("namespace", "", "")
	shortFlag := cmd.String("n", "", "")

	cmd.Parse(os.Args[2:])
	if !cmd.Parsed() {
		return ""
	}

	if *fullFlag == "" && *shortFlag == "" {
		return ""
	}

	ns := *fullFlag
	if ns == "" {
		ns = *shortFlag
	}

	return ns
}
