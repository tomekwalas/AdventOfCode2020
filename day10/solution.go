package day10

import (
	"sort"
)

func FindAdaptersChain(inputs []int) int {
	sort.Ints(inputs)
	inputs = append(inputs, inputs[len(inputs)-1]+3)
	diffMap := map[int]int{}
	for i := 0; i < len(inputs)-1; i++ {
		diff := inputs[i+1] - inputs[i]
		diffMap[diff]++
	}
	return diffMap[1] * (diffMap[3])
}

func FindAdaptersPermutations(inputs []int) int {
	groups := [][]int{}
	group := []int{}
	permutations := 1
	validPermutations := map[int]int{1: 1, 2: 1, 3: 2, 4: 4, 5: 7}

	sort.Ints(inputs)
	inputs = append(inputs, inputs[len(inputs)-1]+3)

	for i, v := range inputs {
		if i == len(inputs)-1 {
			break
		}
		next := inputs[i+1]
		diff := next - v

		if diff < 3 {
			group = append(group, i)
		} else if diff == 3 && len(group) > 0 {
			group = append(group, i)
			groups = append(groups, group)
			group = []int{}
		}
	}

	for _, g := range groups {
		v, ok := validPermutations[len(g)]
		if ok {
			permutations *= v
		}
	}

	return permutations
}
