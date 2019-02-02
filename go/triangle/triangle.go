// Observations on Mingan's solution
package triangle

import "math"

// iota constant
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// defines Kind as type int
type Kind int

func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}
	if a+b < c || b+c < a || a+c < b {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
