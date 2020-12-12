package day12

import (
	"advent-of-code-2020/utils"
	"math"
	"strconv"
	"strings"
)

type direction struct {
	units int
	value string
}

const (
	north   = "N"
	south   = "S"
	west    = "W"
	east    = "E"
	forward = "F"
	left    = "L"
	right   = "R"
)

func ConvertToSlice(filename string) []direction {
	directions := []direction{}
	for _, input := range utils.ReadFile(filename) {
		d := input[:1]
		units, err := strconv.Atoi(strings.TrimSpace(input[1:]))
		if err != nil {
			panic(err.Error())
		}
		direction := direction{
			value: d,
			units: units,
		}
		directions = append(directions, direction)
	}
	return directions
}

func FindBoatDirection(inputs []direction) int {
	directions := map[string]int{north: 0, south: 0, east: 0, west: 0}
	compass := map[int]string{0: north, 90: east, 180: south, 270: west}
	currentPos := 90

	for _, direction := range inputs {
		if direction.value == forward {
			directions[compass[currentPos]] += direction.units
		} else if direction.value == left {
			currentPos = (currentPos + 360 - direction.units) % 360

		} else if direction.value == right {
			currentPos = (currentPos + direction.units) % 360

		} else {
			directions[direction.value] += direction.units
		}
	}

	result := math.Abs(float64(directions[north]-directions[south])) + math.Abs(float64(directions[west]-directions[east]))
	return int(result)
}

func FindBoatDirectionFromWaypoint(inputs []direction) int {
	directions := map[string]int{north: 0, south: 0, east: 0, west: 0}
	compass := map[int]string{0: north, 90: east, 180: south, 270: west}
	reverseCompass := map[string]int{north: 0, east: 90, south: 180, west: 270}
	waypoint := map[int]int{90: 10, 0: 1}

	for _, input := range inputs {
		if input.value == forward {
			for d, v := range waypoint {
				directions[compass[d]] += v * input.units
			}
			continue
		}
		if input.value == left || input.value == right {
			newWaypoint := map[int]int{}
			for d, v := range waypoint {
				var currentPos int
				if input.value == right {
					currentPos = (d + input.units) % 360
				} else {
					currentPos = (d + 360 - input.units) % 360
				}
				newWaypoint[currentPos] = v
			}
			waypoint = newWaypoint
			continue
		}

		direction := reverseCompass[input.value]
		if _, ok := waypoint[direction]; ok {
			waypoint[direction] += input.units
			continue
		}

		reversedDirection := (direction + 180) % 360
		if diff := waypoint[reversedDirection] - input.units; diff >= 0 {
			waypoint[reversedDirection] -= input.units
			continue
		} else {
			delete(waypoint, direction)
			waypoint[reversedDirection] = diff
		}
	}

	result := math.Abs(float64(directions[north]-directions[south])) + math.Abs(float64(directions[west]-directions[east]))
	return int(result)
}
