package adventofcode2018

import (
	"fmt"
	"strconv"
	"strings"
)

func DayFourteenPartOne() {
	fmt.Println("Day 14 - Part One")

	recipes := []int {3,7}

	elf1 := 0
	elf2 := 1

	experiments := 190221
	recipesCount := 10;

	for ; len(recipes) <(experiments+recipesCount); {

		//Create new recipe
		newRecipe := recipes[elf1] + recipes[elf2]

		if newRecipe > 9 {
			recipes = append(recipes, 1)
			newRecipe = newRecipe - 10
		}
		recipes = append(recipes, newRecipe)

		//Move elf's
		elf1Pos := (elf1 + 1 + recipes[elf1]) % len(recipes)
		elf2Pos := (elf2 + 1 + recipes[elf2]) % len(recipes)

		//fmt.Println("Recipes: ", recipes)
		//fmt.Println("Elf 1 pos: ", elf1Pos)
		//fmt.Println("Elf 2 pos: ", elf2Pos)
		elf1 = elf1Pos
		elf2 = elf2Pos
	}

	fmt.Println("Recipes: ", recipes[experiments:])
}


func DayFourteenPartTwo() {
	fmt.Println("Day 14 - Part Two")

	recipes := "37"

	elf1 := 0
	elf2 := 1

	//experiments := 190221
	//recipesCount := 10;

	target := "190221"
	for ; !strings.Contains(recipes,target); {
		fmt.Println(len(recipes))
		//Create new recipe
		r1,_ := strconv.Atoi(string(recipes[elf1]))
		r2,_ := strconv.Atoi(string(recipes[elf2]))

		newRecipe := r1+r2
		if newRecipe > 9 {
			recipes = recipes +  "1"
			newRecipe = newRecipe - 10
		}
		recipes = recipes + strconv.Itoa(newRecipe)

		//Move elf's
		elf1Pos := (elf1 + 1 + r1) % len(recipes)
		elf2Pos := (elf2 + 1 + r2) % len(recipes)

		//fmt.Println("Recipes: ", recipes)
		//fmt.Println("Elf 1 pos: ", elf1Pos)
		//fmt.Println("Elf 2 pos: ", elf2Pos)
		elf1 = elf1Pos
		elf2 = elf2Pos
	}

	//fmt.Println("Recipes: ", recipes[experiments:])
	fmt.Println(strings.Index(recipes,target))
}
