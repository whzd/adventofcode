package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func longestPath(field [][]string, visited []string) {

}

func convertLinesToMatrix(line string, field *[][]string) {
	lineArray := strings.Split(line, "")
	*field = append(*field, lineArray)
}

func findStartingPosition(field [][]string) (int, int) {
	for i, lst := range field {
		for j := range lst {
			if field[i][j] == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {

	file, err := os.Open("./test_input")
	if err != nil {
		log.Fatalf("Unable to open file: %v\n", err)
	}

	scanner := bufio.NewScanner(file)

	field := [][]string{}
	for scanner.Scan() {
		convertLinesToMatrix(scanner.Text(), &field)
	}
	fmt.Println(field)
	startX, startY := findStartingPosition(field)
	fmt.Println(startX)
	fmt.Println(startY)
}
