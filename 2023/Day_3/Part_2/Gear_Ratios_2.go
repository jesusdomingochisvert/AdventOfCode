package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

const FILE_PATH = "engine_diagram.txt"
const RESULT = "The total sum adjacent numbers is: %d"

const DELIMITER_N = "\n"
const DELIMITER_POINT = '.'
const DELIMITER_ASTERISK = '*'

const MIN_LIMIT_TO_FIND_NUMBER = '0'
const MAX_LIMIT_TO_FIND_NUMBER = '9'
const MAX_LIMIT_GEAR = 2
const MIN_LIMIT_NUM = 0
const OUT_INDEX = -1
const DIRECTION_INDEX = 1

func main() {
	file, err := os.ReadFile(FILE_PATH)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof(RESULT, sumGears(string(file)))
}

func sumGears(s string) int {
	sum := 0

	scorable := map[string][]int{}

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

				if found, fx, fy := hasAdjacent(lines, x, firstDigit, lastDigit, isAsterisk); found {
					scorable[fmt.Sprintf("%d, %d", fx, fy)] =
						append(scorable[fmt.Sprintf("%d, %d", fx, fy)], num)
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

			if found, fx, fy := hasAdjacent(lines, x, firstDigit, lastDigit, isAsterisk); found {
				scorable[fmt.Sprintf("%d, %d", fx, fy)] =
					append(scorable[fmt.Sprintf("%d, %d", fx, fy)], num)
			}

			firstDigit = OUT_INDEX
			lastDigit = OUT_INDEX
		}
	}

	for _, value := range scorable {
		if len(value) != MAX_LIMIT_GEAR {
			continue
		}

		sum += value[0] * value[1]
	}

	return sum
}

func hasAdjacent(engine []string, x, firstDigitY, lastDigitY int, check func(byte) bool) (bool, int, int) {
	for i := firstDigitY; i <= lastDigitY; i++ {
		if x > MIN_LIMIT_NUM {
			if check(engine[x-DIRECTION_INDEX][i]) {
				return true, x - DIRECTION_INDEX, i
			}
		}

		if x < len(engine)-DIRECTION_INDEX {
			if check(engine[x+DIRECTION_INDEX][i]) {
				return true, x + DIRECTION_INDEX, i
			}
		}
	}

	if firstDigitY > MIN_LIMIT_NUM {
		if check(engine[x][firstDigitY-DIRECTION_INDEX]) {
			return true, x, firstDigitY - DIRECTION_INDEX
		}

		if x > 0 {
			if check(engine[x-DIRECTION_INDEX][firstDigitY-DIRECTION_INDEX]) {
				return true, x - DIRECTION_INDEX, firstDigitY - DIRECTION_INDEX
			}
		}

		if x < len(engine)-DIRECTION_INDEX {
			if check(engine[x+DIRECTION_INDEX][firstDigitY-DIRECTION_INDEX]) {
				return true, x + DIRECTION_INDEX, firstDigitY - DIRECTION_INDEX
			}
		}
	}

	if lastDigitY < len(engine[x])-DIRECTION_INDEX {
		if check(engine[x][lastDigitY+DIRECTION_INDEX]) {
			return true, x, lastDigitY + DIRECTION_INDEX
		}

		if x > MIN_LIMIT_NUM {
			if check(engine[x-DIRECTION_INDEX][lastDigitY+DIRECTION_INDEX]) {
				return true, x - DIRECTION_INDEX, lastDigitY + DIRECTION_INDEX
			}
		}

		if x < len(engine)-DIRECTION_INDEX {
			if check(engine[x+DIRECTION_INDEX][lastDigitY+DIRECTION_INDEX]) {
				return true, x + DIRECTION_INDEX, lastDigitY + DIRECTION_INDEX
			}
		}
	}

	return false, OUT_INDEX, OUT_INDEX
}

func isSymbol(b byte) bool {
	return !(b >= MIN_LIMIT_TO_FIND_NUMBER && b <= MAX_LIMIT_TO_FIND_NUMBER || b == DELIMITER_POINT)
}

func isAsterisk(b byte) bool {
	return b == DELIMITER_ASTERISK
}
