package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Set struct {
	Red   *int
	Green *int
	Blue  *int
}

type Game struct {
	Num    int
	Sets   []Set
	MinSet Set
	Power  int
}

func takeNum(s string, pos int) int {
	var num int

	_, err := fmt.Sscanf(s, "%d", &num)

	if err != nil {
		return 0
	}

	return num
}

func findColor(colors []string, color string) *int {
	pos := 0

	for _, c := range colors {

		if strings.Contains(c, color) {
			val := takeNum(c, pos)
			return &val
		}
	}

	return nil
}

func makeNewSet(set string) Set {
	colors := strings.Split(set, ", ")

	red := findColor(colors, "red")
	green := findColor(colors, "green")
	blue := findColor(colors, "blue")

	return Set{red, green, blue}
}

func addNewSetsList(setsList []string) []Set {
	var sets []Set

	for _, set := range setsList {
		sets = append(sets, makeNewSet(set))
	}

	return sets
}

func calculateMinSet(sets []Set) Set {
	var red, green, blue int

	for _, set := range sets {

		if set.Red != nil && *set.Red > red {
			red = *set.Red
		}

		if set.Green != nil && *set.Green > green {
			green = *set.Green
		}

		if set.Blue != nil && *set.Blue > blue {
			blue = *set.Blue
		}
	}

	return Set{&red, &green, &blue}
}

func calculatePower(minSet Set) int {
	red := 1
	green := 1
	blue := 1

	if minSet.Red != nil {
		red = *minSet.Red
	}

	if minSet.Green != nil {
		green = *minSet.Green
	}

	if minSet.Blue != nil {
		blue = *minSet.Blue
	}

	return red * green * blue
}

func makeNewGame(game string) Game {
	pos := 0
	gameIndex := 0
	gameNumIndex := 1
	setsIndex := 1

	twoPointsDelimiter := ": "
	blanckDelimiter := " "
	semicolonDelimiter := "; "

	splitGameByTwoPointsList := strings.Split(game, twoPointsDelimiter)
	splitGameByBlankWithGameIndexList := strings.Split(splitGameByTwoPointsList[gameIndex], blanckDelimiter)
	splitGameBySemicolonWithSetsIndexList := strings.Split(splitGameByTwoPointsList[setsIndex], semicolonDelimiter)

	num := takeNum(splitGameByBlankWithGameIndexList[gameNumIndex], pos)
	sets := addNewSetsList(splitGameBySemicolonWithSetsIndexList)
	minSet := calculateMinSet(sets)
	power := calculatePower(minSet)

	return Game{num, sets, minSet, power}
}

func main() {
	file, err := os.Open("cube_games.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var sumPowers int

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for scanner.Scan() {
		line := scanner.Text()

		sumPowers += makeNewGame(line).Power
	}

	fmt.Printf("Sum of Powers: %d\n", sumPowers)
}
