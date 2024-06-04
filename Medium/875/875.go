package Medium

import "math"

func minEatingSpeed(piles []int, h int) int {

	high := 0
	low := 1

	for _, val := range piles {
		if val > high {
			high = val
		}
	}

	for low < high {
		//set variables
		k := low + (high-low)/2
		hours := 0

		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}

		if hours <= h {
			high = k
		} else {
			low = k + 1
		}

	}
	return high

}
