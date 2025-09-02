package year2024

import (
	"advent/run"
	"advent/types"

	"fmt"
	"testing"
)


func TestDay11(t *testing.T) {
	var SMALL = []string{
		"125 17",
	}

	d := getDay(11)
	r := d.DayRunner
	d11 := r.(*AdventDay11)

	d11.numBlinks = 1
	doTest(t, uint(7), run.RunFromTestInput(*d))

	d11.numBlinks = 6
	doTest(t, uint(22), r.Run(SMALL))

	d11.numBlinks = 25
	doTest(t, uint(55312), r.Run(SMALL))

	d11.numBlinks = 25
	doTest(t, uint(203457), run.RunFromFile(*d))

	d11.numBlinks = 75
	doTest(t, uint(241394363462435), run.RunFromFile(*d))
}

func doTest(t *testing.T, want any, res any) {
	if res != want {
		t.Errorf("run result %d does not match expected %d", res, want)
	}
}

func getDay(day int) *types.AdventDay {
	d, err := types.NewDay(2024, day)
	if err != nil {
		run.Exit(run.ERR_INVALID_DAY, fmt.Sprintf("invalid advent day %d", day))
	}
	return d
}
