package palindrome

import (
	"fmt"
	"sort"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{s: "babad"},
			want: "bab",
		}, {
			name: "2",
			args: args{s: "cbbd"},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBestHand(t *testing.T) {
	t.Log(bestHand([]int{4, 1, 2, 3, 5}, []byte{'b', 'a', 'a', 'a', 'a'}))

	t.Log(zeroFilledSubarray([]int{2, 10, 2019}))
}

func bestHand(ranks []int, suits []byte) string {
	if equalSuits(suits) {
		return "Flush"
	} else if equalRanks(ranks, 3) {
		return "Three of a Kind"
	} else if equalRanks(ranks, 2) {
		return "Pair"
	}
	return "High Card"
}

func equalSuits(suits []byte) bool {
	m := make(map[byte]int)
	for _, v := range suits {
		m[v]++
	}
	return len(m) == 1
}

func equalRanks(ranks []int, count int) bool {
	m := make(map[int]int)
	for _, v := range ranks {
		m[v]++
	}
	for _, v := range m {
		if v >= count {
			return true
		}
	}
	return false
}

func zeroFilledSubarray(nums []int) int64 {
	count := int64(0)
	ans := int64(0)
	for _, v := range nums {
		if v == 0 {
			count++
		} else {
			ans += count * (1 + count) / 2
			count = 0
		}
	}
	ans += count * (1 + count) / 2
	return ans
}

func TestCountExcellentPairs(t *testing.T) {
	t.Log(countExcellentPairs([]int{1, 2, 3, 1}, 3))
	t.Log(countExcellentPairs([]int{1, 2, 3, 1, 536870911}, 3))
}

func countExcellentPairs(nums []int, k int) int64 {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v] = bitCount(v)
	}
	fmt.Println(counts)
	sort.Slice(nums, func(i, j int) bool {
		if counts[nums[i]] == counts[nums[j]] {
			return nums[i] < nums[j]
		}
		return counts[nums[i]] < counts[nums[j]]
	})
	fmt.Println(nums)
	left, right := 0, len(nums)-1
	m := make(map[int]bool)
	ans := int64(0)
	for left < len(nums) && right >= 0 {
		m[nums[right]] = true
		sum := counts[nums[left]] + counts[nums[right]]
		if sum < k {
			left++
			for left < len(nums) && left != 0 && nums[left] == nums[left-1] {
				left++
			}
			ans += int64(len(m)) - 1
		} else if sum >= k {
			right--
		}
	}
	if counts[nums[0]]+counts[nums[len(nums)-1]] >= k {
		ans += int64(len(m))
	}
	return ans
}

func bitCount(n int) int {
	c := 0
	for n != 0 {
		c++
		n &= (n - 1)
	}
	return c
}
