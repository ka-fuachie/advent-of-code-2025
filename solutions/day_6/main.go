package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ka-fuachie/advent-of-code-2025/util"
)

func part1(numRows *[][]int, operators *[]string) string {
  totalSolutionsSum := 0
  for col, operator := range *operators {
    var solution int
    switch operator {
    case "+":
      solution = 0
      for _, nums := range *numRows {
        solution += nums[col]
      }
    case "*":
      solution = 1
      for _, nums := range *numRows {
        solution *= nums[col]
      }
    default:
      panic(errors.New("Unknown operator: " + operator))
    }

    totalSolutionsSum += solution
  }

  return strconv.Itoa(totalSolutionsSum)
}

func part2(numRows *[][]int, operators *[]string) string {
  return ""
}

func main() {
  input, err := util.ReadInput(6)
  if err != nil {
    panic(err)
  }
//   input := `123 328  51 64 
//  45 64  387 23 
//   6 98  215 314
// *   +   *   + `

  inputRows := strings.Split(input, "\n")

  spacesRegExp := regexp.MustCompile(`\s+`)

  numListStrRows := inputRows[0:len(inputRows) - 1]
  numListRows := make([][]int, 0, len(numListStrRows))
  for _, numListStr := range numListStrRows {
    numStrList := spacesRegExp.Split(strings.Trim(numListStr, " "), -1)
    numList := make([]int, 0, len(numStrList))

    for _, numStr := range numStrList {
      num, err := strconv.Atoi(numStr)
      if err != nil {
        panic(err)
      }
      numList = append(numList, num)
    }

    numListRows = append(numListRows, numList)
  }

  operators := spacesRegExp.Split(strings.Trim(inputRows[len(inputRows) - 1], " "), -1)

  fmt.Println(part1(&numListRows, &operators))
  fmt.Println(part2(&numListRows, &operators))
}
