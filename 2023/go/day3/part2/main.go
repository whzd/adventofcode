package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type partNumber struct {
	lineNumber int
	number     int
	minLimit   int
	maxLimit   int
}

type gear struct {
	lineNumber int
	index      int
}

func feedDataStructures(line string, lineNumber int, numbers *[]partNumber, symbols *[]gear) {
	// fmt.Println("Line: ", lineNumber)

	specialChars := regexp.MustCompile("[\\*]")
	digits := regexp.MustCompile("[0-9]+")

	indexSpCharsFound := specialChars.FindAllStringIndex(line, -1)
	numbersFound := digits.FindAllString(line, -1)
	indexNumbersFound := digits.FindAllStringIndex(line, -1)

	for _, spChar := range indexSpCharsFound {
		g := gear{
			lineNumber: lineNumber,
			index:      spChar[0],
		}
		*symbols = append(*symbols, g)
	}
	// fmt.Println(symbols)

	for i, indexFound := range indexNumbersFound {
		value, err := strconv.Atoi(numbersFound[i])
		if err != nil {
			fmt.Printf("Could not convert value to int. Error: %v", err)
		}
		pN := partNumber{
			lineNumber: lineNumber,
			number:     value,
			minLimit:   indexFound[0] - 1,
			maxLimit:   indexFound[0] + len(numbersFound[i]),
		}
		*numbers = append(*numbers, pN)
	}

	// fmt.Println(numbers)
}

func calculateSumGearRatio(numbers []partNumber, symbols []gear) (total int) {
	total = 0
	for _, gear := range symbols {
		ratio := 1
		rationCount := 0
		for _, prtNmb := range numbers {
			if (prtNmb.lineNumber == gear.lineNumber-1) && (gear.index >= prtNmb.minLimit) && (gear.index <= prtNmb.maxLimit) {
				ratio *= prtNmb.number
				rationCount += 1
				continue
			}
			if (prtNmb.lineNumber == gear.lineNumber) && (gear.index >= prtNmb.minLimit) && (gear.index <= prtNmb.maxLimit) {
				ratio *= prtNmb.number
				rationCount += 1
				continue
			}
			if (prtNmb.lineNumber == gear.lineNumber+1) && (gear.index >= prtNmb.minLimit) && (gear.index <= prtNmb.maxLimit) {
				ratio *= prtNmb.number
				rationCount += 1
				continue
			}
		}
		if (ratio > 1) && (rationCount == 2) {
			total += ratio
		}
	}

	return total
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	numbers := []partNumber{}
	symbols := []gear{}
	lineNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		feedDataStructures(scanner.Text(), lineNumber, &numbers, &symbols)
		lineNumber += 1
	}

	total := calculateSumGearRatio(numbers, symbols)
	fmt.Println(total)
}
