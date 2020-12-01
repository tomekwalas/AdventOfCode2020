package day1

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
