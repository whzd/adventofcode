package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func navigate(startingNode string, instructions string, m map[string][]string) (string, int) {
	currentNode := startingNode
	for i := 0; i < len(instructions); i++ {
		if currentNode == "ZZZ" {
			return currentNode, i + 1
		} else if instructions[i] == 'R' {
			currentNode = m[currentNode][1]
		} else if instructions[i] == 'L' {
			currentNode = m[currentNode][0]
		}
	}
	return currentNode, len(instructions)
}

func calculateHops(instructions string, m map[string][]string) int {
	total := 0
	hops := 0
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		currentNode, hops = navigate(currentNode, instructions, m)
		total += hops
	}
	return total
}

func feedNavigation(navigation map[string][]string, line string) {
	tmpLine := strings.Split(line, " = ")
	key := tmpLine[0]
	tmpElements := strings.Split(tmpLine[1], ", ")
	values := []string{tmpElements[0][1:], tmpElements[1][:3]}
	navigation[key] = values
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}

	scanner := bufio.NewScanner(file)

	var instructions string
	var navigation = map[string][]string{}
	lineCount := 0
	for scanner.Scan() {
		if lineCount == 0 {
			instructions = scanner.Text()
		}
		if lineCount > 1 {
			feedNavigation(navigation, scanner.Text())
		}
		lineCount++
	}
	fmt.Println(calculateHops(instructions, navigation))

}
