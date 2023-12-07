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

func convertLinesToRaces(lines []string) []race {
	races := []race{}
	digit := regexp.MustCompile("[0-9]+")
	times := digit.FindAllString(lines[0], -1)
	distaces := digit.FindAllString(lines[1], -1)
	for i, time := range times {
		t, err := strconv.Atoi(time)
		if err != nil {
			fmt.Printf("Unable to convert time to int! Error: %v\n", err)
		}
		d, err := strconv.Atoi(distaces[i])
		if err != nil {
			fmt.Printf("Unable to convert distance to int! Error: %v\n", err)
		}
		r := race{
			timeMs: t,
			distMm: d,
		}
		races = append(races, r)
	}
	return races
}

func calculateWaysToWin(races []race) int {
	total := 1
	lowestValue := 0
	highestValue := 0
	for _, race := range races {
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
		total *= highestValue - lowestValue + 1
	}
	return total
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
