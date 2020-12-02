package day2

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// PasswordPolicy struct for password policy
type PasswordPolicy struct {
	minLength    int
	maxLength    int
	requiredChar string
}

//Password struct for password value and password policy
type Password struct {
	pp    PasswordPolicy
	value string
}

//ValidatorFunc function for validate password
type ValidatorFunc func(Password) bool

//GetValidPasswordsNumber gets number of valid password based on validator function
func GetValidPasswordsNumber(passwords []Password, validate ValidatorFunc) int {
	validPasswords := []Password{}
	for _, password := range passwords {
		if isValid := validate(password); !isValid {
			continue
		}
		validPasswords = append(validPasswords, password)
	}
	return len(validPasswords)
}

//ValidateOldCompanyPassword validates password for old company
func ValidateOldCompanyPassword(pass Password) bool {
	counter := 0
	for i := 0; i < len(pass.value); i++ {
		if string(pass.value[i]) == pass.pp.requiredChar {
			counter++
		}
	}
	return counter <= pass.pp.maxLength && counter >= pass.pp.minLength
}

//ValidateCurrenctCompanyPassword validates password for current company
func ValidateCurrenctCompanyPassword(pass Password) bool {
	firstIdx := pass.pp.minLength - 1
	secondIdx := pass.pp.maxLength - 1

	if len(pass.value) <= secondIdx {
		return string(pass.value[firstIdx]) == pass.pp.requiredChar
	}

	if string(pass.value[firstIdx]) == pass.pp.requiredChar && string(pass.value[secondIdx]) == pass.pp.requiredChar {
		return false
	}

	return string(pass.value[firstIdx]) == pass.pp.requiredChar || string(pass.value[secondIdx]) == pass.pp.requiredChar
}

//ConvertToSlice converting file to input slice
func ConvertToSlice(filename string) []Password {
	passwords := []Password{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	for _, input := range strings.Split(string(data), "\n") {
		values := strings.Split(input, ":")
		passwordPolicy := values[0]
		password := strings.TrimSpace(values[1])

		newPassword := Password{
			pp:    parseToPasswordPolicy(passwordPolicy),
			value: password,
		}

		passwords = append(passwords, newPassword)
	}
	return passwords
}

func parseToPasswordPolicy(pps string) PasswordPolicy {
	values := strings.Split(pps, " ")
	length := values[0]
	requiredChar := values[1]
	minMax := strings.Split(length, "-")

	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		panic(err)
	}

	return PasswordPolicy{
		minLength:    min,
		maxLength:    max,
		requiredChar: requiredChar,
	}
}
