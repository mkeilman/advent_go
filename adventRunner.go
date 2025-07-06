package main

import (
	"advent/utils/debug"
	"advent/utils/cli"
	"flag"
	"fmt"
	"os"
	"slices"
)

var MODES = []string{"test", "file", "all"}


func main() {

	m := cli.NewArgsList(MODES, "all")

	usage := func() {
		fmt.Fprintf(os.Stderr, "Usage: go run advent %s <year> <day>\n", m.Usage())
	}

	flag.Var(*m, "mode", "run mode")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		usage()
		os.Exit(1)
	}


	//d := New(2024, 11)
	//var d Runner

	if slices.Contains([]string{"test", "all"}, *m.Val) {
		debug.DebugPrintln("TEST:")
		//d.run_from_test_input()
	}

	if slices.Contains([]string{"file", "all"}, *m.Val) {
		debug.DebugPrintln("FILE:")
		//d.run_from_file()
	}

	//debug.DebugPrintln("A: %s F: %s", os.Args[1:], flag.Args())
}
