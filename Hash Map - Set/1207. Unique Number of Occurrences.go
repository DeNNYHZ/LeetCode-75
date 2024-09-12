package Hash_Map___Set

func uniqueOccurrences(arr []int) bool {
	countOccurMap := map[int]int{}
	markOccurMap := map[int]bool{}

	for _, value := range arr {
		countOccurMap[value] = countOccurMap[value] + 1
	}

	for _, value := range countOccurMap {
		markOccurMap[value] = true
	}
	return len(countOccurMap) == len(markOccurMap)
}
