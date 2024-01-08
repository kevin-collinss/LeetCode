package Easy

func findMaxConsecutiveOnes(nums []int) int {

	max := 0
	current_count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current_count++

			if current_count > max {
				max = current_count
			}
		} else {
			current_count = 0
		}
	}

	return max
}
