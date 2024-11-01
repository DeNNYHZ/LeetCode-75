package Backtracking

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var backtrack func(int, int, int, []int)
	backtrack = func(start, target, k int, path []int) {
		if target == 0 && k == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if target < 0 || k < 0 {
			return
		}
		for i := start; i < len(nums); i++ {
			backtrack(i+1, target-nums[i], k-1, append(path, nums[i]))
		}
	}
	backtrack(0, n, k, []int{})
	return res
}
