package day10

import (
	"testing"
)

func TestFindAdaptersChain(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find adapetrs chain",
			args: args{
				inputs: []int{
					0,
					16,
					10,
					15,
					5,
					1,
					11,
					7,
					19,
					6,
					12,
					4,
				},
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAdaptersChain(tt.args.inputs); got != tt.want {
				t.Errorf("FindAdaptersChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAdaptersPermutations(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find adapetrs permtuations",
			args: args{
				inputs: []int{
					0,
					16,
					10,
					15,
					5,
					1,
					11,
					7,
					19,
					6,
					12,
					4,
				},
			},
			want: 8,
		},
		{
			name: "Should find adapetrs permtuations",
			args: args{
				inputs: []int{
					0,
					28,
					33,
					18,
					42,
					31,
					14,
					46,
					20,
					48,
					47,
					24,
					23,
					49,
					45,
					19,
					38,
					39,
					11,
					1,
					32,
					25,
					35,
					8,
					17,
					7,
					9,
					4,
					2,
					34,
					10,
					3,
				},
			},
			want: 19208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAdaptersPermutations(tt.args.inputs); got != tt.want {
				t.Errorf("FindAdaptersPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
