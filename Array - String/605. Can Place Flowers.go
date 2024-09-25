package Array___String

func canPlaceFlowers(flowerbed []int, n int) bool {
	var res int
	for i := 0; i < len(flowerbed); i++ {
		left, right := false, false
		if i == 0 || flowerbed[i-1] == 0 {
			left = true
		}
		if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
			right = true
		}
		if flowerbed[i] == 0 && left && right {
			flowerbed[i] = 1
			res++
		}
		if res >= n {
			return true
		}
	}
	return false
}
