package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var numbers map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func main() {
	f, err := os.Open("day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineCoordinates := scanner.Text()

		var first, last *int
		var carryover string
		for _, c := range lineCoordinates {
			carryover += string(c)
			parsed, match, resetCarryover := numberMatch(carryover, string(c))
			if resetCarryover {
				carryover = ""
			}
			if !match {
				continue
			}
			if first == nil {
				first = &parsed
			}
			last = &parsed
		}

		if first == nil && last == nil {
			log.Fatalf("could not parse integers for line %s", lineCoordinates)
		}

		coord := (*first * 10) + *last

		log.Printf("word: %s; first: %v, last: %v, coord: %v", lineCoordinates, *first, *last, coord)

		sum += coord
	}

	log.Printf("sum: %v", sum)
}

func numberMatch(carryover string, c string) (i int, match bool, resetCarryover bool) {
	parsed, err := strconv.Atoi(c)
	if err != nil {
		return prefixMatch(carryover)
	}
	return parsed, true, true
}

func prefixMatch(carryover string) (i int, match bool, resetCarryover bool) {
	for i := 0; i < len(carryover); i++ {
		sub := carryover[i:]
		num, ok := numbers[sub]
		if ok {
			return num, true, false
		}
	}

	return -1, false, false
}
