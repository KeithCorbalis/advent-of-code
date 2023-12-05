package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	Num         int
	WinningNums map[int]struct{}
	PlayerNums  map[int]struct{}
}

func main() {

	p1 := time.Now()
	part1()
	log.Println(fmt.Sprintf("part 1: %v", time.Since(p1)))

	p2 := time.Now()
	part2()
	log.Println(fmt.Sprintf("part 2: %v", time.Since(p2)))
}

func part1() {
	f, err := os.Open("day4-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		card := parseCard(line)

		winningPoints, _ := evaluateCard(card)
		sum += winningPoints
	}

	// log.Println(fmt.Sprintf("part 1: sum of cards: %v", sum))
}

func part2() {
	f, err := os.Open("day4-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	sum := 0
	iterator := 0
	cards := []Card{}
	copies := []int{1}
	for scanner.Scan() {
		line := scanner.Text()

		card := parseCard(line)
		cards = append(cards, card)
		_, matchingCount := evaluateCard(card)

		for i := matchingCount; i > 0; i-- {
			foundCopyIndex := iterator + 1 + (matchingCount - i)
			if foundCopyIndex >= len(copies) {
				for j := foundCopyIndex; j >= len(copies)-1; j-- {
					copies = append(copies, 1)
				}
			}
			copies[foundCopyIndex] += copies[iterator]
		}

		if iterator >= len(copies) {
			copies = append(copies, 1)
		}
		sum += copies[iterator]
		iterator++
	}

	// log.Println(fmt.Sprintf("part 2: sum of cards won: %v", sum))
}

func parseCard(line string) Card {
	cardDescriptorSeparator := strings.Split(line, ": ")

	num, err := strconv.Atoi(strings.Fields(cardDescriptorSeparator[0])[1])
	if err != nil {
		log.Fatal(err)
	}

	cardValueSeparator := strings.Split(cardDescriptorSeparator[1], "|")

	winningNums := loadNumbers(strings.Fields(cardValueSeparator[0]))
	playerNums := loadNumbers(strings.Fields(cardValueSeparator[1]))

	return Card{
		Num:         num,
		WinningNums: winningNums,
		PlayerNums:  playerNums,
	}
}

func evaluateCard(card Card) (total int, count int) {
	for num := range card.PlayerNums {
		_, ok := card.WinningNums[num]
		if ok {
			if count == 0 {
				total = 1
			} else {
				total *= 2
			}
			count++
		}
	}

	return total, count
}

func loadNumbers(vals []string) map[int]struct{} {
	nums := map[int]struct{}{}

	for _, s := range vals {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		nums[num] = struct{}{}
	}
	return nums
}
