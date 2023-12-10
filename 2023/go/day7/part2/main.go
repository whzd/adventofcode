package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type hand struct {
	cards string
	bid   int
}

var wg sync.WaitGroup

func lineToHand(line string) hand {
	tmpHand := strings.Split(line, " ")
	bid, err := strconv.Atoi(tmpHand[1])
	if err != nil {
		fmt.Printf("Unable to convert bid to int! Error: %v\n", err)
	}
	h := hand{
		cards: tmpHand[0],
		bid:   bid,
	}
	return h
}

func evaluateHand(cards string) (appear int, twoPair bool, fullHouse bool) {
	charMap := make(map[rune]int)
	countJoker := 0
	for _, char := range cards {
		if char == 'J' {
			countJoker += 1
		} else {
			charMap[char]++
		}
	}

	// if hand is only jokers
	if countJoker == 5 {
		return countJoker, false, false
	}
	pair := false
	twoPair = false
	repetition := 1
	for _, count := range charMap {
		if count > repetition {
			repetition = count
		}
		if count == 2 && !pair {
			pair = true
		} else if count == 2 && pair {
			twoPair = true
		}
	}
	//two pairs and a joker, hand is always fullhouse
	repetition += countJoker
	if twoPair && repetition == 3 {
		return repetition, false, true
	}
	// regular full house
	if pair && repetition == 3 && countJoker == 0 {
		return repetition, false, true
	}
	// every other case
	return repetition, twoPair, false
}

func indexOf(list []rune, char rune) int {
	for i, value := range list {
		if value == char {
			return i
		}
	}
	return -1
}

func handsOrdered(hand1 hand, hand2 hand) bool {
	cardOrder := []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

	for i := 0; i < len(hand1.cards); i++ {
		indexCard1 := indexOf(cardOrder, rune(hand1.cards[i]))
		indexCard2 := indexOf(cardOrder, rune(hand2.cards[i]))
		if indexCard1 == indexCard2 {
			continue
		} else if indexCard1 < indexCard2 {
			return false
		} else {
			return true
		}
	}
	return true
}

func orderHandList(handList *[]hand) {
	defer wg.Done()

	// bubble sort
	size := len(*handList)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if handsOrdered((*handList)[i], (*handList)[j]) == false {
				(*handList)[i], (*handList)[j] = (*handList)[j], (*handList)[i]
			}
		}
	}
}

func calculateWinnings(handList []hand, rank int, total chan int) {
	defer wg.Done()
	wins := 0
	for i, hand := range handList {
		multiplier := rank + i + 1
		wins += hand.bid * multiplier
	}
	total <- wins
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to open the file: %v\n", err)
	}

	scanner := bufio.NewScanner(file)

	fiveOfAKind := []hand{}
	fourOfAKind := []hand{}
	fullHouse := []hand{}
	threeOfAKind := []hand{}
	twoPair := []hand{}
	onePair := []hand{}
	highCard := []hand{}

	for scanner.Scan() {
		h := lineToHand(scanner.Text())
		appear, twoP, fullH := evaluateHand(h.cards)
		switch appear {
		case 1:
			highCard = append(highCard, h)
		case 2:
			if twoP {
				twoPair = append(twoPair, h)
			} else {
				onePair = append(onePair, h)
			}
		case 3:
			if fullH {
				fullHouse = append(fullHouse, h)
			} else {
				threeOfAKind = append(threeOfAKind, h)
			}
		case 4:
			fourOfAKind = append(fourOfAKind, h)
		case 5:
			fiveOfAKind = append(fiveOfAKind, h)
		}
	}

	wg.Add(7)

	go orderHandList(&highCard)
	go orderHandList(&onePair)
	go orderHandList(&twoPair)
	go orderHandList(&threeOfAKind)
	go orderHandList(&fullHouse)
	go orderHandList(&fourOfAKind)
	go orderHandList(&fiveOfAKind)

	wg.Wait()

	wg.Add(7)

	total := 0
	rank := 0
	highCardTotal := make(chan int, 1)
	go calculateWinnings(highCard, rank, highCardTotal)
	total += <-highCardTotal

	rank += len(highCard)
	onePairTotal := make(chan int, 1)
	go calculateWinnings(onePair, rank, onePairTotal)
	total += <-onePairTotal

	rank += len(onePair)
	twoPairTotal := make(chan int, 1)
	go calculateWinnings(twoPair, rank, twoPairTotal)
	total += <-twoPairTotal

	rank += len(twoPair)
	threeOfAKindTotal := make(chan int, 1)
	go calculateWinnings(threeOfAKind, rank, threeOfAKindTotal)
	total += <-threeOfAKindTotal

	rank += len(threeOfAKind)
	fullHouseTotal := make(chan int, 1)
	go calculateWinnings(fullHouse, rank, fullHouseTotal)
	total += <-fullHouseTotal

	rank += len(fullHouse)
	fourOfAKindTotal := make(chan int, 1)
	go calculateWinnings(fourOfAKind, rank, fourOfAKindTotal)
	total += <-fourOfAKindTotal

	rank += len(fourOfAKind)
	fiveOfAKindTotal := make(chan int, 1)
	go calculateWinnings(fiveOfAKind, rank, fiveOfAKindTotal)
	total += <-fiveOfAKindTotal

	wg.Wait()
	fmt.Println(total)

}
