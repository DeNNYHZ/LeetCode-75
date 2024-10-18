package Prefix_Sum

func largestAltitude(gain []int) int {
	altitude := 0
	highestAltitude := 0

	for gainValue := range gain {
		altitude += gain[gainValue]
		if altitude > highestAltitude {
			highestAltitude = altitude
		}
	}
	return highestAltitude
}
