package day06

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var spacesRegExp = regexp.MustCompile(`\s+`)

func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}

func parseInput(input string) (*[]string, *[]string) {
  inputRows := strings.Split(input, "\n")
  numListStrRows := inputRows[0:len(inputRows) - 1]
  operators := spacesRegExp.Split(strings.Trim(inputRows[len(inputRows) - 1], " "), -1)

  return &numListStrRows, &operators
}

type solution struct {}

func (s solution) Part1(input string) string {
  numListStrRows , operators := parseInput(input)
  numListRows := make([][]int, 0, len(*numListStrRows))
  for _, numListStr := range *numListStrRows {
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

  totalSolutionsSum := 0
  for col, operator := range *operators {
    var solution int
    switch operator {
    case "+":
      solution = 0
      for _, nums := range numListRows {
        solution += nums[col]
      }
    case "*":
      solution = 1
      for _, nums := range numListRows {
        solution *= nums[col]
      }
    default:
      panic(errors.New("Unknown operator: " + operator))
    }

    totalSolutionsSum += solution
  }

  return strconv.Itoa(totalSolutionsSum)
}

func (s solution) Part2(input string) string {
  // numListStrRows , operators := parseInput(input)
  return ""
}

var Solution solution = solution{}
