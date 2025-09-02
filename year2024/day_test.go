package year2024

import (
	"testing"
)


func TestDay11(t *testing.T) {
	var SMALL = []string{
		"125 17",
	}

	d := AdventDay11{numBlinks: 1}

	var want uint = 7
	res := d.Run(d.TestInput())
	n := res.(uint)
	if n != want {
		t.Errorf("run result %d does not match expected %d", res, want)
	}

	d.numBlinks = 6
	want = 22
	res = d.Run(SMALL)
	n = res.(uint)
	if n != want {
		t.Errorf("run result %d does not match expected %d", res, want)
	}

	d.numBlinks = 25
	want = 55312
	res = d.Run(SMALL)
	n = res.(uint)
	if n != want {
		t.Errorf("run result %d does not match expected %d", res, want)
	}


}