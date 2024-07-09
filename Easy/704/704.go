package Easy

func search(nums []int, target int) int {

	low := 0
	high := len(nums)

	for low < high {
		med := low + (high-low)/2
		val := nums[med]

		if val == target {
			return med
		} else if val > target {
			high = med
		} else {
			low = med + 1
		}
	}
	return -1
}
