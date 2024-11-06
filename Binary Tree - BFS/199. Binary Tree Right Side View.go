package Binary_Tree___BFS

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := []int{}

	// Loop untuk setiap level pada pohon
	for len(queue) > 0 {
		size := len(queue)

		// Ambil nilai node terakhir dari setiap level
		result = append(result, queue[size-1].Val)

		// Proses setiap node di level saat ini
		for i := 0; i < size; i++ {
			node := queue[i]

			// Tambahkan anak kiri dan kanan node ke dalam queue untuk level berikutnya
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Hapus node yang sudah diproses dari queue
		queue = queue[size:]
	}

	return result
}
