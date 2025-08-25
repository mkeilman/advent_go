package main

import (
	"advent/utils/debug"
	"advent/year2024"
	"fmt"
)

var adventRunners = map[int]map[int]func(runArgs []string)any {
	2024: {
		11: year2024.NewDay11,
	},
}

type Runner interface {
	Config(cfgFileName string, ijnpuFileName string)
	InputFile() string
	Run(input []string) int
	TestInput() []string
}


// ADD RUNARGS (?)
func NewDay(year int, day int, runArgs []string) *Runner {
	debug.DebugPrintln("DAY ARGS %s", runArgs)

	d := adventRunners[year][day](runArgs)
	d.(Runner).Config(configFile(year, day), inputFile(year, day))
	return d.(*Runner)
}


// Run using local test input
//
// Args:
//
//	input ([]string): array of strings to use as test input
//func RunFromTestInput(d Runner, input []string) int {
//	if len(input) > 0 {
//		return d.Run(input)
//	} else {
//		return d.Run(d.data.TestInput)
//	}
//}

// Run using input from file
//func RunFromFile(d Runner) int {
//
//	return d.Run([]string{})
//	content, err = os.ReadFile(d.data.InputFile)
//	if err != nil {
//
//	}
	//with open(d.input_file, "r") as f:
	//	d.input = [x.strip() for x in f.readlines()]
	//	return d.run()
//}

func configFile(year int, day int) string {
	return fmt.Sprintf("year%d/day%02d.toml", year, day)
}

func inputFile(year int, day int) string {
	return fmt.Sprintf("year%d/inputDay%02d.txt", year, day)
}