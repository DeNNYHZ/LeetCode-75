package Graphs___DFS

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	var dfs func(room int)
	dfs = func(room int) {
		visited[room] = true
		for _, key := range rooms[room] {
			if !visited[key] {
				dfs(key)
			}
		}
	}
	dfs(0)
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
