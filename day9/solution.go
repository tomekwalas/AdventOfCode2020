package day9

import (
	"advent-of-code-2020/day1"
	"sort"
)

func FindNumber(inputs []int, offset int) int {

	for i, v := range inputs[offset:] {
		preamble := inputs[i : offset+i]
		if ok := day1.FindSolution(preamble, v, 1, 2); ok > 0 {
			continue
		}
		return v
	}
	return 0
}

func FindContiguousSet(inputs []int, sum int) int {
	slice := []int{}
	currentSum := 0
	for i := 0; i < len(inputs); i++ {
		currentSum = inputs[i]
		for j := i + 1; j <= len(inputs); j++ {
			if currentSum == sum {
				slice = append(slice, inputs[i:j-1]...)
				break
			}
			if currentSum > sum || j == len(inputs) {
				break
			}
			currentSum += inputs[j]
		}
	}

	sort.Ints(slice)

	return slice[0] + slice[len(slice)-1]
}
