package day09

import (
	"math"
	"strconv"
	"strings"
)

type position struct {
  row int
  col int
}

func parsePosition(positionStr string) (position, error) {
  axisValStrList := strings.Split(positionStr, ",")

  col, err := strconv.Atoi(axisValStrList[0])
  if err != nil {
    return position{}, nil
  }
  row, err := strconv.Atoi(axisValStrList[1])
  if err != nil {
    return position{}, nil
  }

  return position{row: row, col: col}, nil
}

func calculateRectArea(positionA position, positionB position) int {
  return int(
    math.Abs(float64(positionA.col - positionB.col + 1) *
    math.Abs(float64(positionA.row - positionB.row + 1))),
  )
}

func parseInput(input string) ([]position, error) {
  positionStrList := strings.Split(input, "\n")
  positions := make([]position, 0, len(positionStrList))
  
  for _, positionStr := range positionStrList {
    positionVal, err := parsePosition(positionStr)
    if err != nil {
      return []position{}, err
    }

    positions = append(positions, positionVal)
  }

  return positions, nil
}

type solution struct {}

func (s solution) Part1(input string) string {
  tilePositions, err := parseInput(input)
  if err != nil {
    panic(err)
  }

  maxRectArea := 0

  for i := 0; i < len(tilePositions) - 1; i++ {
    for j := i + 1; j < len(tilePositions); j++ {
      area := calculateRectArea(tilePositions[i], tilePositions[j])
      maxRectArea = int(math.Max(float64(area), float64(maxRectArea)))
    }
  }

  return strconv.Itoa(maxRectArea)
}

func (s solution) Part2(input string) string {
  return ""
}

var Solution solution = solution{}
