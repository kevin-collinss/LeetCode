package Easy

func numIdenticalPairs(nums []int) int {

	count := 0
	for i, val := range nums {
		for j, valInner := range nums {
			if val == valInner && j > i {
				count++
			}
		}
	}
	return count
}
