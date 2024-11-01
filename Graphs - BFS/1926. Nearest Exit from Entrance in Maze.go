package Graphs___BFS

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{entrance[0], entrance[1]})
	maze[entrance[0]][entrance[1]] = '+'
	steps := 0
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			x, y := node[0], node[1]
			if (x == 0 || x == m-1 || y == 0 || y == n-1) && (x != entrance[0] || y != entrance[1]) {
				return steps
			}
			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && maze[newX][newY] == '.' {
					maze[newX][newY] = '+'
					queue = append(queue, [2]int{newX, newY})
				}
			}
		}
		steps++
	}
	return -1
}
