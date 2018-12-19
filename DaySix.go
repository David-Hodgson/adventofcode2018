package adventofcode2018

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func DaySixPartOne() {
	fmt.Println("Day 6 - Part One")

	input := strings.Split(ReadFile("day6-input.txt"), "\n")

	pointList := make([]point, len(input))
	maxX := 0
	minX := 0
	maxY := 0
	minY := 0
	for i := 0; i < len(input); i++ {
		inputParts := strings.Split(input[i], ",")
		x, _ := strconv.Atoi(inputParts[0])

		if x > maxX {
			maxX = x
		}

		if x < minX {
			minX = x
		}

		y, _ := strconv.Atoi(strings.Trim(inputParts[1], " "))

		if y > maxY {
			maxY = y
		}

		if x < minY {
			minY = y
		}
		newPoint := point{x, y}
		pointList[i] = newPoint
	}

	fmt.Println(pointList)
	fmt.Println("Max x: ", maxX)
	fmt.Println("Min x: ", minX)
	fmt.Println("Max y: ", maxY)
	fmt.Println("Min y: ", minY)

	blobSize1 := processPoints(pointList, 25, maxX, maxY)
	blobSize2 := processPoints(pointList, 50, maxX, maxY)

	maxBlobSize := 0
	for i := 0; i < len(blobSize1); i++ {
		fmt.Println("Point: ", i, " - ", blobSize1[i], " - ", blobSize2[i])

		if blobSize1[i] == blobSize2[i] {
			if blobSize1[i] > maxBlobSize {
				maxBlobSize = blobSize1[i]
			}
		}
	}

	fmt.Println("Max blobsize: ", maxBlobSize)
}

func DaySixPartTwo() {
	fmt.Println("Day 6 - Part Two")

	input := strings.Split(ReadFile("day6-input.txt"), "\n")

	pointList := make([]point, len(input))
	maxX := 0
	minX := 0
	maxY := 0
	minY := 0
	for i := 0; i < len(input); i++ {
		inputParts := strings.Split(input[i], ",")
		x, _ := strconv.Atoi(inputParts[0])

		if x > maxX {
			maxX = x
		}

		if x < minX {
			minX = x
		}

		y, _ := strconv.Atoi(strings.Trim(inputParts[1], " "))

		if y > maxY {
			maxY = y
		}

		if x < minY {
			minY = y
		}
		newPoint := point{x, y}
		pointList[i] = newPoint
	}

	fmt.Println(pointList)
	fmt.Println("Max x: ", maxX)
	fmt.Println("Min x: ", minX)
	fmt.Println("Max y: ", maxY)
	fmt.Println("Min y: ", minY)

	regionSize := processPointsToFindTotalDistances(pointList, 10000, maxX, maxY)
	fmt.Println("Region Size: ", regionSize)
}

func processPointsToFindTotalDistances(pointList []point, maxDistance int, maxX int, maxY int) int {
	grid := make([][]int, maxX+(maxDistance*2))

	regionSize := 0
	for x := 0; x < maxX+(maxDistance*2); x++ {
		grid[x] = make([]int, maxY+(maxDistance*2))
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			for pointNumber := 0; pointNumber < len(pointList); pointNumber++ {
				xDist := abs(pointList[pointNumber].x + maxDistance - x)
				yDist := abs(pointList[pointNumber].y + maxDistance - y)
				manDist := xDist + yDist
				grid[x][y] = grid[x][y] + manDist
			}

			if grid[x][y] < maxDistance {
				regionSize++
			}
		}
	}

	return regionSize
}

func processPoints(pointList []point, gridPadding int, maxX int, maxY int) []int {

	blobSizes := make([]int, len(pointList))
	grid := make([][]int, maxX+(gridPadding*2))

	for x := 0; x < maxX+(gridPadding*2); x++ {
		grid[x] = make([]int, maxY+(gridPadding*2))
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			maxDistance := maxX + maxY
			isMulti := false
			cloestPoint := -1
			//iterate the points
			for pointNumber := 0; pointNumber < len(pointList); pointNumber++ {
				xDist := abs(pointList[pointNumber].x + gridPadding - x)
				yDist := abs(pointList[pointNumber].y + gridPadding - y)
				manDist := xDist + yDist

				if manDist == maxDistance {
					isMulti = true
				} else if manDist < maxDistance {
					isMulti = false
					maxDistance = manDist
					cloestPoint = pointNumber
				}
			}
			if isMulti {
				grid[x][y] = -1
			} else {
				grid[x][y] = cloestPoint
				blobSizes[cloestPoint] = blobSizes[cloestPoint] + 1
			}
		}
	}

	return blobSizes
}
