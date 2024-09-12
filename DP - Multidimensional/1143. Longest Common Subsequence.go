package main

func longestCommonSubsequence(text1 string, text2 string) int {
	// Membuat tabel 2D untuk menyimpan panjang subsequence
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	// Mengisi tabel dengan panjang subsequence yang sesuai
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	// Mengembalikan panjang subsequence yang paling panjang
	return dp[len(text1)][len(text2)]
}
