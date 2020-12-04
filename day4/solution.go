package day4

import (
	"advent-of-code-2020/utils"
	"regexp"
	"strconv"
	"strings"
)

//ConvertToSlice converting file to input slice
func ConvertToSlice(filename string) []string {
	return utils.ReadFile(filename)
}

type passwordField struct {
	value         string
	validatorFunc func(string) bool
}

//FindValidPassportsNumber finds valid passports number
func FindValidPassportsNumber(inputs []string, strict bool) int {
	validFields := []passwordField{
		{
			value:         "byr",
			validatorFunc: byrValidatorFunc,
		},
		{
			value:         "iyr",
			validatorFunc: iyrValidatorFunc,
		},
		{
			value:         "eyr",
			validatorFunc: eyrValidatorFunc,
		},
		{
			value:         "hgt",
			validatorFunc: hgtValidatorFunc,
		},
		{
			value:         "hcl",
			validatorFunc: hclValidatorFunc,
		},
		{
			value:         "ecl",
			validatorFunc: eclValidatorFunc,
		},
		{
			value:         "pid",
			validatorFunc: pidValidatorFunc,
		},
	}

	validPassportsNumber := 0
	validFieldsNumber := 0
	for _, input := range inputs {
		if input == "" {
			if validFieldsNumber == len(validFields) {
				validPassportsNumber++
			}
			validFieldsNumber = 0
			continue
		}

		fields := strings.Split(input, " ")
		for _, field := range fields {
			keyValue := strings.Split(field, ":")
			for _, passwordField := range validFields {
				if passwordField.value == keyValue[0] && (!strict || passwordField.validatorFunc(keyValue[1])) {
					validFieldsNumber++
					break
				}
			}
		}
	}
	if validFieldsNumber == len(validFields) {
		validPassportsNumber++
	}

	return validPassportsNumber
}

func byrValidatorFunc(value string) bool {
	byr, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return byr >= 1920 && byr <= 2002
}

func iyrValidatorFunc(value string) bool {
	iyr, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return iyr >= 2010 && iyr <= 2020
}

func eyrValidatorFunc(value string) bool {
	eyr, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return eyr >= 2020 && eyr <= 2030
}

func hgtValidatorFunc(value string) bool {
	if strings.Contains(value, "cm") {
		hgt, err := strconv.Atoi(strings.Replace(value, "cm", "", -1))
		if err != nil {
			return false
		}
		return hgt >= 150 && hgt <= 193
	}
	if strings.Contains(value, "in") {
		hgt, err := strconv.Atoi(strings.Replace(value, "in", "", -1))
		if err != nil {
			return false
		}
		return hgt >= 59 && hgt <= 76
	}

	return false
}

func hclValidatorFunc(value string) bool {
	match, err := regexp.MatchString("^#[a-fA-F0-9]{6}$", value)
	if err != nil {
		return false
	}
	return match
}

func eclValidatorFunc(value string) bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range validColors {
		if color == value {
			return true
		}
	}

	return false
}

func pidValidatorFunc(value string) bool {
	match, err := regexp.MatchString("^[0-9]{9}$", value)
	if err != nil {
		return false
	}
	return match
}
