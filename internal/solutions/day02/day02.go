package day02

import (
	"errors"
	"strconv"
	"strings"
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

func parseInput(input string) [][]int {
  rangeStrs := strings.Split(input, ",")
  rangeVals := make([][]int, len(rangeStrs))

  for i, rangeStr := range rangeStrs {
    rangeVal, err := parseRange(rangeStr)
    if err != nil {
      panic(err)
    }

    rangeVals[i] = rangeVal
  }

  return rangeVals
}

type solution struct {}

func (s solution) Part1(input string) string {
  rangeVals := parseInput(input)
  totalInvalidIdsSum := 0
  for _, rangeVal := range rangeVals {
    for _, invalidId := range getNumsWithTwiceRepeatedDigitsSequenceWithinRange(rangeVal) {
      totalInvalidIdsSum += invalidId
    }
  }
  return strconv.Itoa(totalInvalidIdsSum)
}

func (s solution) Part2(input string) string {
  rangeVals := parseInput(input)
  totalInvalidIdsSum := 0
  for _, rangeVal := range rangeVals {
    for _, invalidId := range getNumsWithRepeatedDigitsSequenceWithinRange(rangeVal) {
      totalInvalidIdsSum += invalidId
    }
  }
  return strconv.Itoa(totalInvalidIdsSum)
}

var Solution solution = solution{}
