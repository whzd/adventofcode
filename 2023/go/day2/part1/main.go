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

func testGame(line string, limits map[string]int) (gameId int) {
	digits := regexp.MustCompile("[0-9]+")
	splitLine := strings.Split(line, ":")
	id, err := strconv.Atoi(digits.FindString(splitLine[0]))
	if err != nil {
		fmt.Println("Could not convert game id string to int.")
	}
	// fmt.Println("Game id: ", id)
	sets := splitLine[1]
	// fmt.Println("Sets: ", sets)
	pulls := []map[string]int{}
	for _, subsets := range strings.Split(sets, ";") {
		// fmt.Println("Subsets: ", subsets)
		tmpMap := map[string]int{}
		for _, reveal := range strings.Split(subsets, ",") {
			// fmt.Println("Reveals: ", reveal)
			grabs := strings.Split(reveal, " ")
			// fmt.Println("Color: ", grabs[2], "Number: ", grabs[1])
			color := grabs[2]
			number, err := strconv.Atoi(grabs[1])
			if err != nil {
				fmt.Println("Could not covert grab number string to int.")
			}
			tmpMap[color] = number
		}
		pulls = append(pulls, tmpMap)
	}
	// fmt.Println()
	for _, pull := range pulls {
		for key, _ := range pull {
			// fmt.Println("Pull key:", key, "Pull value:", value[key])
			// fmt.Println("Limit key:", key, "Limit value:", limits[key])
			// fmt.Println()
			if pull[key] > limits[key] {
				return -1
			}
		}
	}
	return id
}

func main() {

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	fileScanner := bufio.NewScanner(file)
	var total int
	for fileScanner.Scan() {
		id := testGame(fileScanner.Text(), limits)
		if id != -1 {
			total += id
		}
	}
	fmt.Println("Total: ", total)

}
