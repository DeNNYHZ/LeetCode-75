package Graphs___DFS

func minReorder(n int, connections [][]int) int {
	graph := make(map[int][][2]int)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], [2]int{conn[1], 1}) // 1 means original direction from conn[0] to conn[1]
		graph[conn[1]] = append(graph[conn[1]], [2]int{conn[0], 0}) // 0 means reverse direction (from conn[1] to conn[0])
	}

	visited := make([]bool, n)
	queue := []int{0}
	reorderCount := 0

	// Perform BFS traversal
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true

		for _, neighbor := range graph[current] {
			if !visited[neighbor[0]] {
				// If the road direction is 1 (outgoing from the current node), we need to reorder
				if neighbor[1] == 1 {
					reorderCount++
				}
				queue = append(queue, neighbor[0])
			}
		}
	}

	return reorderCount
}
