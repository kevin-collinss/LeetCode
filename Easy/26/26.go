package Easy

func removeDuplicates(nums []int) int {
	count := 0

	for i := range nums {
		if nums[i] != nums[count] {
			count += 1
			nums[count] = nums[i]
		}
	}

	return count + 1
}
