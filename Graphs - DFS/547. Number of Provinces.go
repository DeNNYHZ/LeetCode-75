package Graphs___DFS

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	count := 0
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			dfs(isConnected, visited, i)
			count++
		}
	}
	return count
}

func dfs(connected [][]int, visited []bool, i int) {
	visited[i] = true
	for j := 0; j < len(connected); j++ {
		if connected[i][j] == 1 && !visited[j] {
			dfs(connected, visited, j)
		}
	}
}
