package libs

import "math"

func MultiReturn(a, b float64) (max float64, result string) {
	if a == b {
		max = a
		result = "same"
		return
	}
	max = math.Max(a, b)
	result = "max value"
	return
}
