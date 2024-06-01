package Easy

func defangIPaddr(address string) string {
	defangAddr := []rune{}
	for _, char := range address {
		if char == '.' {
			defangAddr = append(defangAddr, '[', '.', ']')
		} else {
			defangAddr = append(defangAddr, char)
		}
	}

	return string(defangAddr)
}
