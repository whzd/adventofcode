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

func main() {

	//file, err := os.Open("./test_input")
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	var allMuls []string

	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	for fscanner.Scan() {
		foundMuls := r.FindAllString(fscanner.Text(), -1)
		allMuls = append(allMuls, foundMuls...)
	}

	var allPairs [][]int

	r, _ = regexp.Compile("[0-9]{1,3},[0-9]{1,3}")
	for _, mul := range allMuls {
		strPairs := strings.Split(r.FindString(mul), ",")
		left, _ := strconv.Atoi(strPairs[0])
		right, _ := strconv.Atoi(strPairs[1])
		allPairs = append(allPairs, []int{left, right})
	}

	total := 0
	for _, pair := range allPairs {
		total += pair[0] * pair[1]
	}
	fmt.Println(total)
}
