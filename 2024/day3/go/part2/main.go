package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

	total := 0
	enable := true

	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	rNum, _ := regexp.Compile("[0-9]{1,3},[0-9]{1,3}")
	rDo, _ := regexp.Compile("do\\(\\)")
	rDont, _ := regexp.Compile("don't\\(\\)")
	for fscanner.Scan() {
		line := make(map[int]string)

		foundMuls := r.FindAllString(fscanner.Text(), -1)
		foundMulsIndex := r.FindAllStringIndex(fscanner.Text(), -1)
		foundDos := rDo.FindAllStringIndex(fscanner.Text(), -1)
		foundDonts := rDont.FindAllStringIndex(fscanner.Text(), -1)

		for i, mul := range foundMuls {
			line[foundMulsIndex[i][0]] = mul
		}
		for _, v := range foundDos {
			line[v[0]] = "do"
		}
		for _, v := range foundDonts {
			line[v[0]] = "dont"
		}

		keys := make([]int, 0, len(line))
		for k, _ := range line {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			if enable && line[k] == "dont" {
				enable = false
			} else if !enable && line[k] == "do" {
				enable = true
			} else if enable && line[k] != "do" {
				strPairs := strings.Split(rNum.FindString(line[k]), ",")
				left, _ := strconv.Atoi(strPairs[0])
				right, _ := strconv.Atoi(strPairs[1])
				total += left * right
			}
		}
	}

	fmt.Println(total)

}
