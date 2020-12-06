package day6

//FindSumOfAnswers finds sum of answers to any question in the group
func FindSumOfAnswers(inputs []string) int {
	answeredQuestions := 0
	groupAnsweredQuestions := []string{}
	for _, input := range inputs {
		if input == "" {
			answeredQuestions += len(groupAnsweredQuestions)
			groupAnsweredQuestions = []string{}
			continue
		}

		for _, question := range input {
			if contains(groupAnsweredQuestions, string(question)) {
				continue
			}
			groupAnsweredQuestions = append(groupAnsweredQuestions, string(question))
		}
	}
	answeredQuestions += len(groupAnsweredQuestions)

	return answeredQuestions
}

//FindSumOfAllAnswers finds sum of answers to every question in the group
func FindSumOfAllAnswers(inputs []string) int {
	answeredQuestions := 0
	groupSize := 0
	answers := map[string]int{}

	for _, input := range inputs {
		if input == "" {
			for _, v := range answers {
				if groupSize == v {
					answeredQuestions++
				}
			}
			answers = map[string]int{}
			groupSize = 0
			continue
		}

		for _, k := range input {
			answers[string(k)]++
		}
		groupSize++
	}

	for _, v := range answers {
		if groupSize == v {
			answeredQuestions++
		}
	}

	return answeredQuestions
}

func contains(s []string, w string) bool {
	for _, v := range s {
		if v == w {
			return true
		}
	}
	return false
}
