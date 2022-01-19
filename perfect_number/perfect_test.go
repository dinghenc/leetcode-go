package perfect_number

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	for i := 1; i <= 1e8; i++ {
		if CheckPerfectNumber(i) {
			t.Log(i)
		}
	}
	t.Log("done")
}
