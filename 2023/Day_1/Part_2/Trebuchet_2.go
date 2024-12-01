package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const FILE_PATH = "calibration.txt"
const RESULT = "The sum of all calibrtion values is:"
const FIRST_CODE_NUMBER = 0

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var NumbersRegex = regexp.MustCompile("\\D")

func getLineNumbers(line string) string {

	for word, digit := range wordToDigit {

		newWord := string(word[FIRST_CODE_NUMBER]) + fmt.Sprint(digit) + string(word[len(word)-1])

		line = strings.ReplaceAll(line, word, newWord)
	}

	return NumbersRegex.ReplaceAllString(line, "")
}

func sumTotalCalibrationValues(file string) int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	totalSum := 0

	for _, line := range lines {
		numbers := getLineNumbers(line)
		firstDigit := string(numbers[FIRST_CODE_NUMBER])
		lastDigit := numbers[len(numbers)-1:]

		finalNumberCode := firstDigit + lastDigit
		finalNumber, err := strconv.Atoi(finalNumberCode)
		if err != nil {
			panic(err)
		}
		totalSum += finalNumber
	}

	return totalSum
}

func main() {

	totalSum := sumTotalCalibrationValues(FILE_PATH)

	fmt.Println(RESULT, totalSum)
}
