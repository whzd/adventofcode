package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func testGame(line string) (power int) {
	splitLine := strings.Split(line, ":")
	// fmt.Println("Game id: ", id)
	sets := splitLine[1]
	// fmt.Println("Sets: ", sets)
	game := []map[string]int{}
	for _, subsets := range strings.Split(sets, ";") {
		// fmt.Println("Subsets: ", subsets)
		pull := map[string]int{}
		for _, reveal := range strings.Split(subsets, ",") {
			// fmt.Println("Reveals: ", reveal)
			grabs := strings.Split(reveal, " ")
			// fmt.Println("Color: ", grabs[2], "Number: ", grabs[1])
			color := grabs[2]
			number, err := strconv.Atoi(grabs[1])
			if err != nil {
				fmt.Println("Could not covert grab number string to int.")
			}
			pull[color] = number
		}
		game = append(game, pull)
	}
	// fmt.Println()
	base := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}
	for _, pull := range game {
		for key, _ := range pull {
			if pull[key] > base[key] {
				base[key] = pull[key]
			}
		}
	}
	return base["blue"] * base["red"] * base["green"]
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	fileScanner := bufio.NewScanner(file)
	var total int
	for fileScanner.Scan() {
		id := testGame(fileScanner.Text())
		if id != -1 {
			total += id
		}
	}
	fmt.Println("Total: ", total)

}
