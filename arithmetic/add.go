// Package arithmetic provides functions to make do arithmetic
package arithmetic

import "golang.org/x/exp/constraints"

// Number is a constraint for math operations
type Number interface {
	constraints.Integer | constraints.Float
}

// Add return the sum of a and b
func Add[N Number](a, b N) N {
	return a + b
}
