package year2024

import (
	//"advent/utils/debug"
)

//TEST = [
//         "0 1 10 99 999",
//    ]

	
type AdventDay struct {
	input     *[]string
	inputFile string
	testInput []string
	TEST      []string
}



func (d AdventDay) Run(input []string) int {
	*d.input = input
	return 0
}





