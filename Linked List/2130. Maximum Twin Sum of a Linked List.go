package Linked_List

func pairSum(head *ListNode) int {
	var values []int
	for current := head; current != nil; current = current.Next {
		values = append(values, current.Val)
	}

	maxSum := 0
	n := len(values)
	for i := 0; i < n/2; i++ {
		pairSum := values[i] + values[n-i-1]
		if pairSum > maxSum {
			maxSum = pairSum
		}
	}
	return maxSum
}
