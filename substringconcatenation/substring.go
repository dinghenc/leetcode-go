package substringconcatenation

import "strings"

func findSubstring(s string, words []string) []int {
	indexOfWords := make(map[string][]int)
	for _, w := range words {
		begin, ns := 0, s
		var indexArr []int
		index := strings.Index(ns, w)
		for index >= 0 {
			indexArr = append(indexArr, index+begin)
			begin += index + 1
			ns = ns[index+1:]
			index = strings.Index(ns, w)
		}
		if len(indexArr) == 0 {
			return nil
		}
		indexOfWords[w] = indexArr
	}
	indexWord := make([]string, len(s))
	for w, is := range indexOfWords {
		for _, i := range is {
			indexWord[i] = w
		}
	}

	wordCount := make(map[string]int)
	for _, w := range words {
		wordCount[w]++
	}

	wordLen := len(words[0])
	N := len(words)
	var ans []int
	for left := 0; left < len(s); left++ {
		m := make(map[string]int)
		cnt := 0
		for right := left; right < len(s); {
			if w := indexWord[right]; w == "" || m[w] >= wordCount[w] {
				break
			} else if cnt+1 == N {
				ans = append(ans, left)
				break
			} else {
				m[w]++
				right += wordLen
				cnt++
			}
		}
	}
	return ans
}
