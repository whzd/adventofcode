package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func incDecErrors(sequence []int) int{
  inc:=0
  dec:=0
  for i := 0; i < len(sequence) - 1; i++ {
    if sequence[i+1] - sequence[i] > 0{
      inc++
    } else {
      dec++
    }
  }
  if inc > dec{
    return dec
  } else {
    return inc
  }
}

func intervalErros(sequence []int) int {
  errors := 0
  for i := 0; i < len(sequence) -1; i++ {
    diff := sequence[i+1] - sequence[i]
    if diff < 0 {
      diff = -diff
    }
    if diff < 1 || diff > 3 {
      errors++
      if i+2 < len(sequence){
        diff := sequence[i+2] - sequence[i]
        if diff < 0 {
          diff = -diff
        }
        if diff < 1 || diff > 3 {
          errors++
        }
      }
    }
  }
  return errors
}

func countErrors(sequence []int) int {
  count := incDecErrors(sequence)
  count += intervalErros(sequence)
  fmt.Println(sequence)
  fmt.Println(count)
  return count
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

func main(){

  file, err := os.Open("./test_input")
  //file, err := os.Open("./input")
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
    errors := countErrors(sequence)
    if errors < 2 {
      total += 1
    }
  }

  fmt.Println(total)
}
