package Binary_Search

func guessNumber(n int) int {
	l, r := 1, n

	for l <= r {
		m := (l + r) / 2
		res := guess(m)

		if res > 0 {
			l = m + 1
		} else if res < 0 {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}

func guess(num int) int {
	return 0
}
