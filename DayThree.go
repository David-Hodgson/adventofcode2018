package adventofcode2018

import (
	"fmt"
	"strconv"
	"strings"
)

type box struct {
	id     string
	left   int
	top    int
	width  int
	height int
}

func DayThreePartOne() {

	input := strings.Split(ReadFile("day3-input.txt"), "\n")

	boxList := make([]box, len(input))

	for i := 0; i < len(input); i++ {
		box := parseStringToBox(input[i])
		boxList[i] = box
	}

	fabric := createFabric(1000)
	fmt.Println(fabric)

	for i := 0; i < len(boxList); i++ {
		addBoxToFabric(fabric, boxList[i])
	}

	overlapCount := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				overlapCount++
			}
		}
	}

	fmt.Println("Overlap count: ", overlapCount)

}

func parseStringToBox(stringDescription string) box {
	idDimension := strings.Split(stringDescription, "@")
	positionSize := strings.Split(idDimension[1], ":")
	leftTop := strings.Split(positionSize[0], ",")
	left, _ := strconv.Atoi(strings.Trim(leftTop[0], " "))
	top, _ := strconv.Atoi(strings.Trim(leftTop[1], " "))
	widthHeight := strings.Split(positionSize[1], "x")
	width, _ := strconv.Atoi(strings.Trim(widthHeight[0], " "))
	height, _ := strconv.Atoi(widthHeight[1])
	box := box{}
	box.id = idDimension[0]
	box.left = left
	box.top = top
	box.width = width
	box.height = height
	return box
}

func createFabric(size int) [][]int {

	fabric := make([][]int, size)

	for i := 0; i < size; i++ {
		fabric[i] = make([]int, size)
	}

	return fabric
}

func addBoxToFabric(fabric [][]int, box box) {

	fmt.Println("Add box: ", box.id)

	for i := box.left; i < box.left+box.width; i++ {
		for j := box.top; j < box.top+box.height; j++ {
			fabric[i][j] = fabric[i][j] + 1
		}
	}
}
