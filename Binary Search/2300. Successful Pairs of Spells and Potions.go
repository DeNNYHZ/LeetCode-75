package Binary_Search

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	n := len(potions)
	result := make([]int, len(spells))

	for i, spell := range spells {
		requiredPotionStrength := (success + int64(spell) - 1) / int64(spell)
		index := sort.Search(n, func(j int) bool {
			return int64(potions[j]) >= requiredPotionStrength
		})

		result[i] = n - index
	}
	return result
}
