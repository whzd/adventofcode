package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func validSequence(numA, numB int, inc bool) bool {
  diff := numB - numA
  if diff > 0 && inc {
    return true
  } else if diff < 0 && !inc{
    return true
  }
  return false
}

func validDifference(numA, numB int) bool {
  diff := numA - numB
  if diff < 0{
    diff = -diff
  }
  if diff >= 1 && diff <= 3{
    return true
  }
  return false
}

func safeSequence(sequence []string) bool {
  inc:= false
  for i := 0; i < len(sequence) - 1; i++ {
    valueA, err := strconv.Atoi(sequence[i])
    if err != nil {
      log.Fatal(err)
    }
    valueB, err := strconv.Atoi(sequence[i+1])
    if err != nil {
      log.Fatal(err)
    }

    if i == 0 {
      inc = (valueB - valueA) > 0
    }

    if !validSequence(valueA, valueB, inc) {
      return false
    }

    if !validDifference(valueA, valueB){
      return false
    }
  }
  return true
}

func main(){

  //file, err := os.Open("./test_input")
  file, err := os.Open("./input")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var allSequences [][]string

  fscanner := bufio.NewScanner(file)

  // Create lists from file lines
  for fscanner.Scan() {
    if len(strings.Fields(fscanner.Text())) > 0 {
      allSequences = append(allSequences, strings.Fields(fscanner.Text()))
    }
  }

  total := 0
  for _, sequence := range allSequences {
    if safeSequence(sequence){
      total += 1
    }
  }

  fmt.Println(total)
}
