package adventofcode2018

import (
	"fmt"
	"strings"
)

func DayTwelvePartOne() {
	fmt.Println("Day 12 - Part One")

	//initalState := "#..#.#..##......###...###"

	//ruleInput := strings.Split("...## => #\n..#.. => #\n.#... => #\n.#.#. => #\n.#.## => #\n.##.. => #\n.#### => #\n#.#.# => #\n#.### => #\n##.#. => #\n##.## => #\n###.. => #\n###.# => #\n####. => #", "\n")

	input := strings.Split(ReadFile("day12-input.txt"), "\n")

	//fmt.Println(input)

	initalState := input[0][len("initial state: "):]
	//fmt.Println(input[2:])
	ruleInput := input[2:]
	rules := make(map[string]byte)

	for i := 0; i < len(ruleInput); i++ {

		ruleParts := strings.Split(ruleInput[i], " => ")
		rules[ruleParts[0]] = ruleParts[1][0]

	}
	fmt.Println("Initial State: ", initalState)
	//fmt.Println("Rules: ", rules)
	fmt.Println(initalState)

	currentState := initalState
	offset := 10000
	for padding := 0; padding < offset; padding++ {
		currentState = "." + currentState + "."
	}
	for second := 1; second <= 500; second++ {
		fmt.Println(second, ":", " currentLength: ", len(currentState)+2)
		currentState = currentState // + ".."
		nextState := currentState

		//	fmt.Println("Current State: ", currentState)
		//	fmt.Println("\tOld Next State: ", nextState)
		//	fmt.Println("offset:= ", offset)

		for i := 3; i < len(currentState)-2; i++ {
			//		fmt.Println("i: ", i, " - ", currentState[i-2:i+3])
			if plant, exists := rules[currentState[i-2:i+3]]; exists {
				//fmt.Println("\tpos: ", i, " matches: ", currentState[i-2:i+3], " - output: ", string(plant))
				nextState = nextState[:i] + string(plant) + nextState[i+1:]
			} else {
				nextState = nextState[:i] + "." + nextState[i+1:]
			}

		}
		//	fmt.Println("\tNew Next State: ", nextState)
		currentState = nextState

		//fmt.Println(currentState)
		//fmt.Println("")

		plantCount := 0
		plantTotal := 0
		for pp := 0; pp<len(currentState); pp++ {
			if currentState[pp] == '#' {
				plantCount++
				plantTotal += pp - offset
			}
		}
		fmt.Println("Plant Count: ", plantCount)
		fmt.Println("Plant Total: ", plantTotal)
		//fmt.Println("")
		//	if strings.Index(currentState, "#") <= 3 {
		//		currentState = "." + currentState
		//		offset++
		//	}
	}

	fmt.Println("Offset: ", offset)

	total := 0
	minPlant := len(currentState)
	maxPlant := 0
	for i := 0; i < len(currentState); i++ {
		if currentState[i] == '#' {
			if i < minPlant {
				minPlant = i
			}
			if i > maxPlant {
				maxPlant = i
			}
			total += i - offset
		}
	}

	fmt.Println("Total: ", total)
	fmt.Println("Min Plant: ", minPlant-offset)
	fmt.Println("Max Plant: ", maxPlant-offset)
}
