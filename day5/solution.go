package day5

import (
	"sort"
)

//FindHighestBoardingPassID finds highest boarding pass id
func FindHighestBoardingPassID(inputs []string) int {
	highestID := 0
	ids := []int{}
	for _, input := range inputs {
		row := findRow(input[:7])
		column := findColumn(input[7:])
		id := row*8 + column
		ids = append(ids, id)

		if id > highestID {
			highestID = id
		}
	}

	return highestID
}

//FindYourSeat finds your seat
func FindYourSeat(inputs []string) int {
	ids := []int{}
	for _, input := range inputs {
		row := findRow(input[:7])
		column := findColumn(input[7:])
		id := row*8 + column
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		if ids[i]+1 != ids[i+1] {
			return ids[i] + 1
		}
	}
	return 0
}

func findRow(inputs string) int {
	return search(0, 127, 0, inputs)
}

func findColumn(inputs string) int {
	return search(0, 7, 0, inputs)
}

func search(from, to, idx int, direction string) int {
	if len(direction) == idx+1 {
		if string(direction[idx]) == "F" || string(direction[idx]) == "L" {
			return from
		}
		return to
	}

	size := (to - from) + 1
	size /= 2
	if string(direction[idx]) == "F" || string(direction[idx]) == "L" {
		return search(from, from+size-1, idx+1, direction)
	}
	return search(from+size, to, idx+1, direction)
}
