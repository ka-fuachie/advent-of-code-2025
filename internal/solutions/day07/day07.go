package day07

import (
	"maps"
	"strconv"
	"strings"
	"unicode/utf8"
)

type solution struct {}

func (s solution) Part1(input string) string {
  beamSet := make(map[int]struct{})
  manifold := strings.Split(input, "\n")
  totalSplitsCount := 0

  for row := range manifold {
    if row == 0 {
      beamSet[strings.Index(manifold[row], "S")] = struct{}{}
      continue
    }

    currBeamSet := make(map[int]struct{})
    maps.Copy(currBeamSet, beamSet)
    for beamCol := range currBeamSet {
      smth := string(manifold[row][beamCol])

      if smth == "^" {
        delete(beamSet, beamCol)

        if beamCol > 0 {
          beamSet[beamCol - 1] = struct{}{}
        }
        if beamCol < utf8.RuneCountInString(manifold[row]) - 1 {
          beamSet[beamCol + 1] = struct{}{}
        }
         totalSplitsCount++
      }
    }
  }

  return strconv.Itoa(totalSplitsCount)
}

type tree interface {
  nextChild() <- chan tree
  isSolution() bool
}

func dfsCached(node tree, cache *map[tree]int) (bool, int) {
  cache_ := *cache
  cachedTotalSolutions, exists := cache_[node]
  if exists {
    return cachedTotalSolutions > 0, cachedTotalSolutions
  }

  if node.isSolution() {
    return true, 1
  }

  totalSolutions := 0
  for childNode := range node.nextChild() {
    _, solutions := dfsCached(childNode, cache)
    totalSolutions += solutions
  }

  cache_[node] = totalSolutions
  return totalSolutions > 0, totalSolutions
}

type beamNode struct {
  row int
  col int
  manifold *[]string
}

func (node beamNode) nextChild() <- chan tree {
  ch := make(chan tree)
  go func() {
    manifold := *(node.manifold)
    nextRow := node.row + 1
    if nextRow == len(manifold) {
      close(ch)
      return
    }

    nextManifoldRow := manifold[nextRow]
    smth := string(nextManifoldRow[node.col])

    switch(smth) {
    case ".":
      ch <- beamNode{row: nextRow, col: node.col, manifold: node.manifold}
    case "^":
      if node.col > 0 {
        ch <- beamNode{row: node.row + 1, col: node.col - 1, manifold: node.manifold}
      }
      if node.col < utf8.RuneCountInString(nextManifoldRow) - 1 {
        ch <- beamNode{row: node.row + 1, col: node.col + 1, manifold: node.manifold}
      }
    }

    close(ch)
  }()
  return ch
}

func (node beamNode) isSolution() bool {
  return node.row == len(*node.manifold) - 1
}

func (s solution) Part2(input string) string {
  manifold := strings.Split(input, "\n")
  initialBeamNode := beamNode{row: 0, col: strings.Index(manifold[0], "S"), manifold: &manifold}
  solutionsCache := make(map[tree]int)
  _, solutions := dfsCached(&initialBeamNode, &solutionsCache)
  return strconv.Itoa(solutions)
}

var Solution solution = solution{}
