package day11

import (
	"strings"
	"testing"
)

func TestFindCooupiedSeats(t *testing.T) {
	type args struct {
		inputs [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find occupied seats",
			args: args{
				inputs: [][]string{
					strings.Split("L.LL.LL.LL", ""),
					strings.Split("LLLLLLL.LL", ""),
					strings.Split("L.L.L..L..", ""),
					strings.Split("LLLL.LL.LL", ""),
					strings.Split("L.LL.LL.LL", ""),
					strings.Split("L.LLLLL.LL", ""),
					strings.Split("..L.L.....", ""),
					strings.Split("LLLLLLLLLL", ""),
					strings.Split("L.LLLLLL.L", ""),
					strings.Split("L.LLLLL.LL", ""),
				},
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCooupiedSeats(tt.args.inputs); got != tt.want {
				t.Errorf("FindCooupiedSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOccupiedSeatsBetter(t *testing.T) {
	type args struct {
		inputs [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find occupied seats",
			args: args{
				inputs: [][]string{
					strings.Split("L.LL.LL.LL", ""),
					strings.Split("LLLLLLL.LL", ""),
					strings.Split("L.L.L..L..", ""),
					strings.Split("LLLL.LL.LL", ""),
					strings.Split("L.LL.LL.LL", ""),
					strings.Split("L.LLLLL.LL", ""),
					strings.Split("..L.L.....", ""),
					strings.Split("LLLLLLLLLL", ""),
					strings.Split("L.LLLLLL.L", ""),
					strings.Split("L.LLLLL.LL", ""),
				},
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOccupiedSeatsBetter(tt.args.inputs); got != tt.want {
				t.Errorf("FindOccupiedSeatsBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
