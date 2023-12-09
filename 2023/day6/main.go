package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	races := parseInput("day6-input.txt", part1RaceParser)

	log.Println(fmt.Sprintf("part 1 result: %v", determineMarginOfError(races)))
}

func part2() {
	races := parseInput("day6-input.txt", part2RaceParser)

	log.Println(fmt.Sprintf("part 2 result: %v", determineMarginOfError(races)))
}

func determineMarginOfError(races []Race) int {
	margin := 1
	for _, race := range races {
		margin *= determineWaysToBeat(race)
	}
	return margin
}

func determineWaysToBeat(race Race) int {

	var minTime int
	for minTime = 0; minTime < race.Time; minTime++ {
		coastDistance := minTime
		coastTime := race.Time - coastDistance

		if coastDistance*coastTime > race.Distance {
			// We found our min button push time
			break
		}
	}

	var maxTime int
	for maxTime = race.Time; maxTime > 0; maxTime-- {
		coastDistance := maxTime
		coastTime := race.Time - coastDistance

		if coastDistance*coastTime > race.Distance {
			// We found our max button push time
			break
		}
	}

	// Total numbers that beat time plus 1 for inclusive
	return maxTime - minTime + 1
}

type Race struct {
	Time     int
	Distance int
}

func parseInput(filename string, raceParser func(times, distances []string) []Race) []Race {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	timesLine := scanner.Text()
	scanner.Scan()
	distanceLine := scanner.Text()

	times := strings.Fields(strings.TrimPrefix(timesLine, "Time: "))
	distances := strings.Fields(strings.TrimPrefix(distanceLine, "Distance: "))

	if len(times) != len(distances) {
		log.Fatal(fmt.Sprintf("wrong lengths, %d, %d", len(times), len(distances)))
	}

	return raceParser(times, distances)
}

func part1RaceParser(times, distances []string) []Race {
	races := make([]Race, len(times))
	for i := 0; i < len(times) && i < len(distances); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		races[i] = Race{Time: time, Distance: distance}
	}

	return races
}

func part2RaceParser(times, distances []string) []Race {
	kernedTimes := strings.Join(times, "")
	kernedDistances := strings.Join(distances, "")

	time, _ := strconv.Atoi(kernedTimes)
	distance, _ := strconv.Atoi(kernedDistances)

	return []Race{
		{
			Time:     time,
			Distance: distance,
		},
	}
}
