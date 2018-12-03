package adventofcode2018

import (
	"fmt"
	"strings"
)

func hasCharacterCount(input string, count int) bool {

	letterCount  := make(map[rune]int)

	for _,letter := range input {
		letterCount[letter]++
	}

	for _,value := range letterCount{
		if (value == count) {
			return true
		}
	}
	return false
}

func getChecksum(boxList []string) int {

	twoCount := 0
	threeCount := 0
	for i:=0; i<len(boxList); i++ {
		if (hasCharacterCount(boxList[i],2)) {
			twoCount++;
		}
		if (hasCharacterCount(boxList[i],3)) {
			threeCount++;
		}
	}

	return twoCount * threeCount
}

func DayTwoPartOne() {

	fmt.Println("Day Two - Part One")

	input := ReadFile("day2-input.txt")

	boxes := strings.Split(input, "\n")
	checksum := getChecksum(boxes)

	fmt.Println("Checksum: ", checksum)
}
