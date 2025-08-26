package types

import (
	"flag"
)

type AdventDay struct {
	InputFile string
	Year int
	Day int
	DayRunner Runner
}

type Runner interface {
	Flags() *flag.FlagSet
	Run(input []string) int
	TestInput() []string
}

