package Easy

import "slices"

func destCity(paths [][]string) string {
	var front_array []string
	for i := range paths {
		front_array = append(front_array, paths[i][0])
	}

	for j := range paths {
		if !slices.Contains(front_array, paths[j][1]) {
			return paths[j][1]
		}
	}
	return "None Found"
}
