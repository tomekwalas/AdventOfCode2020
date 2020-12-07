package day7

import (
	"strconv"
	"strings"
)

type containingBag struct {
	name   string
	number int
}
type bag struct {
	name     string
	contains []containingBag
}

func ConvertToSlice(inputs []string) []bag {
	bags := []bag{}
	for _, input := range inputs {
		values := strings.Split(input, "contain")
		containedBags := []containingBag{}
		name := parseName(values[0])
		if !strings.Contains(values[1], "no other bags") {
			for _, bag := range strings.Split(values[1], ",") {
				bagName := parseName(bag[2:])
				bagNumber, err := strconv.Atoi(bag[1:2])
				if err != nil {
					panic(err.Error())
				}
				containedBags = append(containedBags, containingBag{
					name:   strings.TrimSpace(bagName),
					number: bagNumber,
				})

			}
		}
		b := bag{
			name:     strings.TrimSpace(name),
			contains: containedBags,
		}
		bags = append(bags, b)
	}

	return bags
}
func parseName(value string) string {
	name := strings.Replace(value, "bags", "", -1)
	name = strings.Replace(name, "bag", "", -1)
	if strings.Contains(name, ".") {
		name = strings.Replace(name, ".", "", -1)
	}
	return name
}

func SearchForGoldBagInBags(inputs []bag) int {
	bagMap := map[string]int{}
	searchForBagInBags(inputs, "shiny gold", bagMap)
	return len(bagMap)
}

func FindBagsNumberInGoldBag(inputs []bag) int {
	return searchForBagsNumberInBag(inputs, "shiny gold") - 1
}

func searchForBagInBags(bags []bag, searchedBag string, bagMap map[string]int) {
	for _, bag := range bags {
		for _, containedBag := range bag.contains {
			if strings.Contains(containedBag.name, searchedBag) {
				bagMap[bag.name]++
				searchForBagInBags(bags, bag.name, bagMap)
			}
		}
	}
}

func searchForBagsNumberInBag(bags []bag, searchedBag string) int {
	counter := 1
	for _, bag := range bags {
		if bag.name != searchedBag {
			continue
		}
		for _, containedBag := range bag.contains {
			count := searchForBagsNumberInBag(bags, containedBag.name)
			counter += containedBag.number * count
		}
	}
	return counter
}
