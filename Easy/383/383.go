package Easy

func canConstruct(ransomNote string, magazine string) bool {

	mapR := make(map[rune]int)

	for _, v := range ransomNote {
		mapR[v]++
	}

	for _, v := range magazine {
		if mapR[v] > 0 {
			mapR[v]--
		}
	}

	for _, v := range mapR {
		if v > 0 {
			return false
		}
	}

	return true
}
