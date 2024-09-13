package Array_String

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, value := range candies {
		if value > max {
			max = value
		}
	}

	result := make([]bool, len(candies))
	for i, value := range candies {
		if value+extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}
