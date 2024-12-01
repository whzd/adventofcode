package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

  // Sort lists
  slices.Sort(listA)
  slices.Sort(listB)

  total := 0
  for i, v := range listA {
    diff := v - listB[i]
    // Add the absolute diff
    if diff < 0{
      total += -diff
    } else {
      total += diff
    }
  }
  fmt.Println(total)

}

