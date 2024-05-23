package Easy

import (
	"math"
)

func scoreOfString(s string) int {
	total := 0
	asciiValues := []int{}

	for i, c := range s {
		asciiValues = append(asciiValues, int(c))
		if i > 0 {
			total += int(math.Abs(float64(asciiValues[i-1] - asciiValues[i])))
		}
	}

	return total
}
