package main

import (
	"advent/types"
	//"advent/utils/debug"
	"advent/year2024"
	"fmt"
)

// move to types?
var adventRunners = map[int]map[int]types.Runner {
	2024: {
		11: &year2024.AdventDay11{},
	},
}

// ADD RUNARGS (?)
func NewDay(year int, day int) *types.AdventDay {
	d := types.AdventDay{
		Year: year,
		Day: day,
		InputFile: inputFile(year, day),
		DayRunner: adventRunners[year][day],
	}
	return &d
}

func inputFile(year int, day int) string {
	return fmt.Sprintf("year%d/inputDay%02d.txt", year, day)
}