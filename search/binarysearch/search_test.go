package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 12, 13, 14}
	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}
	for i := 0; i < 20; i++ {
		if _, ok := m[i]; !ok {
			m[i] = -1
		}
		assert.Equal(t, m[i], Locate(arr, i), fmt.Sprintf("target=%d", i))
	}
}
