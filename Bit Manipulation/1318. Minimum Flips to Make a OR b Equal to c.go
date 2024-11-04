package Bit_Manipulation

func minFlips(a int, b int, c int) int {
	flips := 0

	for a > 0 || b > 0 || c > 0 {
		bit_a := a & 1
		bit_b := b & 1
		bit_c := c & 1

		if bit_c == 1 {
			if bit_a == 0 && bit_b == 0 {
				flips += 1
			}
		} else {
			flips += bit_a + bit_b
		}
		a >>= 1
		b >>= 1
		c >>= 1
	}
	return flips
}
