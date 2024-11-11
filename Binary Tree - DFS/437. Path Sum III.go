package Binary_Tree___DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	count := 0
	if root.Val == sum {
		count++
	}

	count += dfs(root.Left, sum-root.Val)
	count += dfs(root.Right, sum-root.Val)

	return count
}
