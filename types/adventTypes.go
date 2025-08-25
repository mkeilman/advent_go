package types

type AdventDay struct {
	ConfigFile string
	InputFile string
	Year int
	Day int
	DayRunner Runner
}

type Runner interface {
	Config(cfgFileName string)
	Run(input []string) int
	TestInput() []string
}

