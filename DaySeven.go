package adventofcode2018

import (
	"fmt"
	"strings"
)

func DaySevenPartOne() {
	fmt.Println("Day 7 - Part One")

	input := strings.Split(ReadFile("day7-input.txt"), "\n")

	stateMap := make(map[string]map[string]bool)
	for i := 0; i < len(input); i++ {

		preStep := input[i][5:6]
		step := input[i][36:37]
		if _, exists := stateMap[step]; !exists {
			stateMap[step] = make(map[string]bool)
		}

		if _, exists := stateMap[preStep]; !exists {
			stateMap[preStep] = make(map[string]bool)
		}

		stateMap[step][preStep] = true
	}

	output := ""

	for len(stateMap) > 0 {

		startingStep := "Z"
		for step, preSteps := range stateMap {

			if len(preSteps) == 0 && step < startingStep {
				startingStep = step
			}

		}

		output += startingStep

		delete(stateMap, startingStep)
		for _, preSteps1 := range stateMap {
			delete(preSteps1, startingStep)
		}

	}

	fmt.Println("Output: ", output)
}


func DaySevenPartTwo() {
	fmt.Println("Day 7 - Part Two")

	input := strings.Split(ReadFile("day7-input.txt"), "\n")

	stateMap := make(map[string]map[string]bool)
	for i := 0; i < len(input); i++ {

		preStep := input[i][5:6]
		step := input[i][36:37]
		if _, exists := stateMap[step]; !exists {
			stateMap[step] = make(map[string]bool)
		}

		if _, exists := stateMap[preStep]; !exists {
			stateMap[preStep] = make(map[string]bool)
		}

		stateMap[step][preStep] = true
	}

	output := ""
	currentTime := 0
	maxWorkers := 5
	workers := make(map[string]int)

	for len(stateMap) > 0 {

		for state,time := range workers {
			workers[state] = time -1
			//TODO check for done states
			if workers[state] == 0 {
				delete(workers, state)
				delete(stateMap, state)
				for _,preState := range stateMap {
					delete(preState, state)
				}
				output += state
				fmt.Println("output: ", output)
			}
		}
		
		if len(workers) < maxWorkers {
			//Look for somethhing to do
		}
		currentTime++
	}

	for state, _ := range stateMap {

		bob := int(state[0]) - 64
		fmt.Println("State: ", state[0:1], ", Value: ", bob)	
		
		currentTime++
	}	

	fmt.Println("Elapsed Time: ", currentTime)
	fmt.Println("Output: ", output)
}
