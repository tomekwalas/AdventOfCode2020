package day8

import (
	"advent-of-code-2020/utils"
	"strconv"
	"strings"
)

type cmd struct {
	name  string
	value int
}

//ConvertToSlice converting file to input slice
func ConvertToSlice(filename string) []cmd {
	cmds := []cmd{}
	for _, input := range utils.ReadFile(filename) {
		c := strings.Split(input, " ")
		v, err := strconv.Atoi(c[1])
		if err != nil {
			panic(err)
		}

		cmds = append(cmds, cmd{
			name:  c[0],
			value: v,
		})
	}

	return cmds
}

//FindAccumulatorValue finds accumulator value before cmd infinite loop
func FindAccumulatorValue(cmds []cmd) int {
	return runProgram(cmds, false)
}

//FixProgram fixes program loop and returns correct accumulator value
func FixProgram(cmds []cmd) int {
	for i, cmd := range cmds {
		lastName := ""
		if cmd.name == "acc" {
			continue
		}
		if cmd.name == "nop" {
			lastName = "nop"
			(&cmds[i]).name = "jmp"
		} else if cmd.name == "jmp" {
			lastName = "jmp"
			(&cmds[i]).name = "nop"
		}

		result := runProgram(cmds, true)
		if result > 0 {
			return result
		}
		(&cmds[i]).name = lastName
	}

	return 0
}

func runProgram(cmds []cmd, breakOnCmdRepeat bool) int {
	accumulator := 0
	i := 0
	execMap := map[int]int{}

	for {
		if execMap[i] > 0 {
			if breakOnCmdRepeat {
				return 0
			}
			return accumulator
		}
		if i >= len(cmds) {
			return accumulator
		}
		c := cmds[i]
		if c.name == "nop" {
			execMap[i]++
			i++
		} else if c.name == "acc" {
			accumulator += c.value
			execMap[i]++
			i++
		} else if c.name == "jmp" {
			execMap[i]++
			i += c.value
		}
	}
}
