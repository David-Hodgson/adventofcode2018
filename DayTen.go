package adventofcode2018

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

type star struct {
	x, y   int
	dx, dy int
}

func DayTenExample() {

	fmt.Println("Day Ten - Example")

	input := strings.Split(ReadFile("day10-input.txt"), "\n")
	starList := make([]star, len(input))

	maxX := 0
	minX := 0
	minY := 0
	maxY := 0
	for i := 0; i < len(input); i++ {
		starList[i] = convertStringToStar(input[i])

		if starList[i].x < minX {
			minX = starList[i].x
		}
		if starList[i].x > maxX {
			maxX = starList[i].x
		}
		if starList[i].y < minY {
			minY = starList[i].y
		}
		if starList[i].y > maxY {
			maxY = starList[i].y
		}
	}

	fmt.Println("Stars: ", starList)
	//	width, height := (maxX - minX) , (maxY - minY)

	for i := 0; i < 1000000; i++ {
		//		createImage(width, height, i, starList)
		myMinX := 0
		myMaxX := 0
		myMinY := 0
		myMaxY := 0
		for j := 0; j < len(starList); j++ {
			currentStar := starList[j]
			currentStar.x = currentStar.x + currentStar.dx
			currentStar.y = currentStar.y + currentStar.dy
			starList[j] = currentStar

			if currentStar.x > myMaxX {
				myMaxX = currentStar.x
			}
			if currentStar.x < myMinX {
				myMinX = currentStar.x
			}
			if currentStar.y < myMinY {
				myMinY = currentStar.y
			}
			if currentStar.y > myMaxY {
				myMaxY = currentStar.y
			}
		}

		//		fmt.Println("Height: ", myMaxY-myMinY)
		if (myMaxY - myMinY) < 500 {
			width := (myMaxX - myMinX) * 10
			height := (myMaxY - myMinY) * 10
			if lookForStraightVerticalLines(starList) {
				createImage(width, height, i, starList)
			}
		}
		//		fmt.Println("minX: ", myMinX, ", maxX: ", myMaxX, ", minY: ", myMinY, ", maxY: ", myMaxY)
	}
}

func lookForStraightVerticalLines(stars []star) bool {

	for i := 0; i < len(stars); i++ {
		x := stars[i].x
		y := stars[i].y

		for j := 0; j < len(stars); j++ {
			if stars[j].x == x &&
				stars[j].y == y+1 {
				for k := 0; k < len(stars); k++ {
					if stars[k].x == x &&
						stars[k].y == y+2 {

						for l := 0; l < len(stars); l++ {

							if stars[l].x == x && stars[l].y == y+3 {
								return true
							} //l if
						} //l for
					} //k if
				} //k for
			} //j if
		} //j for
	} //i for

	return false
}
func createImage(width int, height int, imageNumber int, stars []star) {
	// Create a colored image of the given width and height.

	xOffset := width / 2
	yOffset := height / 2

	fmt.Println("Creating image of ", width, " by ", height)
	img := image.NewGray(image.Rect(0, 0, width+xOffset+20, height+yOffset+20))

	fmt.Println("xOffset: ", xOffset)
	fmt.Println("yOffset: ", yOffset)
	for i := 0; i < len(stars); i++ {
		x := stars[i].x * 10
		y := stars[i].y * 10

		x += xOffset
		y += yOffset
		for bob := x; bob < x+10; bob++ {
			for frank := y; frank < y+10; frank++ {
				img.Set(bob, frank, color.Gray{
					Y: uint8(255)})
			}
		}
	}

	fileName := "/home/david/day10/image_" + strconv.Itoa(imageNumber) + ".jpeg"
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if err := jpeg.Encode(f, img, nil); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func convertStringToStar(inputString string) star {
	star := star{}

	positionPair := strings.Split(inputString[10:strings.Index(inputString, ">")], ",")
	x, _ := strconv.Atoi(strings.Trim(positionPair[0], " "))
	y, _ := strconv.Atoi(strings.Trim(positionPair[1], " "))

	velocityPair := strings.Split(inputString[strings.Index(inputString, ">")+12:len(inputString)-1], ",")
	dx, _ := strconv.Atoi(strings.Trim(velocityPair[0], " "))
	dy, _ := strconv.Atoi(strings.Trim(velocityPair[1], " "))

	star.x = x
	star.y = y
	star.dx = dx
	star.dy = dy
	return star
}
