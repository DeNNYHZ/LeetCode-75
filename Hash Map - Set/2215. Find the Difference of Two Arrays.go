package Hash_Map___Set

func findDifference(nums1 []int, nums2 []int) [][]int {
	numMap1 := make(map[int]bool)
	numMap2 := make(map[int]bool)

	for _, num := range nums1 {
		numMap1[num] = true
	}
	for _, num := range nums2 {
		numMap2[num] = true
	}

	var res1, res2 []int
	for num := range numMap1 {
		if !numMap2[num] {
			res1 = append(res1, num)
		}
	}
	for num := range numMap2 {
		if !numMap1[num] {
			res2 = append(res2, num)
		}
	}

	return [][]int{res1, res2}
}
