package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/ka-fuachie/advent-of-code-2025/util"
)

func parseRange(rangeStr string) ([]int, error) {
  rangeStrList := strings.Split(rangeStr, "-")
  if len(rangeStrList) != 2 {
    return []int{-1, -1}, errors.New("Expected two numbers as range, got " + rangeStr)
  }

  start, err := strconv.Atoi(string(rangeStrList[0]))
  if err != nil {
    return []int{-1, -1}, err
  }

  end, err := strconv.Atoi(string(rangeStrList[1]))
  if err != nil {
    return []int{-1, -1}, err
  }

  return []int{start, end}, nil
}

func getNumsWithTwiceRepeatedDigitsSequenceWithinRange(rangeVal []int) []int {
  // Brute force
  var nums []int
  for i := rangeVal[0]; i <= rangeVal[1]; i++ {
    numStr := strconv.Itoa(i)
    if len(numStr) % 2 != 0 {
      continue
    }

    if numStr[0:len(numStr)/2] != numStr[len(numStr)/2:] {
      continue
    }

    nums = append(nums, i)
  }
  return nums
}

func getNumsWithRepeatedDigitsSequenceWithinRange(rangeVal []int) []int {
  // Brute force
  var nums []int
  for i := rangeVal[0]; i <= rangeVal[1]; i++ {
    numStr := strconv.Itoa(i)

    isRepeating := false
    IntervalLoop:
    for j := 1; j <= len(numStr)/2; j++ {
      if j != 1 && len(numStr) % j != 0 {
        continue
      }

      intervalDigits,isRepeatingForInterval := numStr[0:j], false
      for k := 0; k < len(numStr); k+=j {
        if numStr[k:k+j] == intervalDigits  {
          isRepeatingForInterval = true
          continue
        }

        isRepeatingForInterval = false
        break
      }

      if isRepeatingForInterval {
        isRepeating = true
        break IntervalLoop
      }
    }

    if !isRepeating {
      continue
    }

    nums = append(nums, i)
  }
  return nums
}

func part1(rangeVals [][]int) string {
  totalInvalidIdsSum := 0
  for _, rangeVal := range rangeVals {
    for _, invalidId := range getNumsWithTwiceRepeatedDigitsSequenceWithinRange(rangeVal) {
      totalInvalidIdsSum += invalidId
    }
  }
  return strconv.Itoa(totalInvalidIdsSum)
}

func part2(rangeVals [][]int) string {
  totalInvalidIdsSum := 0
  for _, rangeVal := range rangeVals {
    for _, invalidId := range getNumsWithRepeatedDigitsSequenceWithinRange(rangeVal) {
      totalInvalidIdsSum += invalidId
    }
  }
  return strconv.Itoa(totalInvalidIdsSum)
}

func main() {
  input, err := util.ReadInput(2)
  if err != nil {
    panic(err)
  }
  // input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

  rangeStrs := strings.Split(input, ",")
  rangeVals := make([][]int, len(rangeStrs))

  for i, rangeStr := range rangeStrs {
    rangeVal, err := parseRange(rangeStr)
    if err != nil {
      panic(err)
    }

    rangeVals[i] = rangeVal
  }

  fmt.Println(part1(rangeVals))
  fmt.Println(part2(rangeVals))
}
