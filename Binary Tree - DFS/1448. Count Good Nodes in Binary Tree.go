package Binary_Tree___DFS

func goodNodes(root *TreeNode) int {

	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, max int) int {
		if root == nil {
			return 0
		}
		count := 0
		if root.Val >= max {
			count++
			max = root.Val
		}
		count += dfs(root.Left, max)
		count += dfs(root.Right, max)
		return count
	}
	return dfs(root, root.Val)
}
