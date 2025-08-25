package types

type AdventData struct {
	Input     *[]string
	InputFile *string
	TestInput []string
	TEST      []string
}

type Runner interface {
	Config(cfgFileName string, ijnpuFileName string)
	InputFile() string
	Run(input []string) int
	TestInput() []string
}
