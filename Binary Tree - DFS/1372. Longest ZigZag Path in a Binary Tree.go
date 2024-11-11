package Binary_Tree___DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	root = &TreeNode{Left: root}
	longest := longestZigZag(root)
	longest = max(longest, longestZigZag(root.Right))
	return longest

}
