package palindrome

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := range dp {
		dp[i][i] = true
	}

	begin := 0
	maxLen := 1
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if i+1 == j {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] && maxLen < l {
				maxLen = l
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
