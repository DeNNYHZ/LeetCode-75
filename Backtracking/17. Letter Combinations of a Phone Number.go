package Backtracking

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res []string
	var dfs func(int, string)
	m := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	dfs = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		for _, c := range m[digits[index]] {
			dfs(index+1, path+string(c))
		}
	}
	dfs(0, "")
	return res
}
