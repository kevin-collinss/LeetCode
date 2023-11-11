package Easy

func canThreePartsEqualSum(arr []int) bool {

	sumTotal := 0
	runningTotal := 0
	for _, i := range arr {
		sumTotal += i
	}

	if sumTotal%3 != 0 {
		return false
	}

	thirdsVal := sumTotal / 3
	count := 0

	for j := range arr {
		runningTotal += arr[j]
		if runningTotal == thirdsVal {
			runningTotal = 0
			count++
		}
	}

	return count >= 3
}
