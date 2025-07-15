package main

import (
	"fmt"
	//"reflect"
	//"golang.org/x/tools/go/packages"
)

type Runner interface {
	Run(input []string) int
	RunFromTestInput(input []string) int
	RunFromFile() int
}

// Base class for all advent "days"
type AdventDay struct {
	input     *[]string
	inputFile string
	testInput []string
	TEST      []string
}

// ADD RUNARGS
func New(year int, day int) *AdventDay {
	return &AdventDay{inputFile: fmt.Sprintf("year%d/inputDay%02d.txt", year, day), TEST: []string{}}
}

// ADD RUNARGS
func Build(year int, day int) {
	//TODO: build "subclass"
}

// Generic run method. Subclasses will override this to perform specific calculations.
//
// Args:
//
//	input ([]string): array of strings to use as input
func (d AdventDay) Run(input []string) int {
	*d.input = input
	return 0
}

//Adds command-line arguments to the instance
//
//Args:
//	run_args (dict): command line arguments
//    def add_args(self, run_args):
//
//        v = vars(self.args_parser.parse_args(run_args))
//        for arg in v:
//            setattr(self, arg, v[arg])

// Run using local test input
//
// Args:
//
//	input ([]string): array of strings to use as test input
func (d AdventDay) RunFromTestInput(input []string) int {
	if len(input) > 0 {
		return d.Run(input)
	} else {
		return d.Run(d.testInput)
	}
}

// Run using input from file
func (d AdventDay) RunFromFile() int {

	return d.Run([]string{})
	//with open(d.input_file, "r") as f:
	//	d.input = [x.strip() for x in f.readlines()]
	//	return d.run()
}
