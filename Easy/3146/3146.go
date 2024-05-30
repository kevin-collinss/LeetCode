package Easy

import "math"

func findPermutationDifference(s string, t string) int {

	total := 0
	for i, firstChar := range s {
		for j, secondChar := range t {
			if firstChar == secondChar {
				total += int(math.Abs(float64(i - j)))
			}
		}
	}
	return total
}
