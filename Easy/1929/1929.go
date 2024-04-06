package Easy

func getConcatenation(nums []int) []int {

	ansLen := len(nums) * 2

	ans := make([]int, ansLen)
	posi := 0

	for i := 1; i < 3; i++ {
		for _, j := range nums {
			ans[posi] = j
			posi++
		}
	}

	return ans
}

// func getConcatenation(nums []int) []int {

// 	ansLen := len(nums) * 2

// 	ans := make([]int, ansLen)
// 	posi := 0

// 	for i := 1; i < 3; i++ {
// 		for _, j := range nums {
// 			ans[posi] = j
// 			posi++
// 		}
// 	}

// 	return ans
// }
