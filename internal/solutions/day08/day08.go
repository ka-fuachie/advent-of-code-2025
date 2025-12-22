package day08

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type point struct {
  x int
  y int
  z int
}

func parsePoint(pointStr string) (point, error) {
  axisValStrList := strings.Split(pointStr, ",")

  x, err := strconv.Atoi(axisValStrList[0])
  if err != nil {
    return point{}, nil
  }
  y, err := strconv.Atoi(axisValStrList[1])
  if err != nil {
    return point{}, nil
  }
  z, err := strconv.Atoi(axisValStrList[2])
  if err != nil {
    return point{}, nil
  }

  return point{x: x, y: y, z: z}, nil
}

func squaredDist(pointA *point, pointB *point) int {
  return int(
    math.Pow(float64(pointA.x - pointB.x), 2) + 
    math.Pow(float64(pointA.y - pointB.y), 2) + 
    math.Pow(float64(pointA.z - pointB.z), 2))
}

func parseInput(input string) ([]point, []struct{pointA *point; pointB *point; value int}, error) {
  pointStrList := strings.Split(input, "\n")
  points := make([]point, 0, len(pointStrList))

  for _, pointStr := range pointStrList {
    pointVal, err := parsePoint(pointStr)
    if err != nil {
      return []point{}, []struct{pointA *point; pointB *point; value int}{}, err
    }

    points = append(points, pointVal)
  }

  squaredDists := make([]struct{
    pointA *point
    pointB *point
    value int
  }, 0, (len(points) - 1) * len(points) / 2) // Triangular numbers formula
  for i := range points {
    pointA := points[i]
    for j := i + 1; j < len(points); j++ {
      pointB := points[j]
      squaredDists = append(squaredDists, struct{pointA *point; pointB *point; value int}{
        &pointA,
        &pointB,
        squaredDist(&pointA, &pointB),
      })
    }
  }

  sort.Slice(squaredDists, func (i, j int) bool {
    return squaredDists[i].value < squaredDists[j].value
  })

  return points, squaredDists, nil
}

func connectJunctions(pointA *point, pointB *point, circuits *[]*map[point]struct{}, circuitMap *map[point]*map[point]struct{}) {
  circuitA := (*circuitMap)[*pointA]
  circuitB := (*circuitMap)[*pointB]

  if circuitA != nil && circuitB != nil {
    if circuitA == circuitB {
      return 
    }

    for junction := range *circuitB {
      (*circuitA)[junction] = (*circuitB)[junction]
      (*circuitMap)[junction] = circuitA
    }
    circuitBIndex := slices.IndexFunc(*circuits, func (circuit *map[point]struct{}) bool {
      return circuit == circuitB
    })
    *circuits = slices.Delete(*circuits, circuitBIndex, circuitBIndex + 1)
  } else if circuitA != nil {
    (*circuitA)[*pointB] = struct{}{}
    (*circuitMap)[*pointB] = circuitA
  } else if circuitB != nil {
    (*circuitB)[*pointA] = struct{}{}
    (*circuitMap)[*pointA] = circuitB
  } else {
    circuit := make(map[point]struct{})
    circuit[*pointA] = struct{}{}
    circuit[*pointB] = struct{}{}
    *circuits = append(*circuits, &circuit)
    (*circuitMap)[*pointA] = &circuit
    (*circuitMap)[*pointB] = &circuit
  }
}

type solution struct {}

func (s solution) Part1(input string) string {
  _, squaredDists, err := parseInput(input)
  if err != nil {
    panic(err)
  }

  circuits := make([]*map[point]struct{}, 0)
  circuitMap := make(map[point]*map[point]struct{}, 0)

  for i := range 1000 {
    s := squaredDists[i]
    connectJunctions(s.pointA, s.pointB, &circuits, &circuitMap)
  }

  sort.Slice(circuits, func (i, j int) bool {
    return len(*circuits[i]) > len(*circuits[j])
  })

  return strconv.Itoa(len(*circuits[0]) * len(*circuits[1]) * len(*circuits[2]))
}

func (s solution) Part2(input string) string {
  points, squaredDists, err := parseInput(input)
  if err != nil {
    panic(err)
  }

  circuits := make([]*map[point]struct{}, 0)
  circuitMap := make(map[point]*map[point]struct{}, 0)

  for _, s := range squaredDists {
    connectJunctions(s.pointA, s.pointB, &circuits, &circuitMap)

    if len(circuits) == 1 && len(points) == len(circuitMap) {
      return strconv.Itoa(s.pointA.x * s.pointB.x)
    }
  }

  return "0"
}

var Solution solution = solution{}
