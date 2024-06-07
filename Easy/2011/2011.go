package Easy

func finalValueAfterOperations(operations []string) int {

	X := 0
	for _, op := range operations {

		if op == "++X" || op == "X++" {
			X++
		} else {
			X--
		}
	}
	return X
}
