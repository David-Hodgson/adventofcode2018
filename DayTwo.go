package adventofcode2018

import (
	"fmt"
	"strings"
)

func compareStrings(input1, input2 string) bool {

	diffCount := 0

	for index,_ := range input1 {
		if (input1[index] != input2[index]) {
			diffCount++
		}

		if (diffCount > 1) {
			return false
		}
	}
	
	return diffCount == 1 
}

func findBoxHash(boxList []string) string {
	for i:=0;i<len(boxList);i++ {
		for j:=i;j<len(boxList);j++ {
			if (compareStrings(boxList[i],boxList[j])) {
				fmt.Println("Box 1: ", boxList[i])
				fmt.Println("Box 2: ", boxList[j])
				return boxList[i]
			}
		}
	}

	return "";
}

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

func DayTwoPartTwo() {
	fmt.Println("Day Two - Part Two")

	input := ReadFile("day2-input.txt")
	boxes := strings.Split(input, "\n")

	hash := findBoxHash(boxes)

	fmt.Println("Box Hash: ", hash)
}		
