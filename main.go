package main

import (
	"advent-of-code-2020/day1"
	"advent-of-code-2020/day2"
	"advent-of-code-2020/day3"
	"advent-of-code-2020/day4"
	"fmt"
	"os"
)

func main() {

	selectedDay := os.Args[1]

	switch selectedDay {
	case "day1":
		{
			inputs := day1.ConvertToSlice("./day-1/input.txt")
			fmt.Printf("Solution for Day 1 part 1 is: %d\n", day1.FindSolution(inputs, 2020, 1, 2))
			fmt.Printf("Solution for Day 1 part 2 is: %d\n", day1.FindSolution(inputs, 2020, 1, 3))
			return
		}
	case "day2":
		{
			inputs := day2.ConvertToSlice("./day2/input.txt")
			fmt.Printf("Solution for Day 2 part 1 is: %d\n", day2.GetValidPasswordsNumber(inputs, day2.ValidateOldCompanyPassword))
			fmt.Printf("Solution for Day 2 part 2 is: %d\n", day2.GetValidPasswordsNumber(inputs, day2.ValidateCurrenctCompanyPassword))
			return
		}
	case "day3":
		{
			inputs := day3.ConvertToSlice("./day3/input.txt")
			fmt.Printf("Solution for Day 3 part 1 is: %d\n", day3.FindTreesNumber(inputs, 3, 1))
			fmt.Printf("Solution for Day 3 part 2 is: %d\n", day3.FindMultipliedTreesNumber(inputs))
			return
		}
	case "day4":
		{
			inputs := day4.ConvertToSlice("./day4/input.txt")
			fmt.Printf("Solution for Day 4 part 1 is: %d\n", day4.FindValidPassportsNumber(inputs, false))
			fmt.Printf("Solution for Day 4 part 2 is: %d\n", day4.FindValidPassportsNumber(inputs, true))
			return
		}
	}
}
