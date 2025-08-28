package main

import (
	"advent/types"
	//"advent/utils/debug"
	"advent/year2024"

	"errors"
	"fmt"
)

// move to types?
var adventRunners = map[int]map[int]types.Runner {
	2024: {
		11: &year2024.AdventDay11{},
	},
}

func NewDay(year int, day int) (ad *types.AdventDay, err error) {
	y := adventRunners[year]
	if y == nil {
		return nil, errors.New(fmt.Sprintf("invalid year %d", year))
	}
	r := y[day]
	if r == nil {
		return nil, errors.New(fmt.Sprintf("invalid day %d for year %d", day, year))
	}

	d := types.AdventDay{
		Year: year,
		Day: day,
		InputFile: inputFile(year, day),
		DayRunner: r,
	}
	return &d, nil
}

func inputFile(year int, day int) string {
	return fmt.Sprintf("year%d/inputDay%02d.txt", year, day)
}
