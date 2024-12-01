package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){

  //file, err := os.Open("./test_input")
  file, err := os.Open("./input")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var listA, listB []int

  fscanner := bufio.NewScanner(file)

  // Create lists from file lines
  for fscanner.Scan() {
    numbers := strings.Fields(fscanner.Text())

    // Convert str to int
    valueA, err := strconv.Atoi(numbers[0])
    if err != nil {
      log.Fatal(err)
    }
    listA = append(listA, valueA)

    // Convert str to int
    valueB, err := strconv.Atoi(numbers[1])
    if err != nil {
      log.Fatal(err)
    }
    listB = append(listB, valueB)
  }

  hitmap := make(map[int]int)

  // Fill the similarity map
  for _, number := range listB {
    v, ok := hitmap[number]
    if ok {
      hitmap[number] = v + 1;
    } else {
      hitmap[number] = 1;
    }
  }

  total := 0
  // Calculate the total similarity value
  for _, number := range listA {
    if v, ok := hitmap[number]; ok {
      total += number * v;
    }
  }

  fmt.Println(total)
}

