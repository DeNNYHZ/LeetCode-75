package Binary_Search

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1
	for _, p := range piles {
		right = max(right, p)
	}
	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(piles []int, h int, mid int) bool {
	count := 0
	for _, p := range piles {
		count += (p + mid - 1) / mid
	}
	return count <= h
}
