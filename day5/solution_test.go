package day5

import "testing"

func TestFindHighestBoardingPassId(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find highest boarding pass ID",
			args: args{
				inputs: []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"},
			},
			want: 820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindHighestBoardingPassID(tt.args.inputs); got != tt.want {
				t.Errorf("FindHighestBoardingPassID() = %v, want %v", got, tt.want)
			}
		})
	}
}
