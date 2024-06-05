import "math"

func minEatingSpeed(piles []int, h int) int {
	low := 1
	high := max(piles)

	for low < high {
		k := low + (high-low)/2
		hours := 0

		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}

		if hours > h {
			low = k + 1
		} else {
			high = k
		}
	}
	return high
}

func max(piles []int) int {
	max := 1
	for _, val := range piles {
		if val > max {
			max = val
		}
	}
	return max
}