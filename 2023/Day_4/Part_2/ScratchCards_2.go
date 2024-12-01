package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const URL = "scratch_cards.txt"

func main() {
	data := loadData()
	if data == nil {
		fmt.Println("Error loading data.")
		return
	}

	scratchCards := proccessScratchCards(data)

	fmt.Println(scratchCards)
}

func proccessScratchCards(data []string) int {
	cardCount := make(map[string]int)
	copyCount := make(map[string]int)

	for _, line := range data {
		dataNum := strings.Split(line, ":")[1]
		dataNum = strings.TrimSpace(strings.ReplaceAll(dataNum, "  ", " "))

		winners := strings.Fields(strings.Split(dataNum, "|")[0])
		yourNumbers := strings.Fields(strings.Split(dataNum, "|")[1])

		originalCards := strings.Fields(line)[0]

		for i := 1; i <= len(winners); i++ {
			copyCard := fmt.Sprintf("%s_%d", originalCards, i)

			copyCount[copyCard]++

			for _, num := range winners {
				if contains(yourNumbers, num) {
					cardCount[copyCard]++
				}
			}
		}
	}

	totalScratchCards := 0

	for card, count := range cardCount {
		totalScratchCards += count

		fmt.Printf("Card %s has %d instances.\n", card, count)
	}

	return totalScratchCards
}

func loadData() []string {
	file, err := os.Open(URL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return list
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}
