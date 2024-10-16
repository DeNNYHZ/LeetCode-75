package Hash_Map___Set

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, c := range word1 {
		m1[c]++
	}
	for _, c := range word2 {
		m2[c]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
