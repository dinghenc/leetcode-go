package onlinevote

import (
	"reflect"
	"testing"
)

func TestTopVotedCandidate_Q(t *testing.T) {
	type fields struct {
		persons []int
		times   []int
	}
	type args struct {
		times []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "normal vote",
			fields: fields{
				persons: []int{0, 1, 1, 0, 0, 1, 0},
				times:   []int{0, 5, 10, 15, 20, 25, 30},
			},
			args: args{
				times: []int{3, 12, 25, 15, 24, 8},
			},
			want: []int{0, 1, 1, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor(tt.fields.persons, tt.fields.times)
			var gots []int
			for _, time := range tt.args.times {
				got := c.Q(time)
				gots = append(gots, got)
			}
			if !reflect.DeepEqual(gots, tt.want) {
				t.Errorf("excpted want=%v, but got=%v", tt.want, gots)
			}
		})
	}
}
