package Medium

import (
	"fmt"
	"strconv"
)

func searchMatrix(matrix [][]int, target int) bool {

	highY := len(matrix) - 1
	lowY := 0
	medY := 0

	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0] == target
	}

	for lowY <= highY {

		fmt.Println("LowY: " + strconv.Itoa(lowY))
		fmt.Println("HighY: " + strconv.Itoa(highY))
		fmt.Println("Y: " + strconv.Itoa(medY))
		medY = lowY + (highY-lowY)/2
		current := matrix[medY][0]
		fmt.Println("CurrentY: " + strconv.Itoa(current))

		if current == target {
			return true
		} else if current > target {
			highY = medY - 1
		} else {
			lowY = medY + 1
		}
	}

	if lowY > 0 {
		lowY--
	}

	highX := len(matrix[0]) - 1
	lowX := 0

	for lowX <= highX {

		medX := lowX + (highX-lowX)/2
		fmt.Println("X: " + strconv.Itoa(medX))
		current := matrix[lowY][medX]
		fmt.Println("CurrentX: " + strconv.Itoa(current))
		if current == target {
			return true
		} else if current > target {
			highX = medX - 1
		} else {
			lowX = medX + 1
		}
	}

	return false
}
