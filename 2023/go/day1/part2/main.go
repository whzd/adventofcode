package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

func extractNumber(line string, numbers map[string]string) (int, error) {
	var m = map[int]string{}
	runeString := []rune(line)
	for i := 0; i < len(runeString); i++ {
		if unicode.IsDigit(runeString[i]) {
			m[i] = string(runeString[i])
		}
	}
	for key, value := range numbers {
		re := regexp.MustCompile(key)
		index := re.FindAllStringIndex(line, -1)
		if index == nil {
			continue
		}
		for _, indexes := range index {
			m[indexes[0]] = value
		}

	}
	keys := []int{}
	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return strconv.Atoi(m[keys[0]] + m[keys[len(keys)-1]])
}

func main() {

	var total int
	var numbers = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	readFile, err := os.Open("./input")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		value, err := extractNumber(fileScanner.Text(), numbers)
		if err != nil {
			log.Println("Omg, it broke")
		}
		total += value
	}
	fmt.Println(total)

}
