package Binary_Search_Tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findNode returns a tuple of the found node and its parent
func findNode(t *TreeNode, key int) (*TreeNode, *TreeNode) {
	var parent *TreeNode
	cur := t
	for cur != nil {
		if cur.Val == key {
			break
		}
		parent = cur
		if cur.Val < key {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	return cur, parent
}

// adjustTree handles the deletion and adjustment of subtrees
func adjustTree(t *TreeNode) *TreeNode {
	if t.Left == nil {
		return t.Right
	}
	if t.Right == nil {
		return t.Left
	}

	// Find rightmost node in the left subtree
	cur := t.Left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = t.Right

	return t.Left
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	cur, par := findNode(root, key)
	if cur == nil {
		return root
	}

	newChild := adjustTree(cur)

	if par == nil { // Node to be deleted is the root
		return newChild
	}

	if par.Left == cur {
		par.Left = newChild
	} else {
		par.Right = newChild
	}

	return root
}
