package adventofcode2018

import (
	"fmt"
)

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
	currentMarblePosition := 0

	circle := make([]int, 1)
	circle[0] = 0

	playerScores := make(map[int]int)

	for i := 1; i <= totalNumberOfMarbles; i++ {

		if i%10000 == 0 {
			fmt.Println("Player ", currentPlayer, " is placing marble ", i)
		}

		//fmt.Println(circle)
		//Check rules
		if i%23 == 0 {
			//fmt.Println("Special Rules")
			playerScores[currentPlayer] = playerScores[currentPlayer] + i
			removalPosition := currentMarblePosition - 7
			if removalPosition < 0 {
				removalPosition += len(circle)
			}
			//marbleScore := i + circle[removalPosition]
			//fmt.Println("Marble score: ", marbleScore)
			//fmt.Println(playerScores)
			playerScores[currentPlayer] = playerScores[currentPlayer] + circle[removalPosition]
			circle = append(circle[:removalPosition], circle[removalPosition+1:]...)
			currentMarblePosition = removalPosition
		} else {
			//place marble in correct position
			nextPosition := currentMarblePosition + 2
			if nextPosition > len(circle) {
				nextPosition = nextPosition - len(circle)
			}
			//fmt.Println("inserting at position ", nextPosition)
			circle = append(circle, 0)
			copy(circle[nextPosition:], circle[nextPosition-1:])
			circle[nextPosition] = i
			currentMarblePosition = nextPosition
		}
		//play moves to next player
		currentPlayer++
		if currentPlayer > numberOfPlayers {
			currentPlayer = 1
		}
	}

	maxScore := 0
	maxPlayer := 0
	for playerScore, score := range playerScores {
		if score > maxScore {
			maxScore = score
			maxPlayer = playerScore
		}
	}

	fmt.Println("Player ", maxPlayer, " wins with score ", maxScore)
}
