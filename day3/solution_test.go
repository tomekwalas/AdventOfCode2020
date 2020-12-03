package day3

import (
	"testing"
)

func TestFindTreesNumber(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find trees number",
			args: args{inputs: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTreesNumber(tt.args.inputs, 3, 1); got != tt.want {
				t.Errorf("FindTreesNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMultipliedTreesNumber(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find multiplied trees number",
			args: args{inputs: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			}},
			want: 336,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMultipliedTreesNumber(tt.args.inputs); got != tt.want {
				t.Errorf("FindMultipliedTreesNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
