package minmutation

import (
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	i := int32(1)
	if atomic.CompareAndSwapInt32(&i, 1, 2) {
		t.Log("CAS done: ", i)
		return
	}
	t.Log("No CAS: ", i)
}

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				start: "AACCGGTT",
				end:   "AACCGGTA",
				bank:  []string{"AACCGGTA"},
			},
			want: 1,
		}, {
			name: "test 2",
			args: args{
				start: "AACCGGTT",
				end:   "AAACGGTA",
				bank:  []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			want: 2,
		}, {
			name: "test 3",
			args: args{
				start: "AAAAACCC",
				end:   "AACCCCCC",
				bank:  []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
