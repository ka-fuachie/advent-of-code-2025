package day04

import (
	"errors"
	"strconv"
	"strings"
)

func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}

func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}

func countRollsInAdjacentPositions(grid [][]string, y int, x int) int {
  totalAdjacentRolls := 0
  for row := max(y - 1, 0); row <= min(y + 1, len(grid) - 1); row++ {
    for col := max(x - 1, 0); col <= min(x + 1, len(grid[row]) - 1); col++ {
      if row == y && col == x {
        continue
      }

      if grid[row][col] == "." {
        continue
      } else if grid[row][col] == "@" {
        totalAdjacentRolls += 1
      } else {
        panic(errors.New("Unexpected character at grid position (" + strconv.Itoa(col) + ", " + strconv.Itoa(row) + "): " + grid[row][col]))
      }
    }
  }

  return totalAdjacentRolls
}

func countRollsInAdjacentPositionsWithRemovedPositions(grid [][]string, y int, x int, removedPositions map[string]bool) int {
  totalAdjacentRolls := 0
  for row := max(y - 1, 0); row <= min(y + 1, len(grid) - 1); row++ {
    for col := max(x - 1, 0); col <= min(x + 1, len(grid[row]) - 1); col++ {
      if row == y && col == x {
        continue
      }

      if grid[row][col] == "." {
        continue
      } else if grid[row][col] == "@" {
        if removedPositions[strconv.Itoa(col) + "," + strconv.Itoa(row)] {
          continue
        }
        totalAdjacentRolls += 1
      } else {
        panic(errors.New("Unexpected character at grid position (" + strconv.Itoa(col) + ", " + strconv.Itoa(row) + "): " + grid[row][col]))
      }
    }
  }

  return totalAdjacentRolls
}

func parseInput(input string) [][]string {
  rows := strings.Split(input, "\n")
  grid := make([][]string, len(rows))

  for i, row := range rows {
    grid[i] = strings.Split(row, "")
  }

  return grid
}

type solution struct {}

func (s solution) Part1(input string) string {
  grid := parseInput(input)
  totalAccessibleRolls := 0

  for row := range len(grid) {
    for col := range len(grid[row]) {
      if grid[row][col] != "@" {
        continue
      }

      totalAdjacentRolls := countRollsInAdjacentPositions(grid, row, col)
      if totalAdjacentRolls < 4 {
        totalAccessibleRolls += 1
      }
    }
  }
  
  return strconv.Itoa(totalAccessibleRolls)
}

func (s solution) Part2(input string) string {
  grid := parseInput(input)
  // Brute force
  totalAccessibleRollsAfterRemovals := 0
  removedPositions := make(map[string]bool)

  for {
    totalAccessibleRolls := 0

    for row := range len(grid) {
      for col := range len(grid[row]) {
        if grid[row][col] != "@" {
          continue
        }

        if removedPositions[strconv.Itoa(col) + "," + strconv.Itoa(row)] {
          continue
        }

        totalAdjacentRolls := countRollsInAdjacentPositionsWithRemovedPositions(grid, row, col, removedPositions)
        if totalAdjacentRolls < 4 {
          totalAccessibleRolls += 1
          removedPositions[strconv.Itoa(col) + "," + strconv.Itoa(row)] = true
        }
      }
    }

    if totalAccessibleRolls == 0 {
      break
    }

    totalAccessibleRollsAfterRemovals += totalAccessibleRolls
  }
  
  return strconv.Itoa(totalAccessibleRollsAfterRemovals)
}

var Solution solution = solution{}
