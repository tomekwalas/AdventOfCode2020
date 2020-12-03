package utils

import (
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//ReadFile reads input file and coverts it to string slice
func ReadFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return strings.Split(string(data), "\n")
}
