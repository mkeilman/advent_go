package run

import (
	//"advent/utils/debug"
	"advent/types"

	"fmt"
	"os"
	//"path"
	"path/filepath"
	"runtime"
	//"slices"
	"strings"
)

const (
	ERR_NUM_ARGS = iota + 90
	ERR_INVALID_YEAR
	ERR_INVALID_DAY
	ERR_INVALID_FILE
)

func Exit(code int, msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(code)
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

	_, base, _, _ := runtime.Caller(0)
	content, err := os.ReadFile(filepath.Join(filepath.Dir(base), "..", d.InputFile))
	if err != nil {
		Exit(ERR_INVALID_FILE, fmt.Sprintf("invalid input file %s: %s", d.InputFile, err.Error()))
	}
	return d.DayRunner.Run(strings.Split(string(content), "\n"))
}

