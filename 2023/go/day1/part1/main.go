package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func extractNumber(line string) (int, error) {
	var firstDigit, lastDigit string
	runeString := []rune(line)
	for i := 0; i < len(runeString); i++ {
		if unicode.IsDigit(runeString[i]) {
			firstDigit = string(runeString[i])
			break
		}
	}
	for i := len(runeString) - 1; i >= 0; i-- {
		if unicode.IsDigit(runeString[i]) {
			lastDigit = string(runeString[i])
			break
		}
	}
	return strconv.Atoi(firstDigit + lastDigit)
}

func main() {

	var total int
	readFile, err := os.Open("./input")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		value, err := extractNumber(fileScanner.Text())
		if err != nil {
			log.Println("Omg, it broke")
		}
		total += value
	}
	fmt.Println(total)

}
