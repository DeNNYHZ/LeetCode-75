package Hash_Map___Set

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	rowMap := make(map[string]int)

	// Convert each row into a string and store in a map
	for i := 0; i < n; i++ {
		rowStr := fmt.Sprint(grid[i])
		rowMap[rowStr]++
	}

	// Initialize result count
	count := 0

	// For each column, create its string representation and check in rowMap
	for j := 0; j < n; j++ {
		var col []int
		for i := 0; i < n; i++ {
			col = append(col, grid[i][j])
		}
		colStr := fmt.Sprint(col)
		count += rowMap[colStr]
	}

	return count
}
