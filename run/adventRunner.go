package run

import (
	"advent/types"

	"fmt"
	"os"
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

	content, err := os.ReadFile(d.InputFile)
	if err != nil {
		Exit(ERR_INVALID_FILE, fmt.Sprintf("invalid input file %s", d.InputFile))
	}
	return d.DayRunner.Run(strings.Split(string(content), "\n"))
}

