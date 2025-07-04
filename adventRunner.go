package main

import (
	"advent/utils/debug"
	"advent/utils/cli"
	"flag"
	//"os"
	"slices"
)

var MODES = []string{"test", "file", "all"}


func main() {

	m := cli.NewArgsList(MODES, "all")
	//if m == nil {
	//	os.Exit(1)
	//}

	flag.Var(m, "mode", "run mode")

	flag.Parse()

	debug.DebugPrintln("VAL %s", *m.Val)

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

	//debug.DebugPrintln("%s", flag.Args())
}
