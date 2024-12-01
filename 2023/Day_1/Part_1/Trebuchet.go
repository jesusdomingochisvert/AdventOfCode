package main

import (
	"bufio"
	"fmt"
	"os"
)

func extractCalibVal(line string) int {

	var firstDigit, lastDigit int

	for _, char := range line {
		if char >= '0' && char <= '9' {
			if firstDigit == 0 {
				firstDigit = int(char - '0')
			}
			lastDigit = int(char - '0')
		}
	}

	calibVal := firstDigit*10 + lastDigit

	fmt.Printf("Linea: %s, Valor de calibracion: %d\n", line, calibVal)

	return calibVal
}

func main() {

	file, err := os.Open("calibration.txt")

	if err != nil {

		fmt.Println("Error to open file: ", err)

		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {

		line := scanner.Text()

		calibVal := extractCalibVal(line)

		sum += calibVal
	}

	if err := scanner.Err(); err != nil {

		fmt.Println("Error to read file: ", err)

		return
	}

	fmt.Println("Total sum of calibration: ", sum)
}
