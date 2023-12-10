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
		if currentNode[2] == 'Z' {
			return currentNode, i
		} else if instructions[i] == 'R' {
			currentNode = m[currentNode][1]
		} else if instructions[i] == 'L' {
			currentNode = m[currentNode][0]
		}
	}
	return currentNode, len(instructions)
}

func calculateHops(startingNode string, instructions string, m map[string][]string) int {
	total := 0
	hops := 0
	currentNode := startingNode
	for currentNode[2] != 'Z' {
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

func getStartingNodes(m map[string][]string) []string {
	startingNodes := []string{}
	for key := range m {
		if key[2] == 'A' {
			startingNodes = append(startingNodes, key)
		}
	}
	return startingNodes
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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

	// calculate the number of hops required for each node to reach the destination, then calculate the least common multiplier for those values
	startingNodes := getStartingNodes(navigation)
	listHops := []int{}
	for _, node := range startingNodes {
		hops := calculateHops(node, instructions, navigation)
		listHops = append(listHops, hops)
	}
	fmt.Println(LCM(listHops[0], listHops[1], listHops...))

}
