package day8

import (
	"testing"
)

func TestFindAccumulatorValue(t *testing.T) {
	type args struct {
		cmds []cmd
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should get accumulator value",
			args: args{
				cmds: []cmd{
					{
						name:  "nop",
						value: 0,
					},
					{
						name:  "acc",
						value: 1,
					},
					{
						name:  "jmp",
						value: 4,
					},
					{
						name:  "acc",
						value: 3,
					},
					{
						name:  "jmp",
						value: -3,
					},
					{
						name:  "acc",
						value: -99,
					},
					{
						name:  "acc",
						value: 1,
					},
					{
						name:  "jmp",
						value: -4,
					},
					{
						name:  "acc",
						value: 6,
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAccumulatorValue(tt.args.cmds); got != tt.want {
				t.Errorf("FindAccumulatorValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixProgram(t *testing.T) {
	type args struct {
		cmds []cmd
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should fix program loop and get correct accumulator value",
			args: args{
				cmds: []cmd{
					{
						name:  "nop",
						value: 0,
					},
					{
						name:  "acc",
						value: 1,
					},
					{
						name:  "jmp",
						value: 4,
					},
					{
						name:  "acc",
						value: 3,
					},
					{
						name:  "jmp",
						value: -3,
					},
					{
						name:  "acc",
						value: -99,
					},
					{
						name:  "acc",
						value: 1,
					},
					{
						name:  "jmp",
						value: -4,
					},
					{
						name:  "acc",
						value: 6,
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixProgram(tt.args.cmds); got != tt.want {
				t.Errorf("FixProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
