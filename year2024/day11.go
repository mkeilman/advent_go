package year2024

import (
	"advent/types"
	"advent/utils/debug"
	"advent/utils/collections"
	"advent/utils/mathutils"

	//"flag"
	"maps"
	"math"
	"regexp"
	"strconv"
)

var TEST = []string{
	"0 1 10 99 999",
}

	
type AdventDay11 struct {
	data *types.AdventData
	numBlinks int
}

func NewDay11(runArgs []string) any {
	return AdventDay11{}
}

func (d AdventDay11) Config(cfgFileName string, inputFileName string) {
	debug.DebugPrintln("CONFIG %s INPUT %s", cfgFileName, inputFileName)
	//*d.data.InputFile = inputFileName
}

func (d AdventDay11) InputFile() string {
	return *d.data.InputFile
}

func (d AdventDay11) TestInput() []string {
	return TEST
}

func (d AdventDay11) Run(input []string) int {
	*d.data.Input = input

	digits := regexp.MustCompile(`\d+`)

	// single line
	stones :=  collections.Map(digits.FindAllString((*d.data.Input)[0], -1), func (s string) uint {res, _ := strconv.ParseInt(s, 10, 0); return uint(res) })
	s := d.blink(stones)
	n := d.numStones(s)
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
func (d AdventDay11) blink(stones []uint) map[uint]uint {

	var countMap map[uint]uint

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

	if d.numBlinks < 1 {
		return countMap
	}
	
	for i := range d.numBlinks {
		if i == i {}  // stupid
		var m map[uint]uint
		for _, s := range countMap {
			updateCountMap(d.nextStones(s), m, countMap[s])
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
func (d AdventDay11) nextStones(stone uint) []uint {
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


func (d AdventDay11) numStones(countMap map[uint]uint) uint {
	v := []uint{}
	for i := range(maps.Values(countMap)) {
		v = append(v, i)
	}
	return mathutils.Sum(v)
}
