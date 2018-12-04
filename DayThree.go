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

	fmt.Println(boxList)
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
