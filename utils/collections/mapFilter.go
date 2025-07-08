package collections

import (
	//"slices"
	//"advent/utils/debug"
)

// Filter a slice according to the given criteria
//
// Args:
//     source ([any]): the slice to filter
//     doesPass (func (any) bool): when true, add this element
func Filter[S ~[]E, E any] (source S, doesPass func (E) bool) S {
	res := []E{}
	for _, e := range source {
		if doesPass(e) {
			res = append(res, e)
		}
	}
	return res
}

// Map a slice's values to other values
//
// Args:
//     source ([any]): the slice to map
//     mapping (func (any) any): map to a new value
func Map[S ~[]E, E any, V any] (source S, mapping func (E) V) []V {
	res := []V{}
	for _, e := range source {
		res = append(res, mapping(e))
	}
	return res
}

// Reduce a slice to a single value
//
// Args:
//     source ([any]): the slice to map
//     reducer (func (any, any) any): a function that combines a value with
//         a value of the same type derived from an element of the slice
//     initVal (any): the initial value
func Reduce[S ~[]E, E any, V any] (source S, reducer func (V, E) V, initVal V) V {
	var res = initVal
	for _, e := range source {
		res = reducer(res, e)
	}
	return res
}
