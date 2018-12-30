package adventofcode2018

import (
	"fmt"
)

func DayElevenPartOne() {
	fmt.Println("Day 11 - Part One")

	serialNumber := 5468
	largestTotalPower := 0
	bestX := 0
	bestY := 0

	grid := make([][]int, 300)
	for i := 0; i < 300; i++ {
		grid[i] = make([]int, 300)
	}

	for x := 1; x < 301; x++ {

		for y := 1; y < 301; y++ {
			rackId := x + 10
			powerLevel := rackId * y
			powerLevel += serialNumber
			powerLevel = powerLevel * rackId
			powerLevel = powerLevel % 1000
			powerLevel = powerLevel / 100
			powerLevel -= 5

			grid[x-1][y-1] = powerLevel
		}
	}

	for x := 0; x < 298; x++ {
		for y := 0; y < 298; y++ {
			gridTotal := grid[x][y] + grid[x][y+1] + grid[x][y+2] +
				grid[x+1][y] + grid[x][y+1] + grid[x][y+2] +
				grid[x+2][y] + grid[x+2][y+1] + grid[x+2][y+2]

			if gridTotal > largestTotalPower {
				largestTotalPower = gridTotal
				bestX = x + 1
				bestY = y + 1
			}
		}
	}

	//fmt.Println("Grid: ", grid)
	fmt.Println("Best power: ", largestTotalPower, " at x: ", bestX, ", y: ", bestY)
}

func DayElevenPartTwo() {
	fmt.Println("Day 11 - Part 2")

	serialNumber := 5468

	grid := make([][]int, 300)
	for i := 0; i < 300; i++ {
		grid[i] = make([]int, 300)
	}

	for x := 1; x < 301; x++ {

		for y := 1; y < 301; y++ {
			rackId := x + 10
			powerLevel := rackId * y
			powerLevel += serialNumber
			powerLevel = powerLevel * rackId
			powerLevel = powerLevel % 1000
			powerLevel = powerLevel / 100
			powerLevel -= 5

			grid[x-1][y-1] = powerLevel
		}
	}

	calculateGridPower(grid)
}

func calculateGridPower(grid [][]int) {

	largestTotalPower := 0
	bestX := 0
	bestY := 0
	bestSize := 0

	for squareSize := 1; squareSize <= 300; squareSize++ {
		fmt.Println("Test squares of size: ", squareSize)
		for x := 0; x <= len(grid)-squareSize; x++ {
			for y := 0; y <= len(grid)-squareSize; y++ {

				gridTotal := 0
				for i := 0; i < squareSize; i++ {
					for j := 0; j < squareSize; j++ {

						gridTotal += grid[x+i][y+j]
					}
				}

				if gridTotal > largestTotalPower {
					largestTotalPower = gridTotal
					bestX = x + 1
					bestY = y + 1
					bestSize = squareSize
				}
			}
		}

		fmt.Println("\tCurrent Best power: ", largestTotalPower, ", x: ", bestX, ", y: ", bestY, ", size: ", bestSize)
	}

	fmt.Println("Best power: ", largestTotalPower, ", x: ", bestX, ", y: ", bestY, ", size: ", bestSize)
}
