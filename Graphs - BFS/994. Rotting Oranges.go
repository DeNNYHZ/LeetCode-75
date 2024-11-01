package Graphs___BFS

func orangesRotting(grid [][]int) int {
	minutes := 0
	queue := make([][2]int, 0)
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	freshOranges := 0

	// Initialize the queue with all initially rotten oranges and count fresh oranges.
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	// Early return if there are no fresh oranges.
	if freshOranges == 0 {
		return 0
	}

	// Process the BFS queue to rot adjacent oranges minute by minute.
	for len(queue) > 0 {
		size := len(queue)
		minutePassed := false
		for i := 0; i < size; i++ {
			x, y := queue[i][0], queue[i][1]
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					queue = append(queue, [2]int{nx, ny})
					freshOranges--
					minutePassed = true
				}
			}
		}
		// Increment minutes only if at least one orange was rotted this minute.
		if minutePassed {
			minutes++
		}
		queue = queue[size:]
	}

	// If there are still fresh oranges left, return -1; otherwise, return the total minutes.
	if freshOranges > 0 {
		return -1
	}
	return minutes
}
