package happynumber

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)
	for !m[n] {
		sum := 0
		for ; n != 0; n /= 10 {
			v := n % 10
			sum += v * v
		}
		if sum == 1 {
			return true
		}
		m[n] = true
		n = sum
		fmt.Println(sum)
		if len(m) > 1 {
			return false
		}
	}
	return false
}
