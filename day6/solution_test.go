package day6

import (
	"testing"
)

func TestFindSumOfAnswers(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find sum of answered questions",
			args: args{
				inputs: []string{
					"abc",
					"",
					"a",
					"b",
					"c",
					"",
					"ab",
					"ac",
					"",
					"a",
					"a",
					"a",
					"a",
					"",
					"b",
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSumOfAnswers(tt.args.inputs); got != tt.want {
				t.Errorf("FindSumOfAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSumOfAllAnswers(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find sum of all answered questions",
			args: args{
				inputs: []string{
					"abc",
					"",
					"a",
					"b",
					"c",
					"",
					"ab",
					"ac",
					"",
					"a",
					"a",
					"a",
					"a",
					"",
					"b",
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSumOfAllAnswers(tt.args.inputs); got != tt.want {
				t.Errorf("FindSumOfAllAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
