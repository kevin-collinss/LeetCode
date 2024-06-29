package Easy

func isAnagram(s string, t string) bool {
	if len(s) == len(t) {
		sMap := make(map[rune]int)
		for _, val := range s {
			sMap[val]++
		}
		for _, val := range t {
			sMap[val]--
		}
		for _, val := range sMap {
			if val != 0 {
				return false
			}
		}
		return true
	}
	return false
}
