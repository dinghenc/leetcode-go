package monostack

import (
	"math"
	"testing"
	"unsafe"
)

func TestNextGreaterNumber(t *testing.T) {
	numbers := []int{2, 1, 5, 6, 2, 3}
	t.Log(NextGreaterNumber(numbers))
	t.Log(NextLessNumber(numbers))
	t.Log(reverse(NextGreaterNumber(reverse(numbers))))
	t.Log(PrevGreaterNumber(numbers))
	t.Log(reverse(NextLessNumber(reverse(numbers))))
	t.Log(PrevLessNumber(numbers))
	t.Log(PrevLessNumberTemp(numbers))

	t.Log(unsafe.Sizeof("abc"))
	t.Log(math.MaxInt)
}

func reverse(number []int) []int {
	rev := make([]int, len(number))
	for i, v := range number {
		rev[len(number)-i-1] = v
	}
	return rev
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func TestGcd(t *testing.T) {
	t.Log(gcd(5, 6))
}
