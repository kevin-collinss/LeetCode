package Easy

func removeElement(nums []int, val int) int {

	last := len(nums) - 1
	i := 0

	for i <= last {
		if nums[i] == val {
			nums[i] = nums[last]
			nums[last] = 0
			last--
		} else {
			i++
		}
	}
	return i
}
