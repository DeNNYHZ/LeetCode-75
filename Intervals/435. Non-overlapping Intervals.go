package Intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 1
	end := intervals[0][1]

	for _, interval := range intervals[1:] {
		if interval[0] >= end {
			count++
			end = interval[1]
		} else if interval[1] < end {
			end = min(end, interval[1])
		}
	}

	return len(intervals) - count
}
