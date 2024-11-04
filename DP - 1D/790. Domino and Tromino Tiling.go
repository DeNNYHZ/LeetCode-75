package DP___1D

const mod = 1e9 + 7

func numTilings(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 5
	}

	// Dynamic programming array
	dp := make([]int, n+1)
	dp[0] = 1 // One way to tile an empty board
	dp[1] = 1 // One way to tile a 2x1 board
	dp[2] = 2 // Two ways to tile a 2x2 board
	dp[3] = 5 // Five ways to tile a 2x3 board

	// Fill the dp array for n >= 4
	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]*2 + dp[i-3]) % mod
		if i > 3 {
			dp[i] = (dp[i] + dp[i-4]) % mod
		}
	}

	return dp[n]
}
