package adventofcode2018

import (
	"fmt"
)

type circle struct {
	size        int
	currentItem *circleItem
}

type circleItem struct {
	marbleValue int64
	next        *circleItem
	previous    *circleItem
}

func (c *circle) moveNext() {
	c.currentItem = c.currentItem.next
}

func (c *circle) movePrevious() {
	c.currentItem = c.currentItem.previous
}

func (c *circle) add(newValue int64) {

	if c.size == 0 {
		circle1 := circleItem{}
		circle1.marbleValue = newValue
		circle1.next = &circle1
		circle1.previous = &circle1
		c.size++
		c.currentItem = &circle1
	} else {
		current := c.currentItem
		c.size++
		next := current.next

		circle1 := circleItem{}
		circle1.marbleValue = newValue
		circle1.next = next
		circle1.previous = current

		next.previous = &circle1
		current.next = &circle1

		c.currentItem = &circle1
	}
}

func (c *circle) remove() {
	previous := c.currentItem.previous
	next := c.currentItem.next
	next.previous = previous
	previous.next = next

	c.currentItem = next
}

func DayNinePartOne() {
	fmt.Println("Day 9 - Part One")

	numberOfPlayers := 419
	totalNumberOfMarbles := 72164

	playGame(numberOfPlayers, totalNumberOfMarbles)
}

func DayNinePartTwo() {
	fmt.Println("Day 9 - Part Two")

	numberOfPlayers := 419
	totalNumberOfMarbles := 72164 * 100
	playGame(numberOfPlayers, totalNumberOfMarbles)
}

func playGame(numberOfPlayers int, totalNumberOfMarbles int) {
	currentPlayer := 1

	circle := circle{}
	circle.add(0)
	playerScores := make(map[int]int64)

	for i := 1; i <= totalNumberOfMarbles; i++ {

		//Check rules
		if i%23 == 0 {
			playerScores[currentPlayer] = playerScores[currentPlayer] + int64(i)
			for x := 0; x < 7; x++ {
				circle.movePrevious()
			}
			playerScores[currentPlayer] = playerScores[currentPlayer] + circle.currentItem.marbleValue
			circle.remove()
		} else {
			//place marble in correct position
			for x := 0; x < 1; x++ {
				circle.moveNext()
			}
			circle.add(int64(i))
		}
		//play moves to next player
		currentPlayer++
		if currentPlayer > numberOfPlayers {
			currentPlayer = 1
		}
	}

	maxScore := int64(0)
	maxPlayer := 0
	for playerScore, score := range playerScores {
		if score > maxScore {
			maxScore = score
			maxPlayer = playerScore
		}
	}

	fmt.Println("Player ", maxPlayer, " wins with score ", maxScore)
}
