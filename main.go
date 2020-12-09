package main

import (
	"advent-of-code-2020/day1"
	"advent-of-code-2020/day2"
	"advent-of-code-2020/day3"
	"advent-of-code-2020/day4"
	"advent-of-code-2020/day5"
	"advent-of-code-2020/day6"
	"advent-of-code-2020/day7"
	"advent-of-code-2020/day8"
	"advent-of-code-2020/day9"
	"advent-of-code-2020/utils"
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
	case "day5":
		{
			inputs := utils.ReadFile("./day5/input.txt")
			fmt.Printf("Solution for Day 5 part 1 is: %d\n", day5.FindHighestBoardingPassID(inputs))
			fmt.Printf("Solution for Day 5 part 2 is: %d\n", day5.FindYourSeat(inputs))
		}
	case "day6":
		{
			inputs := utils.ReadFile("./day6/input.txt")
			fmt.Printf("Solution for Day 6 part 1 is: %d\n", day6.FindSumOfAnswers(inputs))
			fmt.Printf("Solution for Day 6 part 2 is: %d\n", day6.FindSumOfAllAnswers(inputs))
		}
	case "day7":
		{
			inputs := utils.ReadFile("./day7/input.txt")
			bags := day7.ConvertToSlice(inputs)
			fmt.Printf("Solution for Day 7 part 1 is: %d\n", day7.SearchForGoldBagInBags(bags))
			fmt.Printf("Solution for Day 7 part 2 is: %d\n", day7.FindBagsNumberInGoldBag(bags))
		}
	case "day8":
		{
			inputs := day8.ConvertToSlice("./day8/input.txt")
			fmt.Printf("Solution for Day 8 part 1 is: %d\n", day8.FindAccumulatorValue(inputs))
			fmt.Printf("Solution for Day 8 part 2 is: %d\n", day8.FixProgram(inputs))
		}
	case "day9":
		{
			inputs := day1.ConvertToSlice("./day9/input.txt")
			result := day9.FindNumber(inputs, 25)
			fmt.Printf("Solution for Day 9 part 1 is: %d\n", result)
			fmt.Printf("Solution for Day 9 part 2 is: %d\n", day9.FindContiguousSet(inputs, result))
		}
	}
}
