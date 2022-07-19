package onlinevote

import "sort"

type TopVotedCandidate struct {
	tops  []int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	voteCounts := map[int]int{-1: -1}
	top := -1
	var tops []int
	for _, p := range persons {
		voteCounts[p]++
		if voteCounts[p] >= voteCounts[top] {
			top = p
		}
		tops = append(tops, top)
	}
	return TopVotedCandidate{
		tops:  tops,
		times: times,
	}
}

func (c *TopVotedCandidate) Q(t int) int {
	return c.tops[sort.SearchInts(c.times, t+1)-1]
}
