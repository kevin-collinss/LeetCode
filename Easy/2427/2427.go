package Easy

func commonFactors(a int, b int) int {

	var top int
	count := 0

	if a < b {
		top = a
	} else {
		top = b
	}

	for top > 0 {
		if a%top == 0 && b%top == 0 {
			count++
		}
		top--
	}
	return count
}
