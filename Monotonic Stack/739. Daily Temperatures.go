package Monotonic_Stack

func dailyTemperatures(temperatures []int) []int {
	temperatures = append(temperatures, 300)
	stack := make([]int, 0)
	result := make([]int, len(temperatures)-1)
	for i := len(temperatures) - 2; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return result
}
