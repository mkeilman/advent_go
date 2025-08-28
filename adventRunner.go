package main

import (
	"advent/run"
	"advent/types"
	"advent/utils/cli"
	"advent/utils/debug"

	"flag"
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
var baseMsg = usageStr(modes.Usage())


func exit(code int, msg string) {
	usage(msg)
	os.Exit(code)
}

func main() {


	args := os.Args[1:]

	if len(args) < 2 {
		run.Exit(ERR_NUM_ARGS, baseMsg)
	}
	
	year, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		run.Exit(run.ERR_INVALID_YEAR, baseMsg)
	}
	day, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		run.Exit(run.ERR_INVALID_DAY, baseMsg)
	}


	runArgs := []string{}

	if len(args) > 2 {
		runArgs = args[2:]
	}
	d, err := NewDay(int(year), int(day))
	if err != nil {
		run.Exit(run.ERR_INVALID_YEAR, err.Error())
	}

	fs := d.DayRunner.Flags()
	fs.Var(*modes, "mode", modes.Usage())
	fs.Usage = usageFn(fs)
	fs.Parse(runArgs)

	if slices.Contains([]string{"test", "all"}, *modes.Val) {
		debug.DebugPrintln("TEST:")
		run.RunFromTestInput(*d)
	}

	if slices.Contains([]string{"file", "all"}, *modes.Val) {
		debug.DebugPrintln("FILE:")
		run.RunFromFile(*d)
	}
}

// Run using local test input
//
// Args:
//		d (types.AdventDay): day struct
//
func RunFromTestInput(d types.AdventDay) any {
	return d.DayRunner.Run(d.DayRunner.TestInput())
}

// Run using input from file
//
// Args:
//		d (types.AdventDay): day struct
//
func RunFromFile(d types.AdventDay) any {

	content, err := os.ReadFile(d.InputFile)
	if err != nil {
		exit(ERR_INVALID_FILE, baseMsg)
	}
	return d.DayRunner.Run(strings.Split(string(content), "\n"))
}

func usage(msg string) {
	fmt.Fprint(os.Stderr, msg)
}

func usageFn(fs *flag.FlagSet) func() {
	u := func() string {
		s := ""
		f := func(fl *flag.Flag) {
			s = s + fmt.Sprintf("%s ", fl.Usage)
		}
		fs.VisitAll(f)
		return s
	}
	return func() { fmt.Fprint(os.Stderr, usageStr(u())) }
}

func usageStr(s string) string {
	return fmt.Sprintf("Usage: go run advent <year> <day> %s", s)
}
