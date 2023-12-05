package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("day3-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	schematic := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}

	part1sum := part1(schematic)
	part2sum := part2(schematic)

	log.Println(fmt.Sprintf("part1sum: %v", part1sum))
	log.Println(fmt.Sprintf("part2sum: %v", part2sum))

}

// Part 1

func part1(schematic [][]rune) int {
	sum := 0
	for i := 0; i < len(schematic); i++ {
		col := 0
		for j := 0; j < len(schematic[i]); j++ {
			if j < col {
				continue
			}

			_, nonnumber := strconv.Atoi(string(schematic[i][j]))
			if nonnumber != nil {
				continue
			}

			partNumber, jump := parsePartNumber(schematic, i, j)

			col = jump

			sum += partNumber
		}
	}
	return sum
}

func parsePartNumber(schematic [][]rune, i, j int) (int, int) {
	partString := ""

	numsToParse := []int{j}

	shouldInclude := false

	lastColumn := j
	for len(numsToParse) > 0 {
		col := numsToParse[0]
		if _, nonnumber := strconv.Atoi(string(schematic[i][col])); nonnumber != nil {
			if !shouldInclude {
				return 0, col
			}
			num, _ := strconv.Atoi(partString)
			return num, col
		}

		shouldInclude = shouldInclude || nearbySymbolFound(schematic, i, col)

		if col < len(schematic[i])-1 {
			numsToParse = append(numsToParse, col+1)
		}

		partString += string(schematic[i][col])
		lastColumn = col
		numsToParse = numsToParse[1:]
	}

	if !shouldInclude {
		return 0, lastColumn + 1
	}

	num, _ := strconv.Atoi(partString)
	return num, lastColumn + 1
}

func nearbySymbolFound(schematic [][]rune, i, j int) bool {
	symbolFound := false

	if i > 1 {
		symbolFound = symbolFound || isSymbol(schematic[i-1][j])

		if j > 1 {
			symbolFound = symbolFound || isSymbol(schematic[i-1][j-1])
		}

		if j < len(schematic[i])-1 {
			symbolFound = symbolFound || isSymbol(schematic[i-1][j+1])
		}
	}

	if i < len(schematic)-1 {
		symbolFound = symbolFound || isSymbol(schematic[i+1][j])

		if j > 1 {
			symbolFound = symbolFound || isSymbol(schematic[i+1][j-1])
		}

		if j < len(schematic[i])-1 {
			symbolFound = symbolFound || isSymbol(schematic[i+1][j+1])
		}
	}

	if j > 1 {
		symbolFound = symbolFound || isSymbol(schematic[i][j-1])
	}

	if j < len(schematic[i])-1 {
		symbolFound = symbolFound || isSymbol(schematic[i][j+1])
	}

	return symbolFound
}

func isSymbol(char rune) bool {
	_, nonnumber := strconv.Atoi(string(char))
	if nonnumber != nil && string(char) != "." {
		return true
	}

	return false
}

// Part 2

type point struct {
	row, col int
}

func part2(schematic [][]rune) int {
	sum := 0

	stars := map[point]string{}
	numbers := map[point]int{}
	for row, line := range schematic {
		for col, char := range line {
			if isStar(char) {
				stars[point{row: row, col: col}] = "*"
			}

			if _, ok := numbers[point{row: row, col: col}]; isDigit(char) && !ok {
				buildNumber(schematic, numbers, row, col)
			}
		}
	}

	for star := range stars {
		neighbors := map[int]struct{}{}
		for location, num := range numbers {
			if star.row >= location.row-1 && star.row <= location.row+1 && star.col >= location.col-1 && star.col <= location.col+1 {
				neighbors[num] = struct{}{}
			}
		}

		if len(neighbors) > 2 {
			log.Fatalf("found too many neighbors: %v", neighbors)
		}

		if len(neighbors) == 2 {
			log.Println(fmt.Sprintf("found neighbors: %v", neighbors))
			product := 1
			for num := range neighbors {
				product *= num
			}
			sum += product
		}
	}

	return sum
}

func buildNumber(schematic [][]rune, numbers map[point]int, row, col int) {
	expansion := col
	partNumberBuilder := ""
	for expansion < len(schematic[row]) && isDigit(schematic[row][expansion]) {
		partNumberBuilder += string(schematic[row][expansion])
		expansion++
	}

	partNumber, _ := strconv.Atoi(partNumberBuilder)
	for i := col; i < expansion; i++ {
		numbers[point{row: row, col: i}] = partNumber
	}
}

func isDigit(char rune) bool {
	return char <= '9' && char >= '0'
}

func isStar(char rune) bool {
	return char == '*'
}
