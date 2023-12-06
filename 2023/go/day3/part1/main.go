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
	index      int
	size       int
}

func feedDataStructures(line string, lineNumber int, numbers *[]partNumber, symbols map[int][]int) {
	// fmt.Println("Line: ", lineNumber)

	specialChars := regexp.MustCompile("[\\/\\*\\%\\@\\+\\$\\-\\=\\#\\&]")
	digits := regexp.MustCompile("[0-9]+")

	indexSpCharsFound := specialChars.FindAllStringIndex(line, -1)
	numbersFound := digits.FindAllString(line, -1)
	indexNumbersFound := digits.FindAllStringIndex(line, -1)

	lstSymbols := []int{}
	for _, spChar := range indexSpCharsFound {
		lstSymbols = append(lstSymbols, spChar[0])
	}
	if len(lstSymbols) > 0 {
		symbols[lineNumber] = lstSymbols
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
			index:      indexFound[0],
			size:       len(numbersFound[i]),
		}
		*numbers = append(*numbers, pN)
	}

	// fmt.Println(numbers)
}

func calculateSumValidPartNumbers(numbers []partNumber, symbols map[int][]int) (total int) {
	total = 0
	for _, prtNmb := range numbers {
		flag := false
		line := prtNmb.lineNumber
		min := prtNmb.index - 1
		max := prtNmb.index + prtNmb.size
		if symbIndxLst, check := symbols[line-1]; check {
			for _, symbIndx := range symbIndxLst {
				if (symbIndx >= min) && (symbIndx <= max) {
					flag = true
					break
				}
			}
			if flag {
				total += prtNmb.number
				continue
			}
		}
		if symbIndxLst, check := symbols[line]; check {
			for _, symbIndx := range symbIndxLst {
				if (symbIndx >= min) && (symbIndx <= max) {
					flag = true
					break
				}
			}
			if flag {
				total += prtNmb.number
				continue
			}
		}
		if symbIndxLst, check := symbols[line+1]; check {
			for _, symbIndx := range symbIndxLst {
				if (symbIndx >= min) && (symbIndx <= max) {
					flag = true
					break
				}
			}
			if flag {
				total += prtNmb.number
				continue
			}
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
	symbols := make(map[int][]int)
	lineNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		feedDataStructures(scanner.Text(), lineNumber, &numbers, symbols)
		lineNumber += 1
	}

	total := calculateSumValidPartNumbers(numbers, symbols)
	fmt.Println(total)
}
