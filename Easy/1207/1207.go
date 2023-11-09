package Easy

func uniqueOccurrences(arr []int) bool {

	mapR := make(map[int]int)

	//_ is the index, v is the value
	for _, v := range arr {
		mapR[v]++
	}

	for outer_key, outer_value := range mapR {
		for inner_key, inner_value := range mapR {
			if outer_key != inner_key && outer_value == inner_value {
				return false
			}
		}
	}

	return true
}
