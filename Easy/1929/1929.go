package Easy

func getConcatenation(nums []int) []int {

	var ans []int

	for i := 1; i < 3; i++ {
		for _, j := range nums {
			ans = append(ans, j)
		}
	}
	return ans
}
