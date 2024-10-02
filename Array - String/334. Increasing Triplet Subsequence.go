package Array___String

import "math"

func increasingTriplet(nums []int) bool {
	small, big := math.MaxInt64, math.MaxInt64
	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= big {
			big = num
		} else {
			return true
		}
	}
	return false
}
