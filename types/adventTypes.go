package types

import (
	"flag"
	"fmt"
)

type AdventDay struct {
	InputFile string
	Year int
	Day int
	DayRunner Runner
}

type Runner interface {
	Flags() *flag.FlagSet
	Run(input []string) any
	TestInput() []string
}

// global
var runnerMap = map[int]map[int]func() Runner {
	2024: {
		11: nil,
	},
}


func NewDay(year int, day int) (ad *AdventDay, err error) {
	y := runnerMap[year]
	if y == nil {
		return nil, fmt.Errorf("invalid year %d", year)
	}
	r := y[day]
	if r == nil {
		return nil, fmt.Errorf("invalid day %d for year %d", day, year)
	}

	d := AdventDay{
		Year: year,
		Day: day,
		InputFile: inputFile(year, day),
		DayRunner: r(),
	}
	return &d, nil
}


func RunnerMap(year int, day int, runner func() Runner) *map[int]map[int]func() Runner {
	r := runnerMap[year][day]
	if r == nil {
		runnerMap[year][day] = runner
	}
	return &runnerMap
}

func inputFile(year int, day int) string {
	return fmt.Sprintf("year%d/inputDay%02d.txt", year, day)
}

