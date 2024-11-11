package Binary_Tree___DFS

func longestZigZag(root *TreeNode) int {
	maxLength := 0

	var dfs func(node *TreeNode, leftLength, rightLength int)
	dfs = func(node *TreeNode, leftLength, rightLength int) {
		if node == nil {
			return
		}

		if leftLength > maxLength {
			maxLength = leftLength
		}
		if rightLength > maxLength {
			maxLength = rightLength
		}

		dfs(node.Left, rightLength+1, 0)

		dfs(node.Right, 0, leftLength+1)
	}

	dfs(root, 0, 0)

	return maxLength
}
