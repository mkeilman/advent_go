package types

import (
	"flag"
)

type AdventDay struct {
	ConfigFile string
	InputFile string
	Year int
	Day int
	DayRunner Runner
}

type Runner interface {
	Config(cfgFileName string)
	Flags() *flag.FlagSet
	Run(input []string) int
	TestInput() []string
}

