package day11

import (
	"advent-of-code-2020/utils"
	"strings"
)

const (
	floor        = "."
	emptySeat    = "L"
	occupiedSeat = "#"
)

func ConvertToSlice(filename string) [][]string {
	inputs := [][]string{}
	for _, input := range utils.ReadFile(filename) {
		inputs = append(inputs, strings.Split(input, ""))
	}

	return inputs
}

func FindCooupiedSeats(inputs [][]string) int {
	changes := 0
	for {
		inputsCopy := copyDeep(inputs)
		for i, input := range inputs {
			for j, char := range input {
				if char == floor {
					continue
				}
				occupied := 0
				directions := [][]int{{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1}, {i, j - 1}, {i, j + 1}, {i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}}

				for _, direction := range directions {
					if isIdxValid(direction[0], len(inputs)) && isIdxValid(direction[1], len(input)) && inputs[direction[0]][direction[1]] == occupiedSeat {
						occupied++
					}
				}

				if char == emptySeat && occupied == 0 {
					inputsCopy[i][j] = occupiedSeat
					changes++
				} else if char == occupiedSeat && occupied >= 4 {
					inputsCopy[i][j] = emptySeat
					changes++
				}
			}
		}
		if changes == 0 {
			allOccupied := 0
			for _, input := range inputsCopy {
				for _, char := range input {
					if char == occupiedSeat {
						allOccupied++
					}
				}
			}
			return allOccupied
		}
		changes = 0
		inputs = copyDeep(inputsCopy)
	}
}

func FindOccupiedSeatsBetter(inputs [][]string) int {
	isOccupied := func(y, x, yDirection, xDirection, yLen, xLen int) bool {
		for {
			y += yDirection
			x += xDirection
			if isIdxValid(y, yLen) && isIdxValid(x, xLen) {
				if inputs[y][x] == emptySeat {
					return false
				}
				if inputs[y][x] == occupiedSeat {
					return true
				}
			} else {
				return false
			}
		}

	}
	changes := 0
	for {
		inputsCopy := copyDeep(inputs)
		for i, input := range inputs {
			for j, char := range input {
				if char == floor {
					continue
				}
				occupied := 0
				directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
				for _, direction := range directions {
					if isOccupied(i, j, direction[0], direction[1], len(inputs), len(input)) {
						occupied++
					}
				}

				if char == emptySeat && occupied == 0 {
					inputsCopy[i][j] = occupiedSeat
					changes++
				} else if char == occupiedSeat && occupied >= 5 {
					inputsCopy[i][j] = emptySeat
					changes++
				}

			}
		}
		if changes == 0 {
			allOccupied := 0
			for _, input := range inputsCopy {
				for _, char := range input {
					if char == occupiedSeat {
						allOccupied++
					}
				}
			}
			return allOccupied
		}
		changes = 0
		inputs = copyDeep(inputsCopy)
	}
}

func isIdxValid(i, l int) bool {
	return i >= 0 && i < l
}

func copyDeep(src [][]string) [][]string {
	duplicate := make([][]string, len(src))
	for i := range src {
		duplicate[i] = make([]string, len(src[i]))
		copy(duplicate[i], src[i])
	}
	return duplicate
}
