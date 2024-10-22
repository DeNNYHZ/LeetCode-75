package Binary_Tree___DFS

import "reflect"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(*TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		if root.Left == nil && root.Right == nil {
			return []int{root.Val}
		}
		return append(dfs(root.Left), dfs(root.Right)...)
	}
	return reflect.DeepEqual(dfs(root1), dfs(root2))
}
