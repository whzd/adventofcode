package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type race struct {
	timeMs int
	distMm int
}

func convertLinesToRaces(lines []string) race {
	digit := regexp.MustCompile("[0-9]+")
	timeParts := digit.FindAllString(lines[0], -1)
	distanceParts := digit.FindAllString(lines[1], -1)
	var time string
	var distance string
	for i, part := range timeParts {
		time += part
		distance += distanceParts[i]
	}
	t, err := strconv.Atoi(time)
	if err != nil {
		fmt.Printf("Unable to convert time to int! Error: %v\n", err)
	}
	d, err := strconv.Atoi(distance)
	if err != nil {
		fmt.Printf("Unable to convert distance to int! Error: %v\n", err)
	}
	r := race{
		timeMs: t,
		distMm: d,
	}
	return r
}

func calculateWaysToWin(race race) int {
	lowestValue := 0
	highestValue := 0
	for i := 0; i < race.timeMs; i++ {
		tmpValue := (race.timeMs - i) * i
		if tmpValue > race.distMm {
			lowestValue = i
			break
		}
	}
	for i := race.timeMs; i > 0; i-- {
		tmpValue := (race.timeMs - i) * i
		if tmpValue > race.distMm {
			highestValue = i
			break
		}
	}
	return highestValue - lowestValue + 1
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	races := convertLinesToRaces(lines)
	fmt.Println(calculateWaysToWin(races))

}
