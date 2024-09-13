package Array_String

import (
	"fmt"
)

// 1768. Merge Strings Alternately
// You are given two strings word1 and word2. Merge the strings by alternating the letters in both strings.
// If one string is longer than the other, append the remaining letters in the longer string.
// Return the merged string.
func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	result := ""

	for i < len(word1) && j < len(word2) {
		result += string(word1[i])
		result += string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		result += string(word1[i])
		i++
	}

	for j < len(word2) {
		result += string(word2[j])
		j++
	}

	return result
}

func main() {
	word1 := "abc"
	word2 := "pqr"
	merged := mergeAlternately(word1, word2)
	fmt.Println(merged) // Output: "apbqcrstu"
}
