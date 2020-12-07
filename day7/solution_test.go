package day7

import (
	"testing"
)

func TestSearchForGoldBag(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find number of gold bags",
			args: args{
				inputs: []string{
					"light red bags contain 1 bright white bag, 2 muted yellow bags.",
					"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
					"bright white bags contain 1 shiny gold bag.",
					"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
					"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
					"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
					"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
					"faded blue bags contain no other bags.",
					"dotted black bags contain no other bags.",
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := ConvertToSlice(tt.args.inputs)
			if got := SearchForGoldBagInBags(inputs); got != tt.want {
				t.Errorf("SearchForGoldBag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindBagsNumberInGoldBag(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find number of bags inside golden bag",
			args: args{
				inputs: []string{
					"light red bags contain 1 bright white bag, 2 muted yellow bags.",
					"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
					"bright white bags contain 1 shiny gold bag.",
					"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
					"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
					"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
					"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
					"faded blue bags contain no other bags.",
					"dotted black bags contain no other bags.",
				},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := ConvertToSlice(tt.args.inputs)
			if got := FindBagsNumberInGoldBag(inputs); got != tt.want {
				t.Errorf("FindBagsNumberInGoldBag() = %v, want %v", got, tt.want)
			}
		})
	}
}
