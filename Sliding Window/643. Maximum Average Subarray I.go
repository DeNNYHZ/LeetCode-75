package Sliding_Window

func findMaxAverage(nums []int, k int) float64 {
	cur_sum := 0
	for i := 0; i < k; i++ {
		cur_sum += nums[i]
	}

	max_sum := cur_sum

	for i := k; i < len(nums); i++ {
		cur_sum += nums[i] - nums[i-k]
		if cur_sum > max_sum {
			max_sum = cur_sum
		}
	}
	return float64(max_sum) / float64(k)
}
