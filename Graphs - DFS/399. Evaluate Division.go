package Graphs___DFS

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Step 1: Build the graph
	graph := make(map[string]map[string]float64)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1.0 / values[i]
	}

	// Step 2: Define DFS function
	var dfs func(curr, target string, visited map[string]bool) float64
	dfs = func(curr, target string, visited map[string]bool) float64 {
		if _, exists := graph[curr]; !exists {
			return -1.0
		}
		if val, exists := graph[curr][target]; exists {
			return val
		}
		visited[curr] = true
		for neighbor, value := range graph[curr] {
			if !visited[neighbor] {
				product := dfs(neighbor, target, visited)
				if product != -1.0 {
					return value * product
				}
			}
		}
		return -1.0
	}

	// Step 3: Process queries
	results := make([]float64, len(queries))
	for i, query := range queries {
		a, b := query[0], query[1]
		if a == b {
			if _, exists := graph[a]; exists {
				results[i] = 1.0
			} else {
				results[i] = -1.0
			}
		} else {
			visited := make(map[string]bool)
			results[i] = dfs(a, b, visited)
		}
	}

	return results
}
