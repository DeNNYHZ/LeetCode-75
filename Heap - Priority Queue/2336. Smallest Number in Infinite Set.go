package Heap___Priority_Queue

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
	minSize int
	heap    IntHeap
	set     map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		minSize: 1,
		heap:    *h,
		set:     make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.heap.Len() > 0 {
		smallest := heap.Pop(&this.heap).(int)
		delete(this.set, smallest)
		return smallest
	}

	smallest := this.minSize
	this.minSize++
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.minSize && !this.set[num] {
		heap.Push(&this.heap, num)
		this.set[num] = true
	}
}
