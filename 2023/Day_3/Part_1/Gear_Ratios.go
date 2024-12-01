package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

const FILE_PATH = "engine_diagram.txt"
const RESULT = "The total sum adjacent numbers is: %d"

const DELIMITER_N = "\n"
const DELIMITER_POINT = '.'

const MIN_LIMIT_TO_FIND_NUMBER = '0'
const MAX_LIMIT_TO_FIND_NUMBER = '9'
const MIN_LIMIT_NUM = 0

const OUT_INDEX = -1
const DIRECTION_INDEX = 1

func main() {
	file, err := os.ReadFile(FILE_PATH)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof(RESULT, sumAdjacentNumbers(string(file)))
}

func sumAdjacentNumbers(s string) int {
	sum := 0

	lines := strings.Split(s, DELIMITER_N)

	firstDigit := OUT_INDEX
	lastDigit := OUT_INDEX

	for x, line := range lines {
		for y, r := range line {
			if r >= MIN_LIMIT_TO_FIND_NUMBER && r <= MAX_LIMIT_TO_FIND_NUMBER {
				if firstDigit == OUT_INDEX {
					firstDigit = y
					lastDigit = y
				}

				lastDigit = y
			}

			if (r == DELIMITER_POINT || isSymbol(byte(r))) && firstDigit != OUT_INDEX {
				num, err := strconv.Atoi(lines[x][firstDigit : lastDigit+DIRECTION_INDEX])

				if err != nil {
					panic(err)
				}

				if hasAdjacentSymbol(lines, x, firstDigit, lastDigit) {
					sum += num
				}

				firstDigit = OUT_INDEX
				lastDigit = OUT_INDEX
			}
		}

		if firstDigit != OUT_INDEX {
			num, err := strconv.Atoi(lines[x][firstDigit : lastDigit+1])

			if err != nil {
				panic(err)
			}

			if hasAdjacentSymbol(lines, x, firstDigit, lastDigit) {
				sum += num
			}

			firstDigit = OUT_INDEX
			lastDigit = OUT_INDEX
		}
	}

	return sum
}

func hasAdjacentSymbol(engine []string, x, firstDigitY, lastDigitY int) bool {
	for i := firstDigitY; i <= lastDigitY; i++ {
		if x > MIN_LIMIT_NUM {
			if isSymbol(engine[x-DIRECTION_INDEX][i]) {
				return true
			}
		}

		if x < len(engine)-DIRECTION_INDEX {
			if isSymbol(engine[x+DIRECTION_INDEX][i]) {
				return true
			}
		}
	}

	if firstDigitY > MIN_LIMIT_NUM {
		if isSymbol(engine[x][firstDigitY-DIRECTION_INDEX]) {
			return true
		}

		if x > MIN_LIMIT_NUM {
			if isSymbol(engine[x-DIRECTION_INDEX][firstDigitY-DIRECTION_INDEX]) {
				return true
			}
		}

		if x < len(engine)-DIRECTION_INDEX {
			if isSymbol(engine[x+DIRECTION_INDEX][firstDigitY-DIRECTION_INDEX]) {
				return true
			}
		}
	}

	if lastDigitY < len(engine[x])-DIRECTION_INDEX {
		if isSymbol(engine[x][lastDigitY+DIRECTION_INDEX]) {
			return true
		}

		if x > MIN_LIMIT_NUM {
			if isSymbol(engine[x-DIRECTION_INDEX][lastDigitY+DIRECTION_INDEX]) {
				return true
			}
		}

		if x < len(engine)-DIRECTION_INDEX {
			if isSymbol(engine[x+DIRECTION_INDEX][lastDigitY+DIRECTION_INDEX]) {
				return true
			}
		}
	}

	return false
}

func isSymbol(b byte) bool {
	return !(b >= MIN_LIMIT_TO_FIND_NUMBER && b <= MAX_LIMIT_TO_FIND_NUMBER || b == DELIMITER_POINT)
}
