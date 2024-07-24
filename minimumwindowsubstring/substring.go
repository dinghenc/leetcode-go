package minimumwindowsubstring

func check(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	counts := make(map[byte]int)
	for _, c := range []byte(t) {
		counts[c]++
	}

	m := make(map[byte]int)
	substr := ""
	for left, right := 0, 0; right < len(s); {
		for ; right < len(s) && !check(counts, m); right++ {
			if counts[s[right]] == 0 {
				continue
			}
			m[s[right]]++
		}
		for ; left < right && check(counts, m); left++ {
			if substr == "" || len(substr) > right-left {
				substr = s[left:right]
			}
			if counts[s[left]] == 0 {
				continue
			}
			m[s[left]]--
		}
	}
	return substr
}
