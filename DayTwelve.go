package adventofcode2018

import (
	"fmt"
	"strings"
)

func DayTwelvePartOne() {
	fmt.Println("Day 12 - Part One")

	initalState := "#..#.#..##......###...###"

	ruleInput := strings.Split("...## => #\n..#.. => #\n.#... => #\n.#.#. => #\n.#.## => #\n.##.. => #\n.#### => #\n#.#.# => #\n#.### => #\n##.#. => #\n##.## => #\n###.. => #\n###.# => #\n####. => #", "\n")

	rules := make(map[string]byte)

	for i := 0; i<len(ruleInput); i++ {

		ruleParts := strings.Split(ruleInput[i], " => ")
		rules[ruleParts[0]] = ruleParts[1][0]

	}
	fmt.Println("Initial State: ", initalState)
	fmt.Println("Rules: ", rules)

	currentState := "..." + initalState
	offset := 3
	for second :=1; second <= 20; second++ {
		fmt.Println(second, ":")
		currentState = currentState + ".."
		nextState := currentState

	//	fmt.Println("Current State: ", currentState)
		fmt.Println("\tOld Next State: ", nextState)
//		fmt.Println("offset:= ", offset)

		for i :=3 ; i < len(currentState)-2; i++ {
			//fmt.Println("i: ", i, " - ", currentState[i-2:i+3])
			if plant,exists := rules[currentState[i-2:i+3]]; exists{
				//fmt.Println("pos: ", i, " matches: ", currentState[i-2:i+3])
				nextState = nextState[:i] + string(plant) + nextState[i+1:]
			} else {
				nextState = nextState[:i] + "." + nextState[i+1:]
			}
		}
		fmt.Println("\tNew Next State: ", nextState)
		currentState = nextState

		if strings.Index(currentState, "#") <= 3 {
			currentState = "." + currentState
			offset++
		}
	}

	fmt.Println("Offset: ", offset)

	total :=0
	for i := 0; i<len(currentState); i++ {
		if (currentState[i] == '#') {
			total += i-offset
		}
	}

	fmt.Println("Total: ", total)
}
