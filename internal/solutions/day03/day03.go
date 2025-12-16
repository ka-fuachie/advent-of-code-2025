package day03

import (
	"strconv"
	"strings"
)

func splitStringToIntSlice(str string) ([]int, error) {
  intSlice := make([]int, len(str))

  for i := 0; i < len(str); i++ {
    digit, err := strconv.Atoi(str[i:i+1])
    if err != nil {
      return []int{}, err
    }
    intSlice[i] = digit
  }

  return intSlice, nil
}

func parseInput(input string) []string {
  return strings.Split(input, "\n")
}

type solution struct {}

func (s solution) Part1(input string) string {
  banks := parseInput(input)
  var totalJoltage int
  for _, bank := range banks {
    digits, err := splitStringToIntSlice(bank)
    if err != nil {
      panic(err)
    }

    left, right := 0, len(bank) - 1
    maxLeftIndex, maxRightIndex := left, right

    for left != right - 1 {
      if digits[maxLeftIndex] > digits[maxRightIndex] {
        right--
        if digits[right] > digits[maxRightIndex] {
          maxRightIndex = right
        }
      } else {
        left++
        if digits[left] > digits[maxLeftIndex] {
          maxLeftIndex = left
        }
      }
    }

    if digits[maxLeftIndex] < digits[maxRightIndex] && maxRightIndex != len(bank) - 1 {
      maxLeftIndex = maxRightIndex
      maxRightIndex = maxLeftIndex + 1
      for i := maxRightIndex + 1; i < len(bank); i++ {
        if digits[i] > digits[maxRightIndex] {
          maxRightIndex = i
        }
      }
    }

    joltage := 10 * digits[maxLeftIndex] + digits[maxRightIndex]
    totalJoltage += joltage
  }

  return strconv.Itoa(totalJoltage)
}

func (s solution) Part2(input string) string {
  banks := parseInput(input)
  const TOTAL_DIGITS_PER_JOLTAGE = 12

  var totalJoltage int
  for _, bank := range banks {
    digits, err := splitStringToIntSlice(bank)
    if err != nil {
      panic(err)
    }
    
    joltage := 0
    nextJthIndex := 0

    // for i := 0; i < TOTAL_DIGITS_PER_JOLTAGE; i++ {
    for i := range TOTAL_DIGITS_PER_JOLTAGE {
      maxJthIndex := nextJthIndex

      for j := nextJthIndex; j <= len(bank) - TOTAL_DIGITS_PER_JOLTAGE + i; j++ {
        if digits[j] > digits[maxJthIndex] {
          maxJthIndex = j
        }
      }

      joltage += digits[maxJthIndex] * int(math.Pow10(TOTAL_DIGITS_PER_JOLTAGE - i - 1))
      nextJthIndex = maxJthIndex + 1
    }

    totalJoltage += joltage
  }

  return strconv.Itoa(totalJoltage)
}

var Solution solution = solution{}
