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

type card struct {
	id             int
	nrOfWins       int
	numberOfCopies int
}

func feedData(line string, cardList *[]card) {
	matchCount := 0
	cardSplit := strings.Split(line, ":")

	digits := regexp.MustCompile("[0-9]+")
	cardId, err := strconv.Atoi(digits.FindString((cardSplit[0])))
	if err != nil {
		fmt.Printf("Error converting card id! Error: %v\n", err)
	}

	numbersSplit := strings.Split(cardSplit[1], " | ")
	for _, player := range strings.Split(numbersSplit[1], " ") {
		for _, winning := range strings.Split(numbersSplit[0], " ") {
			if player != "" && player == winning {
				matchCount += 1
			}
		}
	}

	c := card{
		id:             cardId,
		nrOfWins:       matchCount,
		numberOfCopies: 1,
	}

	*cardList = append(*cardList, c)
}

func updateNumberOfCopies(cardList *[]card) {
	for i, card := range *cardList {
		for repeat := 0; repeat < card.numberOfCopies; repeat++ {
			for j := i + 1; j <= i+card.nrOfWins; j++ {
				(*cardList)[j].numberOfCopies += 1
			}
		}
	}
}

func calculateNumberOfCards(cardList []card) int {
	total := 0
	for _, card := range cardList {
		total += card.numberOfCopies
	}
	return total
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}

	scanner := bufio.NewScanner(file)

	cardList := []card{}
	for scanner.Scan() {
		feedData(scanner.Text(), &cardList)
	}
	updateNumberOfCopies(&cardList)
	fmt.Println(calculateNumberOfCards(cardList))

}
