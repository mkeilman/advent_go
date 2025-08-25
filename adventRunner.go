package main

import (
	"advent/utils/debug"
	"advent/utils/cli"
	//"advent/utils/collections"

	//"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	//"golang.org/x/tools/go/packages"
)

const (
	ERR_NUM_ARGS = iota + 90
	ERR_INVALID_YEAR
	ERR_INVALID_DAY
	ERR_INVALID_FILE
)

var MODES = []string{"test", "file", "all"}
var modes = cli.NewArgsList(MODES, "all")


func exit(code int) {
	usage()
	os.Exit(code)
}

func main() {

	flag.Var(*modes, "mode", "run mode")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		exit(ERR_NUM_ARGS)
	}

	
	year, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		exit(ERR_INVALID_YEAR)
	}
	day, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		exit(ERR_INVALID_DAY)
	}
	debug.DebugPrintln("YEAR %d DAY %d", year, day)


	runArgs := []string{}
	if len(args) > 2 {
		runArgs = args[2:]
	}
	d := NewDay(int(year), int(day), runArgs)

	if slices.Contains([]string{"test", "all"}, *modes.Val) {
		debug.DebugPrintln("TEST:")
		RunFromTestInput(*d)
	}

	if slices.Contains([]string{"file", "all"}, *modes.Val) {
		debug.DebugPrintln("FILE:")
		RunFromFile(*d)
	}

	debug.DebugPrintln("A: %s F: %s", os.Args[1:], flag.Args())
}

// Run using local test input
//
// Args:
//
//	input ([]string): array of strings to use as test input
func RunFromTestInput(d Runner) int {
	return d.Run(d.TestInput())
}

// Run using input from file
func RunFromFile(d Runner) int {

	content, err := os.ReadFile(d.InputFile())
	if err != nil {
		exit(ERR_INVALID_FILE)
	}
	return d.Run(strings.Split(string(content), "\n"))
}

//func inputFile(year int, day int) string {
//	return fmt.Sprintf("year%d/inputDay%02d.txt", year, day)
//}


func usage() {
	fmt.Fprintf(os.Stderr, "Usage: go run advent %s <year> <day>\n", modes.Usage())
}