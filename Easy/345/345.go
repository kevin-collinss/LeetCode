package Easy

func reverseVowels(s string) string {

	left, right := 0, len(s)-1
	runes := []rune(s)

	for left < right {

		if !is_a_vowel(runes[right]) {
			right--
			continue
		}
		if !is_a_vowel(runes[left]) {
			left++
			continue
		}

		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func is_a_vowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
