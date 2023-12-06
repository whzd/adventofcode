package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func cardPoints(line string) float64 {
	matches := 0
	cardNumbers := strings.Split(strings.Split(line, ": ")[1], " | ")
	winningNumbers := strings.Split(cardNumbers[0], " ")
	playerNumbers := strings.Split(cardNumbers[1], " ")
	for _, player := range playerNumbers {
		for _, winner := range winningNumbers {
			if player != "" && player == winner {
				matches += 1
				break
			}
		}
	}

	if matches == 0 {
		return 0
	}
	return math.Pow(2, float64(matches-1))
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	var total float64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += cardPoints(scanner.Text())
	}
	fmt.Println(total)

}
