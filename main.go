package main

import (
	"advent-of-code-2020/day1"
	"advent-of-code-2020/day2"
	"fmt"
	"os"
)

func main() {

	selectedDay := os.Args[1]

	switch selectedDay {
	case "day1":
		{
			inputs := convertToSlice("./day-1/input.txt")
			fmt.Printf("Solution for Day 1 part 1 is: %d\n", day1.FindSolution(inputs, 2020, 1, 2))
			fmt.Printf("Solution for Day 1 part 2 is: %d\n", day1.FindSolution(inputs, 2020, 1, 3))
			return
		}
	case "day2":
		{
			inputs := day2.ConvertToSlice("./day2/input.txt")
			fmt.Printf("Solution for Day2 part 1 is: %d\n", day2.GetValidPasswordsNumber(inputs, day2.ValidateOldCompanyPassword))
			fmt.Printf("Solution for Day2 part 2 is: %d\n", day2.GetValidPasswordsNumber(inputs, day2.ValidateCurrenctCompanyPassword))
		}

	}

}
