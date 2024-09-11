package main

import (
	"fmt"
	"strings"
)

// gcdOfStrings calculates the greatest common divisor of two strings.
// The GCD of two strings is the longest string that can be concatenated to form both strings.
func gcdOfStrings(str1 string, str2 string) string {
	// Swap str1 and str2 if str2 is longer to simplify the comparison
	if len(str2) > len(str1) {
		return gcdOfStrings(str2, str1)
	}

	// If str1 and str2 are equal, they are the GCD
	if str1 == str2 {
		return str1
	}

	// Check if str1 starts with str2
	if strings.HasPrefix(str1, str2) {
		// Recursively call gcdOfStrings with the remaining part of str1
		return gcdOfStrings(str1[len(str2):], str2)
	}

	// If none of the above conditions are met, there is no GCD
	return ""
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	gcd := gcdOfStrings(str1, str2)
	fmt.Println(gcd) // Output: "ABC"
}
