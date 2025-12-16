package day05

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

type idRange struct {
  start int
  end int
}

func (self idRange) includes(id int) bool {
  return self.start <= id && self.end >= id
}

func (self idRange) intersects(otherIdRange idRange) bool {
  return (
    self.includes(otherIdRange.start) ||
    self.includes(otherIdRange.end) ||
    otherIdRange.includes(self.start) ||
    otherIdRange.includes(self.end))
}

func (self idRange) merge(otherIdRange idRange) (idRange, error) {
  if !self.intersects(otherIdRange) {
    return idRange{}, errors.New(
      "Cannot merge no intersecting id ranges: " +
      self.toString() + "," + otherIdRange.toString())
  }

  return idRange{
    start: min(self.start, otherIdRange.start),
    end: max(self.end, otherIdRange.end),
  }, nil
}

func (self idRange) toString() string {
  return strconv.Itoa(self.start) + "-" + strconv.Itoa(self.end)
}

func (self idRange) length() int {
  return self.end - self.start + 1
}

func parseIdRange(idRangeStr string) (idRange, error) {
  rangeStrList := strings.Split(idRangeStr, "-")

  start, err := strconv.Atoi(rangeStrList[0])
  if err != nil {
    return idRange{}, err
  }
  end, err := strconv.Atoi(rangeStrList[1])
  if err != nil {
    return idRange{}, err
  }

  return idRange{start: start, end: end}, nil
}

func _resolveIdRanges(idRanges []idRange) []idRange {
  partiallyResolvedIdRanges := make([]idRange, 0, len(idRanges))

  for _, idRangeItem := range idRanges {
    hasRangeIntersects := false
    for resolvedIdRangeIndex, resolvedIdRange := range partiallyResolvedIdRanges {
      mergedIdRanges, err := resolvedIdRange.merge(idRangeItem)
      if err != nil {
        continue
      }

      partiallyResolvedIdRanges[resolvedIdRangeIndex] = mergedIdRanges
      hasRangeIntersects = true
    }

    if !hasRangeIntersects {
      partiallyResolvedIdRanges = append(partiallyResolvedIdRanges, idRangeItem)
    }
  }

  return partiallyResolvedIdRanges
}

func resolveIdRanges(idRanges []idRange) []idRange {
  // brute force
  resolveIdRanges := _resolveIdRanges(idRanges)
  for {
    partiallyResolveIdRanges := _resolveIdRanges(resolveIdRanges)
    if len(partiallyResolveIdRanges) == len(resolveIdRanges) {
      break
    }

    resolveIdRanges = partiallyResolveIdRanges
  }

  return resolveIdRanges
}

func parseInput(input string) ([]idRange, []int) {
  inputParts := strings.Split(input, "\n\n")
  freshIdRangesPart := inputParts[0]
  availableIdsPart := inputParts[1]

  freshIdRangeStrList := strings.Split(strings.Trim(freshIdRangesPart, "\n"), "\n")
  availableIdStrList := strings.Split(strings.Trim(availableIdsPart, "\n"), "\n")

  freshIdRanges := make([]idRange, 0, len(freshIdRangeStrList))
  for _, idRangeStr := range freshIdRangeStrList {
    idRange, err := parseIdRange(idRangeStr)
    if err != nil {
      panic(err)
    }
    freshIdRanges = append(freshIdRanges, idRange)
  }

  availableIds := make([]int, 0, len(availableIdStrList))
  for _, idStr := range availableIdStrList {
    id, err := strconv.Atoi(idStr)
    if err != nil {
      panic(err)
    }
    availableIds = append(availableIds, id)
  }

  return freshIdRanges, availableIds
}

type solution struct {}

func (s solution) Part1(input string) string {
  freshIdRanges, availableIds := parseInput(input)
  totalFreshAvailableingredients := 0
  for _, id := range availableIds {
    for _, idRange := range freshIdRanges {
      if !idRange.includes(id) {
        continue
      }

      totalFreshAvailableingredients+=1
      break
    }
  }

  return strconv.Itoa(totalFreshAvailableingredients)
}

func (s solution) Part2(input string) string {
  freshIdRanges, _ := parseInput(input)
  totalIdsWithinFreshRanges := 0
  for _, idRange := range resolveIdRanges(freshIdRanges) {
    totalIdsWithinFreshRanges += idRange.length()
  }

  return strconv.Itoa(totalIdsWithinFreshRanges)
}

var Solution solution = solution{}
