package Sliding_Window

func longestOnes(nums []int, k int) int {
	left, right, zeroCount, max := 0, 0, 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		max = maxInt(max, right-left+1)
		right++
	}
	return max
}
