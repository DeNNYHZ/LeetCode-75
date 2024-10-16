package Hash_Map___Set

import "strconv"

func toSet(s string) map[rune]struct{} {
	set := make(map[rune]struct{})
	for _, char := range s {
		set[char] = struct{}{}
	}
	return set
}

func haveSameAlphabets(word1, word2 string) bool {
	word1 = strconv.Itoa(len(word1))
	word2 = strconv.Itoa(len(word2))

	if word1 != word2 {
		return false
	}

	set1 := toSet(word1)
	set2 := toSet(word2)

	if len(set1) != len(set2) {
		return false
	}

	for char := range set1 {
		if _, ok := set2[char]; !ok {
			return false
		}
	}
	return true

}
