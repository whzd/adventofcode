package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func allValuesZero(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}
	return true
}

func calculateDiffList(values []int) []int {
	diffArray := []int{}
	for i := 0; i < len(values)-1; i++ {
		diffArray = append(diffArray, (values[i+1] - values[i]))
	}
	return diffArray
}

func predictNextValue(values []int) int {
	if allValuesZero(values) {
		return 0
	} else {
		diff := calculateDiffList(values)
		return values[len(values)-1] + predictNextValue(diff)
	}
}

func convertLineToData(line string, oasisData *[][]int) {
	tmpValues := []int{}
	for _, value := range strings.Split(line, " ") {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Unable to convert data num to int! Error: %v", err)
		}
		tmpValues = append(tmpValues, num)
	}
	*oasisData = append(*oasisData, tmpValues)
}

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Unable to open the file: %v\n", err)
	}

	scanner := bufio.NewScanner(file)
	oasisData := [][]int{}

	for scanner.Scan() {
		convertLineToData(scanner.Text(), &oasisData)
	}

	total := 0
	for _, list := range oasisData {
		total += predictNextValue(list)
	}
	fmt.Println(total)

}
