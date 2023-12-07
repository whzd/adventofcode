package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type relation struct {
	targetStart int
	sourceStart int
	indexRange  int
}

func feedSeeds(lines []string, seeds *[]int) {
	digit := regexp.MustCompile("[0-9]+")
	numbers := digit.FindAllString(strings.Split(lines[0], ": ")[1], -1)
	for _, number := range numbers {
		seed, err := strconv.Atoi(number)
		if err != nil {
			fmt.Printf("Unable to convert seed to int! Error: %v", err)
		}
		*seeds = append(*seeds, seed)
	}
}

func convertLinesToRelations(lines []string) []relation {
	var relationList = []relation{}
	digit := regexp.MustCompile("[0-9]+")
	for i, line := range lines {
		if i == 0 {
			continue
		}
		numbers := digit.FindAllString(line, -1)
		trgNmb, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Printf("Unable to convert target start id to int! Error: %v", err)

		}
		srcNmb, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Printf("Unable to convert source start id to int! Error: %v", err)
		}
		rng, err := strconv.Atoi(numbers[2])
		if err != nil {
			fmt.Printf("Unable to convert index range to int! Error: %v", err)
		}
		r := relation{
			targetStart: trgNmb,
			sourceStart: srcNmb,
			indexRange:  rng,
		}
		relationList = append(relationList, r)
	}
	return relationList
}

func convertSourceToTarget(lines []string, source []int, target *[]int) {
	relationList := convertLinesToRelations(lines)
	for _, sourceId := range source {
		notFound := true
		for _, relation := range relationList {
			if sourceId >= relation.sourceStart && sourceId <= relation.sourceStart+relation.indexRange {
				targetId := relation.targetStart + (sourceId - relation.sourceStart)
				*target = append(*target, targetId)
				notFound = false
				break
			}
		}
		if notFound {
			*target = append(*target, sourceId)
		}
	}
}

func getLowestValue(list []int) int {
	var min int = list[0]
	for _, value := range list {
		if min > value {
			min = value
		}
	}
	return min
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to open file: %f\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := []int{}
	soil := []int{}
	fertilizer := []int{}
	water := []int{}
	light := []int{}
	temperature := []int{}
	humidity := []int{}
	location := []int{}

	handlerCase := 0
	textBlock := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			textBlock = append(textBlock, line)
		} else {
			switch handlerCase {
			case 0:
				feedSeeds(textBlock, &seeds)
			case 1:
				convertSourceToTarget(textBlock, seeds, &soil)
			case 2:
				convertSourceToTarget(textBlock, soil, &fertilizer)
			case 3:
				convertSourceToTarget(textBlock, fertilizer, &water)
			case 4:
				convertSourceToTarget(textBlock, water, &light)
			case 5:
				convertSourceToTarget(textBlock, light, &temperature)
			case 6:
				convertSourceToTarget(textBlock, temperature, &humidity)
			}
			textBlock = nil
			handlerCase += 1
		}
	}
	if len(textBlock) > 0 && handlerCase == 7 {
		convertSourceToTarget(textBlock, humidity, &location)
		textBlock = nil
	}
	fmt.Println(getLowestValue(location))

}
