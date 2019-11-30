package adventofcode2018

import (
	"fmt"
	"strings"
	"sort"
)

const (
	Left = iota
	Right
	Up
	Down
	Straight
)

type cart struct {
	x,y,direction, intersectionDirection int
	removed bool
}

var (
	interSecs [][]int = make([][]int,4)

)

func init() {

	interSecs[Left] = make([]int,5)
	interSecs[Left][Left] = Down
	interSecs[Left][Straight] = Left
	interSecs[Left][Right] = Up

	interSecs[Right] = make([]int,5)
	interSecs[Right][Left] = Up
	interSecs[Right][Straight] = Right
	interSecs[Right][Right] = Down

	interSecs[Up] = make([]int,5)
	interSecs[Up][Left] = Left
	interSecs[Up][Straight] = Up
	interSecs[Up][Right] = Right

	interSecs[Down] = make([]int,5)
	interSecs[Down][Left] = Right
	interSecs[Down][Straight] = Down
	interSecs[Down][Right] = Left

}

func DayThirteenPartTwo() {
	fmt.Println("Day 13 - Part Two")

	input := ReadFile("day13-input.txt")

	inputRows := strings.Split(input, "\n")

	carts,tracks := parseInput(inputRows)

	for i := 0; i< 100000; i++ {

		fmt.Println("Time: ", i)

		validCarts := 0;
		for x := 0; x <len(carts); x++ {
			if (!carts[x].removed) {
				fmt.Println("x:", carts[x].x, "y:",carts[x].y)
				validCarts++
			}
		}
		if validCarts == 1 {
			break;
		}
		sort.Slice(carts, func(i,j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			} else {
				return carts[i].y < carts[j].y
			}
		})

		for j :=0; j < len(carts); j++ {
			cart := &carts[j]
			if cart.removed {
				continue
			}

			moveCart(cart,tracks)
			checkForCollisions(carts) 
		}
	}
}

func DayThirteenPartOne() {
	fmt.Println("Day 13 - Part One")

	input := ReadFile("day13-input.txt")

	inputRows := strings.Split(input, "\n")

	carts,tracks := parseInput(inputRows)


	for i := 0; i< 1000; i++ {

		fmt.Println("Time: ", i)

		sort.Slice(carts, func(i,j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			} else {
				return carts[i].y < carts[j].y
			}
		})

		for j :=0; j < len(carts); j++ {
			cart := &carts[j]

			moveCart(cart,tracks)
			checkForCollisions(carts)
		}
	}
}

func parseInput(inputRows []string) ([]cart, [][]string) {


	carts := make([]cart,0)
	tracks := make([][]string,len(inputRows))

	for row := 0; row < len(inputRows); row++ {
		tracks[row] = make([]string,len(inputRows[row]))

		for col := 0; col < len(inputRows[row]); col++ {
			bob := string(inputRows[row][col])

			if bob == ">" ||
				bob == "<" ||
				bob == "v" ||
				bob == "^" {
					newCart := cart{}
					newCart.y = row
					newCart.x = col
					newCart.intersectionDirection = Left
					newCart.removed = false

					if bob == ">" || bob == "<" {
						tracks[row][col] = "-"

						if bob == ">" {
							newCart.direction = Right
						} else {
							newCart.direction = Left
						}
					}

					if bob == "v" || bob == "^" {
						tracks[row][col] = "|"

						if bob == "^" {
							newCart.direction = Up
						} else {
							newCart.direction = Down
						}
					}

					carts = append(carts,newCart)
			} else {
				tracks[row][col] = string(inputRows[row][col])
			}
		}
	}

	fmt.Println("Carts: ", carts)

	return carts,tracks
}

func moveCart(cart *cart, tracks [][]string) {
	switch dir := cart.direction; dir {
	case Left:

		next := tracks[cart.y][cart.x-1]
		cart.x = cart.x - 1

		if next == "/" {
			cart.direction = Down
		}

		if next == "\\" {
			cart.direction = Up
		}

		if next == "+" {
			cart.direction = interSecs[dir][cart.intersectionDirection]
			cart.intersectionDirection = getNextIntersectionDirection(cart.intersectionDirection)

		}
	case Right:
		next := tracks[cart.y][cart.x+1]
		cart.x = cart.x + 1

		if next == "/" {
			cart.direction = Up
		}

		if next == "\\" {
			cart.direction = Down
		}


		if next == "+" {
			cart.direction = interSecs[dir][cart.intersectionDirection]
			cart.intersectionDirection = getNextIntersectionDirection(cart.intersectionDirection)
		}
	case Up:
		next := tracks[cart.y-1][cart.x]

		cart.y = cart.y - 1

		if next == "/" {
			cart.direction = Right
		}

		if next == "\\" {
			cart.direction = Left
		}


		if next == "+" {
			cart.direction = interSecs[dir][cart.intersectionDirection]
			cart.intersectionDirection = getNextIntersectionDirection(cart.intersectionDirection)
		}
	case Down:
		next := tracks[cart.y+1][cart.x]
		cart.y = cart.y + 1

		if next == "/" {
			cart.direction = Left
		}

		if next == "\\" {
			cart.direction = Right
		}


		if next == "+" {
			cart.direction = interSecs[dir][cart.intersectionDirection]
			cart.intersectionDirection = getNextIntersectionDirection(cart.intersectionDirection)
		}
	}
}

func checkForCollisions(carts []cart) bool {

	collision := false;
	for i :=0; i<len(carts)-1; i++ {
		if carts[i].removed {
			continue
		}
		for j := i+1; j<len(carts); j++ {
			if carts[j].removed {
				continue
			}

			if (carts[i].x == carts[j].x &&
				carts[i].y == carts[j].y) {
					collision = true
					fmt.Println("Crash at: ", carts[i].x, ",", carts[i].y)
					carts[i].removed = true
					carts[j].removed = true
			}
		}
	}

	return collision
}

func getNextIntersectionDirection(currentDirection int) int {
	switch currentDirection {
	case Left:
		return Straight
	case Straight:
		return Right
	case Right:
		return Left
	}

	return currentDirection
}
