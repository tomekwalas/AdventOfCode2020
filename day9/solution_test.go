package day9

import (
	"testing"
)

func TestFindNumber(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should find invalid number",
			args: args{
				inputs: []int{
					35,
					20,
					15,
					25,
					47,
					40,
					62,
					55,
					65,
					95,
					102,
					117,
					150,
					182,
					127,
					219,
					299,
					277,
					309,
					576,
				},
			},
			want: 127,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumber(tt.args.inputs, 5); got != tt.want {
				t.Errorf("FindNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindContiguousSet(t *testing.T) {
	type args struct {
		inputs []int
		sum    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should find sum of smallest and largest number in contiguous set of sum",
			args: args{
				inputs: []int{
					35,
					20,
					15,
					25,
					47,
					40,
					62,
					55,
					65,
					95,
					102,
					117,
					150,
					182,
					127,
					219,
					299,
					277,
					309,
					576,
				},
				sum: 127,
			},
			want: 62,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindContiguousSet(tt.args.inputs, tt.args.sum); got != tt.want {
				t.Errorf("FindContiguousSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
