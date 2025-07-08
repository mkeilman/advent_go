package collections

import (
	"fmt"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
    source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    want := []int{2, 4, 6, 8, 10}
	fail := []int{1, 3, 5, 7, 9}
	isEven := func (i int) bool {
		return i % 2 == 0
	}

	res := Filter(source, isEven)
    if !slices.Equal(res, want) {
        t.Errorf("Filter result %v does not match expected %v", res, want)
    }
	if slices.Equal(res, fail) {
        t.Errorf("Filter result %v improperly matches %v", res, fail)
    }
}

func TestMap(t *testing.T) {
	type P struct {
		first string
		last string
	}

	firstLast := func (p P) string {
		return fmt.Sprintf("%s %s", p.first, p.last)
	}

	numChars := func (p P) int {
		return len(p.first) + len(p.last)
	}

    source := []P {
		{"Alice", "Jones"},
		{"Bob", "Smith"},
	}

    want := []string {
		"Alice Jones",
		"Bob Smith",
	}

	intWant := []int {
		10,
		8,
	}

	res := Map(source, firstLast)
    if !slices.Equal(res, want) {
        t.Errorf("Map result %v does not match expected %v", res, want)
    }

	intRes := Map(source, numChars)
	if !slices.Equal(intRes, intWant) {
        t.Errorf("Map result %v does not match expected %v", intRes, intWant)
    }
}

func TestReduce(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := func(v1 int, v2 int) int {
		return v1 + v2
	}
	want := 55
	res := Reduce(source, sum, 0)

	if res != want {
		t.Errorf("Reduce result %d does not match expected %d", res, want)
	}
}
