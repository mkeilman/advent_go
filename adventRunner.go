package main

import (
	"advent/types"
	"advent/utils/debug"
	"advent/utils/cli"

	//"errors"
	//"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	// 
	args := os.Args[1:]

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


	runArgs := []string{}

	if len(args) > 2 {
		runArgs = args[2:]
	}
	d := NewDay(int(year), int(day), runArgs)

	fs := d.DayRunner.Flags()
	fs.Var(*modes, "mode", "run mode")
	fs.Parse(runArgs)

	if slices.Contains([]string{"test", "all"}, *modes.Val) {
		debug.DebugPrintln("TEST:")
		RunFromTestInput(*d)
	}

	if slices.Contains([]string{"file", "all"}, *modes.Val) {
		debug.DebugPrintln("FILE:")
		RunFromFile(*d)
	}
}

// Run using local test input
//
// Args:
//		d (types.AdventDay): day struct
func RunFromTestInput(d types.AdventDay) int {
	return d.DayRunner.Run(d.DayRunner.TestInput())
}

// Run using input from file
//
// Args:
//		d (types.AdventDay): day struct
func RunFromFile(d types.AdventDay) int {

	content, err := os.ReadFile(d.InputFile)
	if err != nil {
		exit(ERR_INVALID_FILE)
	}
	return d.DayRunner.Run(strings.Split(string(content), "\n"))
}


func usage() {
	fmt.Fprintf(os.Stderr, "Usage: go run advent <year> <day> %s\n", modes.Usage())
}