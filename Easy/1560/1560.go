package Easy

func mostVisited(n int, rounds []int) []int {

	map_count := make(map[int]int)

	//for every sector, create a key in the map and set the val to 0
	for sector := 1; sector <= n; sector++ {
		map_count[sector] = 0
	}

	//position is init to first round[] value
	position := rounds[0]

	//for every val in rounds array
	for _, v := range rounds {

		// position isn't at the required next position
		for position != v {

			//map_count and the position is increased by one as it passes
			map_count[position]++

			//move on position
			position++

			//if the position is higher than sectors set it back to 1
			if position > n {
				position = 1
			}
		}
	}

	//most visits at any sector
	most_visits := 0

	//for every sector, if value at i > most visits than that is the new most visits to any place
	for i := 1; i <= n; i++ {
		if map_count[i] > most_visits {
			most_visits = map_count[i]
		}
	}

	//most visited sector array RETURN ARRAY
	most_visited_sectors := []int{}

	//for every sector if the value at the sector = the most visits than append the sector number onto the array
	for i := 1; i <= n; i++ {
		if map_count[i] == most_visits {
			most_visited_sectors = append(most_visited_sectors, i)
		}
	}

	return most_visited_sectors
}
