package day1

import (
	"testing"
)

func Test_FindSolution(t *testing.T) {
	type args struct {
		inputs       []int
		partialInput int
		depth        int
		maxDeph      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should calculate first part correctly",
			args: args{
				inputs:       []int{1721, 979, 366, 299, 675, 1456},
				partialInput: 2020,
				depth:        1,
				maxDeph:      2,
			},
			want: 514579,
		},
		{
			name: "Should calculate second part correctly",
			args: args{
				inputs:       []int{1721, 979, 366, 299, 675, 1456},
				partialInput: 2020,
				depth:        1,
				maxDeph:      3,
			},
			want: 241861950,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSolution(tt.args.inputs, tt.args.partialInput, tt.args.depth, tt.args.maxDeph); got != tt.want {
				t.Errorf("FindSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
