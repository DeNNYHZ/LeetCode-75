package Array___String

import "strconv"

func compress(chars []byte) int {
	n := len(chars)
	i, j := 0, 0
	for j < n {
		count := 1
		for j+1 < n && chars[j] == chars[j+1] {
			count++
			j++
		}
		chars[i] = chars[j]
		i++
		if count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars[i] = byte(digit)
				i++
			}
		}
		j++
	}
	return i
}
