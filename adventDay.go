package main

import (
	"fmt"
)

// Base class for all advent "days"
type adventDay struct {
	input     []string
	inputFile string
	testInput []string
	TEST      []string
}

// ADD RUNARGS
func New(year int, day int) adventDay {
	return adventDay{[]string{}, fmt.Sprintf("year_{%d}/input_day_{%02d}.txt", year, day), []string{}, []string{}}
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
func (d adventDay) Run(input []string) int {
	d.input = input
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
func (d adventDay) RunFromTestInput(input []string) int {
	if len(input) > 0 {
		return d.Run(input)
	} else {
		return d.Run(d.testInput)
	}
}

// Run using input from file
func (d adventDay) RunFromFile() int {

	return d.Run([]string{})
	//with open(d.input_file, "r") as f:
	//	d.input = [x.strip() for x in f.readlines()]
	//	return d.run()
}
