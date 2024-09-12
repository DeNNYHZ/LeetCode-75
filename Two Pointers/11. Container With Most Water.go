package Two_Pointers

func maxArea(height []int) int {

	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {

		width := right - left
		minHeight := min(height[left], height[right])
		area := width * minHeight
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
