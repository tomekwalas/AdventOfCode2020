package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToSlice(filename string) []int {
	inputs := []int{}
	data, err := ioutil.ReadFile(filename)
	check(err)
	for _, inputString := range strings.Split(string(data), "\n") {
		input, err := strconv.Atoi(inputString)
		check(err)
		inputs = append(inputs, input)
	}

	return inputs
}
