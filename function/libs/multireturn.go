package libs

import "math"

func MultiReturn(a, b float64) (float64, string) {
	if a == b {
		return a, "same"
	}
	return math.Max(a, b), "max value"
}
