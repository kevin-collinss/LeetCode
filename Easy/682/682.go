package Easy

import "strconv"

func calPoints(operations []string) int {

	scores := []int{}

	for _, val := range operations {
		j := len(scores) - 1
		if num, err := strconv.Atoi(val); err == nil {
			scores = append(scores, num)
		}
		if val == "+" {
			scores = append(scores, (scores[j] + scores[j-1]))
		}
		if val == "D" {
			scores = append(scores, (scores[j] * 2))
		}
		if val == "C" {
			scores = scores[:len(scores)-1]
		}
	}

	sum := 0
	for _, score := range scores {
		sum += score
	}

	return sum
}
