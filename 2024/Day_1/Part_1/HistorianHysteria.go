package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func quickSort(arr []int, low, high int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	quickSortHelper(sortedArr, low, high)
	return sortedArr
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func calculateTotalDistance(leftList, rightList []int) (int, error) {
	if len(leftList) != len(rightList) {
		return 0, fmt.Errorf("left and right lists are not equal")
	}

	sortedLeft := quickSort(leftList, 0, len(leftList)-1)
	sortedRight := quickSort(rightList, 0, len(rightList)-1)

	totalDistance := 0
	for i := 0; i < len(sortedLeft); i++ {
		distance := abs(sortedLeft[i] - sortedRight[i])
		totalDistance += distance
	}

	return totalDistance, nil
}

func readInput(filename string) ([]int, []int, error) {
	var leftList []int
	var rightList []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("input invalid in line: %s", line)
		}

		leftNum, err1 := strconv.Atoi(fields[0])
		rightNum, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("invalid numbers in line: %s", fields[0])
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning input: %v", err)
	}

	if len(leftList) == 0 || len(rightList) == 0 {
		return nil, nil, fmt.Errorf("invalid input, no values found")
	}

	return leftList, rightList, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: program historian_hysteria.txt")
		return
	}
	filename := os.Args[1]

	leftList, rightList, err := readInput(filename)
	if err != nil {
		fmt.Println("Error to read input:", err)
		return
	}

	totalDistance, err := calculateTotalDistance(leftList, rightList)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Distance total is:", totalDistance)
}
