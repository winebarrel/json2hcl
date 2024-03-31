package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	version string
)

type flags struct {
	file string
}

func init() {
	cmdLine := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	cmdLine.Usage = func() {
		fmt.Fprintf(cmdLine.Output(), "Usage: %s [OPTION] [FILE]\n", cmdLine.Name())
		cmdLine.PrintDefaults()
	}

	flag.CommandLine = cmdLine
}

func parseFlags() *flags {
	flags := &flags{}
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		printVersionAndExit()
	}

	args := flag.Args()

	if len(args) == 1 {
		flags.file = args[0]
	} else if len(args) >= 2 {
		printUsageAndExit()
	}

	return flags
}

func printVersionAndExit() {
	v := version

	if v == "" {
		v = "<nil>"
	}

	fmt.Fprintln(flag.CommandLine.Output(), v)
	os.Exit(0)
}

func printUsageAndExit() {
	flag.CommandLine.Usage()
	os.Exit(0)
}
