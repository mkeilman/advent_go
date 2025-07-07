package collections

import (
	//"slices"
	//"advent/utils/debug"
)

func Filter[S ~[]E, E any] (source S, filter func (E) bool) S {
	res := []E{}
	for _, e := range source {
		if filter(e) {
			res = append(res, e)
		}
	}
	return res
}

