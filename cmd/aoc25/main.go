package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day01"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day02"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day03"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day04"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day05"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day06"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day07"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day08"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day09"
	"github.com/ka-fuachie/advent-of-code-2025/internal/solutions/day10"
	"github.com/ka-fuachie/advent-of-code-2025/internal/util"
	flag "github.com/spf13/pflag"
)

var solutions map[int]util.Solution = map[int]util.Solution {
  1: day01.Solution,
  2: day02.Solution,
  3: day03.Solution,
  4: day04.Solution,
  5: day05.Solution,
  6: day06.Solution,
  7: day07.Solution,
  8: day08.Solution,
  9: day09.Solution,
  10: day10.Solution,
}

func main() {
  var inputPath string
  var outputPath string
  flag.StringVarP(&inputPath, "input", "i", "", "Path to input file")
  flag.StringVarP(&outputPath, "output", "o", "", "Path to output file")
  flag.Parse()

  dayStr := flag.Arg(0)
  if dayStr == "" {
    panic(errors.New("Puzzle day is required"))
  }
  day, err := strconv.Atoi(dayStr)
  if err != nil {
    panic(err)
  }

  var input string
  if inputPath != "" {
    content, err := util.ReadFile(inputPath)
    if err != nil {
      panic(err)
    }
    input = content
  } else {
    stdin, err := util.ReadStdIn()
    if err != nil {
      panic(err)
    }
    input = stdin
  }

  if input == "" {
    panic(errors.New("No input found for puzzle"))
  }

  solution, exists := solutions[day]
  if !exists {
    panic(errors.New("Couldn't find solution for day: " + strconv.Itoa(day)))
  }

  output := fmt.Sprintf("%s\n%s\n", solution.Part1(input), solution.Part2(input))
  if outputPath != "" {
    err := util.WriteFile(outputPath, output)
    if err != nil {
      panic(err)
    }
  } else {
    util.WriteStdOut(output)
  }
}
