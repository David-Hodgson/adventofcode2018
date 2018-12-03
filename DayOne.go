package adventofcode2018

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateFrequency(input string) int {

	inputList := strings.Split(input, "\n")

	frequency := 0

	for i := 0; i<len(inputList); i++ {

		change, _ := strconv.Atoi(inputList[i])

		frequency = frequency + change
	}

	return frequency
}

func DayOneExample() {
	fmt.Println("Day One")
	freq := calculateFrequency("+1\n+1\n+1")
	fmt.Println("Res:", freq)
}

func DayOnePartOne() {

	fmt.Println("Day 1 - Part One")

	input := ReadFile("day1-input.txt")

	frequency := calculateFrequency(input)

	fmt.Println("Frequency: ", frequency)
}
