package mathutils

import (
	"advent/utils/collections"

	"golang.org/x/exp/constraints"
)

type Number interface {
    constraints.Integer | constraints.Float
}

func Sum[S ~[]E, E Number] (source S) E {
	sum := func(v1 E, v2 E) E {
		return v1 + v2
	}
	return collections.Reduce(source, sum, 0)
}
