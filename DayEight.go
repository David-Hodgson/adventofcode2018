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

type node struct {
	childCount      int
	metadata        int
	children        []node
	metadataEntries []int
}

func DayEightPartTwo() {
	fmt.Println("Day 8 - Part Two")

	metaDataTotal = 0

	//input := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	input := ReadFile("day8-input.txt")

	nodeList := strings.Split(input, " ")

	_, rootNode := parseNodeString(nodeList)

	bob := getMetaDataTotal(rootNode)
	fmt.Println("Total: ", bob)
}

func parseNodeString(nodeString []string) (int, node) {

	node := node{}
	childNodeCount, _ := strconv.Atoi(nodeString[0])
	metaDataCount, _ := strconv.Atoi(nodeString[1])

	node.childCount = childNodeCount
	offset := 2
	if childNodeCount > 0 {
		for i := 0; i < childNodeCount; i++ {
			newOffset, childNode := parseNodeString(nodeString[offset:])
			offset += newOffset
			node.children = append(node.children, childNode)
		}
	}

	if metaDataCount > 0 {

		for i := 0; i < metaDataCount; i++ {
			metaDataValue, _ := strconv.Atoi(nodeString[offset])
			metaDataTotal += metaDataValue
			node.metadata = node.metadata + metaDataValue
			node.metadataEntries = append(node.metadataEntries, metaDataValue)
			offset++
		}

	}

	return offset, node
}

func getMetaDataTotal(root node) int {

	if root.childCount == 0 {
		return root.metadata
	} else {
		total := 0

		for i := 0; i < len(root.metadataEntries); i++ {

			metaDataEntry := root.metadataEntries[i]
			if metaDataEntry <= root.childCount {
				total += getMetaDataTotal(root.children[metaDataEntry-1])
			}
		}

		return total
	}

	return 0
}
