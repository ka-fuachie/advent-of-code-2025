package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ka-fuachie/advent-of-code-2025/util"
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

func part1(grid [][]string) string {
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

func part2(grid [][]string) string {
  return ""
}

func main() {
  input, err := util.ReadInput(4)
  if err != nil {
    panic(err)
  }
//   input := `..@@.@@@@.
// @@@.@.@.@@
// @@@@@.@.@@
// @.@@@@..@.
// @@.@@@@.@@
// .@@@@@@@.@
// .@.@.@.@@@
// @.@@@.@@@@
// .@@@@@@@@.
// @.@.@@@.@.`

  rows := strings.Split(input, "\n")
  grid := make([][]string, len(rows))

  for i, row := range rows {
    grid[i] = strings.Split(row, "")
  }

  fmt.Println(part1(grid))
  fmt.Println(part2(grid))
}
