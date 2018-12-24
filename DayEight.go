package adventofcode2018

import (
	"fmt"
	"strconv"
	"strings"
)

var metaDataTotal = 0

func DayEightPartOne() {
	fmt.Println("Day 8 - Part One")

	//input := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	input := ReadFile("day8-input.txt")

	nodeList := strings.Split(input, " ")

	parseNodeString(nodeList)

	fmt.Println("MetaData Total: ", metaDataTotal)
}

func parseNodeString(nodeString []string) int {

	childNodeCount, _ := strconv.Atoi(nodeString[0])
	metaDataCount, _ := strconv.Atoi(nodeString[1])

	offset := 2
	if childNodeCount > 0 {
		offset += readNodes(nodeString[2:], childNodeCount)
	}

	if metaDataCount > 0 {

		for i := 0; i < metaDataCount; i++ {
			metaDataValue, _ := strconv.Atoi(nodeString[offset])
			metaDataTotal += metaDataValue
			offset++
		}

	}

	return offset
}

func readNodes(nodeString []string, nodeCount int) int {

	offset := 0
	for i := 0; i < nodeCount; i++ {

		offset += parseNodeString(nodeString[offset:])
	}

	return offset
}
