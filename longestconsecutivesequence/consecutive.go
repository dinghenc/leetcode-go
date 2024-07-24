package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	maxLength := 0
	dp := make([]int, len(nums))
	for i, v := range nums {
		if dp[i] != 0 {
			continue
		}
		num := v
		length := 0
		for {
			preIndex, ok := m[num]
			if !ok {
				break
			}
			if dp[preIndex] != 0 {
				length += dp[preIndex]
				break
			}
			num = nums[preIndex] - 1
			length++
		}
		maxLength = max(maxLength, length)
		for p := v; p > num; p-- {
			dp[m[p]] = length
			length--
		}
	}
	return maxLength
}
