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
	dataNum := make([]string, 0)
	sum := 0

	for _, line := range data {
		dataNum = append(dataNum, strings.Split(line, ":")[1])
	}

	for _, data := range dataNum {
		data = strings.TrimSpace(strings.ReplaceAll(data, "  ", " "))
		winners := strings.Fields(strings.Split(data, "|")[0])
		yourNumbers := strings.Fields(strings.Split(data, "|")[1])

		winners = intersection(winners, yourNumbers)

		numToPlus := 0
		for i := 0; i < len(winners); i++ {
			if numToPlus == 0 {
				numToPlus = 1
				continue
			}
			numToPlus *= 2
		}
		sum += numToPlus
	}

	fmt.Println(sum)
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

func intersection(a, b []string) []string {
	set := make(map[string]bool)
	var result []string
	for _, v := range a {
		set[v] = true
	}
	for _, v := range b {
		if set[v] {
			result = append(result, v)
		}
	}
	return result
}
