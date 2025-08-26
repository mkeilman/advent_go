package year2024

import (
	"advent/types"
	"advent/utils/debug"
	"advent/utils/collections"
	"advent/utils/mathutils"

	"flag"
	"maps"
	"math"
	"regexp"
	"strconv"
)

var TEST = []string{
	"0 1 10 99 999",
}

	
type AdventDay11 struct {
	numBlinks uint
}

func NewDay11(runArgs []string) types.Runner {
	d := &AdventDay11{}
	return d
}

func (d *AdventDay11) Flags() *flag.FlagSet{
	fs := flag.NewFlagSet("2024.11", flag.ExitOnError)
	fs.UintVar(&d.numBlinks, "num-blinks", 25, "--num-blinks=<num_blinks> (num_blinks >= 0)")
	return fs
}

func (d *AdventDay11) TestInput() []string {
	return TEST
}

func (d *AdventDay11) Run(input []string) int {
	digits := regexp.MustCompile(`\d+`)

	// input is on a single line
	stones :=  collections.Map(
		digits.FindAllString(input[0], -1),
		func (s string) uint {res, _ := strconv.ParseInt(s, 10, 0); return uint(res) },
	)
	n := numStones(blink(d.numBlinks, stones))
	debug.DebugPrintln("stones %v, %d blinks -> %d total", stones, d.numBlinks, n)
	return int(n)
}

// Generate new stones from the given array. Luckily the goal of this exercise is to *count* the
// resultant number of stones, *not* to deliver the entire assortment
//
//     Args:
//         stones (int[]): array of stone values
//
//     Returns:
//         (dict) a map of stone values to the number of stones with that value
//
func blink(numBlinks uint, stones []uint) map[uint]uint {

	var countMap map[uint]uint = make(map[uint]uint)

	updateCountMap := func(st []uint, m map[uint]uint, numStones uint) {
		for _, s := range st {
			_, ok := m[s]
			if !ok {
				m[s] = 0
			}
			m[s] += numStones
		}
	}

	updateCountMap(stones, countMap, 1)

	if numBlinks < 1 {
		return countMap
	}
	
	for i := range numBlinks {
		// we must refer to i even if we don't need it
		debug.DebugIf("", "", i < 0)
		var m map[uint]uint = make(map[uint]uint)
		for k, v := range countMap {
			updateCountMap(nextStones(k), m, v)
		}
		countMap = m
	}
	
	return countMap
}

// Generates new stones accordong to the above rules
//
//    Args:
//        stone (int): the value of the current stone
//
//    Returns:
//        ([]int) an array of generated stones
func nextStones(stone uint) []uint {
	numDigits := func(n uint) int {
		return int(math.Log10(float64(n))) + 1
	}

	split := func(s uint) []uint {
		f := math.Pow10(numDigits(s) / 2)
		return []uint{s / uint(f), s % uint(f)}
	}
        
	if stone == 0 {
		return []uint{1}
	}
	if numDigits(stone) % 2 == 0 {
		return split(stone)
	}
	return []uint{stone * 2024}
}


func numStones(countMap map[uint]uint) uint {
	v := []uint{}
	for i := range(maps.Values(countMap)) {
		v = append(v, i)
	}
	return mathutils.Sum(v)
}
