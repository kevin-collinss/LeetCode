package easy

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	result := 0
	for num > 0 {
        //bitwise & operator
		if num&1 == 1 {
			result += 2
		} else {
			result += 1
		}
        
        //bitwise shift
		num >>= 1
	}

	return result - 1
}