package day01

import (
	"errors"
	"strconv"
	"strings"
)

const MAX_DIAL_PTR = 99

func mod(dividend int, divisor int) int {
  return ((dividend % divisor) + divisor) % divisor
}

func abs(num int) int {
  if num < 0 {
    return -num
  }
  return num
}

func getRotationAmount(rotation string) (int, error) {
  direction, stringAmount := string(rotation[0]), string(rotation[1:])
  amount, err := strconv.Atoi(stringAmount)
  if err != nil {
    return 0, err
  }

  switch direction {
  case "L":
    amount *= -1
  case "R":
    // do nothing
  default:
    return 0, errors.New("Expected L or R for direction, got " + direction)
  }

  return amount, nil
}

func parseInput(input string) []string {
  return strings.Split(input, "\n")
}

type solution struct{}

func (s solution) Part1(input string) string {
  rotations := parseInput(input)
  var ptr int = 50

  var password int

  for i := range rotations {
    amount, err := getRotationAmount(rotations[i])
    if err != nil {
      panic(err)
    }

    ptr = mod(ptr +amount, MAX_DIAL_PTR + 1)

    if ptr == 0 {
      password += 1
    }
  }

  return strconv.Itoa(password)
}

func (s solution) Part2(input string) string {
  rotations := parseInput(input)
  var ptr int = 50

  var password int

  for i := range rotations {
    oldPtr := ptr
    amount, err := getRotationAmount(rotations[i])
    if err != nil {
      panic(err)
    }
    ptr = mod(ptr +amount, MAX_DIAL_PTR + 1)

    direction := string(rotations[i][0])

    fullRotationTimes := abs(amount/(MAX_DIAL_PTR + 1))
    password += fullRotationTimes
    if ptr == 0 && amount != 0 {
      password += 1
    } else if direction == "L" && ptr > oldPtr && oldPtr != 0 {
      password += 1
    } else if direction == "R" && ptr < oldPtr && oldPtr != 0 {
      password += 1
    }
  }

  return  strconv.Itoa(password)
}

var Solution solution = solution{}
