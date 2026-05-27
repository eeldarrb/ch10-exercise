// Package math provides simple math functions
package math

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes the sum of two integers and returns the result
func Add[T Number](a, b T) T {
	return a + b
}
