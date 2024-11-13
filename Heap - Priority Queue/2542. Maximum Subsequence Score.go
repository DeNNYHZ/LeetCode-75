package Heap___Priority_Queue

import (
	"container/heap"
	"sort"
)

// A MinHeap to store the minimum values.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	// Create pairs of (nums1[i], nums2[i]).
	pairs := make([][2]int, len(nums1))
	for i := range nums1 {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}

	// Sort pairs by the second element (nums2[i]) in descending order.
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	n1Sum := 0
	res := int64(0)

	for _, pair := range pairs {
		n1, n2 := pair[0], pair[1]
		heap.Push(minHeap, n1)
		n1Sum += n1

		// If the size of the heap exceeds k, remove the smallest element.
		if minHeap.Len() > k {
			n1Sum -= heap.Pop(minHeap).(int)
		}

		// If the heap size is exactly k, update the result.
		if minHeap.Len() == k {
			res = maxInt64(res, int64(n1Sum)*int64(n2))
		}
	}

	return res
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
