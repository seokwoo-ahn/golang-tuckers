package libs

import "math"

func EqualBit(a, b float64) bool {
	return math.Nextafter(a, b) == b
}
