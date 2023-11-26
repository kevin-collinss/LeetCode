package Easy

func subtractProductAndSum(n int) int {

	var arr []int = make([]int, 0)
	product := 1
	sum := 0

	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
	}

	for i := 0; i < len(arr); i++ {
		product *= arr[i]
		sum += arr[i]
	}

	return product - sum
}
