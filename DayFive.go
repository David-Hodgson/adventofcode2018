package adventofcode2018

import (
	"fmt"
	"strings"
)

func DayFivePartOne() {

	input := ReadFile("day5-input.txt")

	//input = "dabAcCaCBAcCcaDA"
	output := input
	for {
		output = reduce(input)

		if len(output) == len(input) {
			break
		}

		input = output
	}

	fmt.Println("Result: ", output)
	fmt.Println("Result length: ", len(output))
}

func reduce(input string) string {

	output := ""
	for i := 0; i < len(input)-1; i++ {

		current := string(input[i])
		next := string(input[i+1])
		if current != next && (current == strings.ToUpper(next) ||
			current == strings.ToLower(next)) {
			output += input[0:i]
			output += input[i+2:]
			break

		}

	}

	if output == "" {
		output = input
	}
	return output
}
