package day3

import (
	"io/ioutil"
	"strings"
)

const (
	tree  = "#"
	blank = "."
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConvertToSlice(filename string) []string {
	inputs := []string{}
	data, err := ioutil.ReadFile(filename)
	check(err)
	for _, inputString := range strings.Split(string(data), "\n") {
		inputs = append(inputs, inputString)
	}

	return inputs
}

func FindTreesNumber(inputs []string, xHop, yHop int) int {
	xPos := 0
	yPos := 0
	counter := 0

	for {
		if yPos >= len(inputs)-1 {
			return counter
		}
		yPos += yHop
		xPos += xHop
		row := inputs[yPos]
		if cell := inputs[yPos][xPos%len(row)]; string(cell) == tree {
			counter++
		}
	}
}

func FindMultipliedTreesNumber(inputs []string) int {
	variants := []struct{ xHop, yHop int }{{xHop: 1, yHop: 1}, {xHop: 3, yHop: 1}, {xHop: 5, yHop: 1}, {xHop: 7, yHop: 1}, {xHop: 1, yHop: 2}}
	counter := 1

	for _, variant := range variants {
		counter *= FindTreesNumber(inputs, variant.xHop, variant.yHop)
	}
	return counter
}
