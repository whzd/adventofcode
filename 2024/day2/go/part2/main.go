package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func validSkip(sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		before := sequence[:i]
		after := sequence[i+1:]

		// The make built-in function allocates and initializes an object of type
		// slice, map, or chan (only).
		//	Slice: The size specifies the length. The capacity of the slice is
		//	equal to its length. A second integer argument may be provided to
		//	specify a different capacity; it must be no smaller than the
		//	length. For example, make([]int, 0, 10) allocates an underlying array
		//	of size 10 and returns a slice of length 0 and capacity 10 that is
		//	backed by this underlying array.
		newSequence := make([]int, len(before), len(before)+len(after))
		_ = copy(newSequence, before)
		newSequence = append(newSequence, after...)

		if validIncDec(newSequence) && validInterval(newSequence) {
			return true
		}
	}
	return false
}

func validIncDec(sequence []int) bool {
	inc := sequence[1]-sequence[0] > 0
	for i := 1; i < len(sequence)-1; i++ {
		if inc && sequence[i+1]-sequence[i] < 0 {
			return false
		}
		if !inc && sequence[i+1]-sequence[i] > 0 {
			return false
		}
	}
	return true
}

func validInterval(sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		diff := absInt(sequence[i+1] - sequence[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func lineToSequence(line []string) []int {
	var res []int
	for _, v := range line {
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, number)
	}
	return res
}

func main() {

	//file, err := os.Open("./test_input")
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var allSequences [][]int

	fscanner := bufio.NewScanner(file)

	// Create lists from file lines
	for fscanner.Scan() {
		if len(strings.Fields(fscanner.Text())) > 0 {
			sequence := lineToSequence(strings.Fields(fscanner.Text()))
			allSequences = append(allSequences, sequence)
		}
	}

	total := 0
	for _, sequence := range allSequences {
		valid := validIncDec(sequence) && validInterval(sequence)
		if !valid {
			valid = validSkip(sequence)
		}
		if valid {
			total++
		}
	}

	fmt.Println(total)
}
