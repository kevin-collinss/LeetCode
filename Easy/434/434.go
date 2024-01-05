package Easy

func countSegments(s string) int {

	runes := []rune(s)
	count := 0

	in_a_segment := false

	for i := 0; i < len(runes); i++ {

		if runes[i] == ' ' && in_a_segment == true {
			count++
			in_a_segment = false
		}

		if runes[i] != ' ' {
			in_a_segment = true
		}
	}

	if in_a_segment == true {
		count++
	}

	return count
}
