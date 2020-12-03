package day1

import (
	"advent-of-code-2020/utils"
	"strconv"
)

//ConvertToSlice converting file to input slice
func ConvertToSlice(filename string) []int {
	inputs := []int{}
	for _, inputString := range utils.ReadFile(filename) {
		input, err := strconv.Atoi(inputString)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, input)
	}

	return inputs

}

// FindSolution finds solution for day 1 problem
func FindSolution(inputs []int, partialInput int, depth int, maxDeph int) int {
	if depth == maxDeph {
		for _, input := range inputs {
			if partialInput == input {
				return partialInput
			}
		}
		return 0
	}

	for _, input := range inputs {
		inputToSearch := partialInput - input
		if inputToSearch < 0 {
			continue
		}
		partialInput := FindSolution(inputs, inputToSearch, depth+1, maxDeph)
		if partialInput == 0 {
			continue
		}
		return input * partialInput
	}

	return 0
}
