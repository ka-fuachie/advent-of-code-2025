package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ka-fuachie/advent-of-code-2025/util"
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

func part1(banks []string) string {
  var totalJoltage int
  for _, bank := range banks {
    digits, err := splitStringToIntSlice(bank)
    if err != nil {
      panic(err)
    }

    left, right := 0, len(bank) - 1
    maxLeftIndex, maxRightIndex := left, right

    for {
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

      if left == right - 1 {
        break
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

func part2(banks []string) string {
  return ""
}

func main() {
  input, err := util.ReadInput(3)
  if err != nil {
    panic(err)
  }
//   input := `987654321111111
// 811111111111119
// 234234234234278
// 818181911112111`

  banks := strings.Split(input, "\n")

  fmt.Println(part1(banks))
  fmt.Println(part2(banks))
}
