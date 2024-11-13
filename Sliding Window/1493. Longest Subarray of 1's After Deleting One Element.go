package Sliding_Window

func longestSubarray(nums []int) int {
	left, right, zeroCount, max := 0, 0, 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		max = maxInt(max, right-left)
		right++
	}
	return max
}
