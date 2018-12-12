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

func DayFivePartTwo() {

	input := ReadFile("day5-input.txt")

	//input = "dabAcCaCBAcCcaDA"

	minLength := len(input)
	bestChar := ""

	for i := 0; i < 26; i++ {

		lowerChar := string('a' + byte(i))
		upperChar := strings.ToUpper(lowerChar)

		trimedString := strings.Replace(input, lowerChar, "", -1)
		trimedString = strings.Replace(trimedString, upperChar, "", -1)
		output := trimedString
		for {
			output = reduce(trimedString)
			if len(output) == len(trimedString) {
				break
			}

			trimedString = output
		}

		if len(output) < minLength {
			minLength = len(output)
			bestChar = lowerChar
		}
	}

	fmt.Println("Result length: ", minLength)
	fmt.Println("Best character: ", bestChar)
}

func reduce(input string) string {

	output := make([]byte, len(input))
	newLength := 0
	for i := 0; i < len(input)-1; i++ {

		current := string(input[i])
		next := string(input[i+1])
		if current != next && (strings.ToUpper(current) == strings.ToUpper(next)) {
			//We've found a set to skip this character
			//and the next
			i++
			if i == len(input)-2 {
				output[newLength] = input[i+1]
				newLength++
			}
		} else {
			//no match add this character
			output[newLength] = input[i]
			newLength++

			if i == len(input)-2 {
				output[newLength] = input[i+1]
				newLength++
			}
		}

	}

	outputString := string(output[0:newLength])
	if outputString == "" {
		outputString = input
	}
	return outputString
}
