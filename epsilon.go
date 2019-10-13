package equal

import "math"

func CompareEpsilon(epsilon float64, v1 float64, v2 float64) bool {
	if epsilon < 0 {
		epsilon = -epsilon
	}
	return math.Abs(v1-v2) <= epsilon
}
