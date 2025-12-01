package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ka-fuachie/advent-of-code-2025/util"
)

const MAX_DIAL_PTR = 99

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

func part1(rotations []string) string {
  var ptr int = 50

  var password int

  for i := 0; i < len(rotations) - 1; i++ {
    amount, err := getRotationAmount(rotations[i])
    if err != nil {
      panic(err)
    }

    ptr += amount
    for ptr < 0 {
      ptr += (MAX_DIAL_PTR + 1)
    }
    ptr %= MAX_DIAL_PTR + 1

    if ptr == 0 {
      password += 1
    }
  }

  return strconv.Itoa(password)
}

func part2(rotations []string) string {
  var ptr int = 50

  var password int

  for i := 0; i < len(rotations) - 1; i++ {
    var fullRotationTimes int
    oldPtr := ptr
    amount, err := getRotationAmount(rotations[i])
    if err != nil {
      panic(err)
    }
    ptr += amount
    for ptr < 0 {
      ptr += (MAX_DIAL_PTR + 1)
    }
    ptr %= MAX_DIAL_PTR + 1
    if amount < 0 {
      currentAmount := amount + (MAX_DIAL_PTR + 1)
      for currentAmount < 0 {
        fullRotationTimes++
        currentAmount += (MAX_DIAL_PTR + 1)
      }
    } else if amount > 0 {
      fullRotationTimes += amount/(MAX_DIAL_PTR + 1)
    }

    direction := string(rotations[i][0])

    password += fullRotationTimes
    if ptr == 0 {
      password += 1
    } else if direction == "L" && ptr > oldPtr && oldPtr != 0 {
      password += 1
    } else if direction == "R" && ptr < oldPtr && oldPtr != 0 {
      password += 1
    }
  }

  return  strconv.Itoa(password)
}


func main() {
  input, err := util.ReadInput(1)
  if err != nil {
    panic(err)
  }

  rotations := strings.Split(input, "\n")

  fmt.Println(part1(rotations))
  fmt.Println(part2(rotations))
}
