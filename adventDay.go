package main

import (
	"advent/types"
	"advent/utils/debug"
	"advent/year2024"
	"fmt"
)

var adventRunners = map[int]map[int]func(runArgs []string)types.Runner {
	2024: {
		11: year2024.NewDay11,
	},
}

// ADD RUNARGS (?)
func NewDay(year int, day int, runArgs []string) types.Runner {
	debug.DebugPrintln("DAY ARGS %s", runArgs)

	d := adventRunners[year][day](runArgs)
	d.Config(configFile(year, day), inputFile(year, day))
	return d
}

func configFile(year int, day int) string {
	return fmt.Sprintf("year%d/day%02d.toml", year, day)
}

func inputFile(year int, day int) string {
	return fmt.Sprintf("year%d/inputDay%02d.txt", year, day)
}