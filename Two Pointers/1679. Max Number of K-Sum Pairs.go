package Two_Pointers

func maxOperations(nums []int, k int) int {
	count := 0
	numMap := make(map[int]int)

	for _, num := range nums {
		complement := k - num
		if numMap[complement] > 0 {
			count++
			numMap[complement]--
		} else {
			numMap[num]++
		}
	}

	return count
}
