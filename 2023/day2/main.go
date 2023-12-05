package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Num   int
	Pulls []pull
}

type pull struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	f, err := os.Open("day2-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	gameSum := 0
	gamePowers := 0
	for scanner.Scan() {
		line := scanner.Text()

		game, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}

		// Part 1
		if gameIsPossible(game, 12, 13, 14) {
			gameSum += game.Num
			// b, _ := json.Marshal(game)
			// log.Println(fmt.Sprintf("found possible game!: %v", string(b)))
		}

		// Part 2
		power := findGamePower(game)
		gamePowers += power
	}

	log.Println(fmt.Sprintf("sum of all games: %v", gameSum))
	log.Println(fmt.Sprintf("power of all games: %v", gamePowers))
}

func gameIsPossible(game Game, red, green, blue int) bool {
	for _, pull := range game.Pulls {
		if pull.Red > red || pull.Green > green || pull.Blue > blue {
			return false
		}
	}

	return true
}

func findGamePower(game Game) int {
	red := 0
	green := 0
	blue := 0

	for _, pull := range game.Pulls {
		red = max(red, pull.Red)
		green = max(green, pull.Green)
		blue = max(blue, pull.Blue)
	}

	power := red * green * blue
	b, _ := json.Marshal(game)
	log.Println(fmt.Sprintf("game %v power: %v. game: %v", game.Num, power, string(b)))

	return power
}

func parseLine(line string) (Game, error) {
	// Splitting the string to get the game number and attempts
	splitString := strings.Split(line, ": ")
	gameNumber, err := strconv.Atoi(strings.Fields(splitString[0])[1]) // Extracting the game number
	if err != nil {
		return Game{}, err
	}

	attempts := strings.Split(splitString[1], ";") // Splitting attempts
	numAttempts := len(attempts)                   // Counting the number of attempts

	parsedAttempts := make([]pull, numAttempts)

	// Parsing each attempt to associate numbers with colors
	for i, attempt := range attempts {
		colors := strings.Split(strings.TrimSpace(attempt), ", ")
		parsedAttempt := pull{}

		for _, color := range colors {
			splitColor := strings.Fields(color)
			num, _ := strconv.Atoi(splitColor[0])
			colorName := splitColor[1]
			switch colorName {
			case "red":
				parsedAttempt.Red = num
			case "green":
				parsedAttempt.Green = num
			case "blue":
				parsedAttempt.Blue = num
			default:
				return Game{}, fmt.Errorf("unknown color: %s, line: %s", colorName, line)
			}
		}

		parsedAttempts[i] = parsedAttempt
	}

	return Game{
		Num:   gameNumber,
		Pulls: parsedAttempts,
	}, nil
}
