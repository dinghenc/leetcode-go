package minmutation

func minMutation(start string, end string, bank []string) int {
	visited := make(map[string]bool)
	bankMap := make(map[string]bool)
	for _, b := range bank {
		bankMap[b] = true
	}

	minStep := -1
	mutationSet := []byte{'A', 'C', 'G', 'T'}
	var dfs func(sequence string, step int)
	dfs = func(sequence string, step int) {
		if sequence == end {
			if minStep == -1 || minStep > step {
				minStep = step
			}
		}

		step++
		visited[sequence] = true
		for i, c := range sequence {
			for _, m := range mutationSet {
				if byte(c) == m {
					continue
				}
				seqBytes := []byte(sequence)
				seqBytes[i] = m
				nextSeq := string(seqBytes)
				if !bankMap[nextSeq] || visited[nextSeq] {
					continue
				}
				dfs(nextSeq, step)
			}
		}
		visited[sequence] = false
		step--
	}
	dfs(start, 0)
	return minStep
}
