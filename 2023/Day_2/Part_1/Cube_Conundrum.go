package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPossible(game string, red, green, blue int) bool {
	cubeCounts := map[string]int{"red": red, "green": green, "blue": blue}

	for _, subset := range strings.Split(game, ";") {
		subset = strings.TrimSpace(subset)

		cubes := strings.Split(subset, ",")

		for _, cube := range cubes {
			parts := strings.Fields(cube)

			if len(parts) != 2 {
				continue
			}

			count, color := atoi(parts[0]), parts[1]

			if _, ok := cubeCounts[color]; !ok {
				continue
			}

			if cubeCounts[color] < count || count < 0 {
				return false
			}
		}
	}

	return true
}

func atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	file, err := os.Open("cube_games.txt")
	if err != nil {
		fmt.Println("Error abriendo el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	red, green, blue := 12, 13, 14
	sumOfIDs := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameID := atoi(strings.Split(line, ":")[0][5:])
		game := strings.Split(line, ":")[1]

		fmt.Println(line)

		if isPossible(game, red, green, blue) {
			sumOfIDs += gameID
			fmt.Println("\nJuego posible. Suma de IDs hasta ahora:", sumOfIDs)
		} else {
			fmt.Println("Juego no posible.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error leyendo el archivo:", err)
		return
	}

	fmt.Println("Suma de IDs de juegos posibles:", sumOfIDs)
}
