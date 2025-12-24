package day10

import (
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

type machine struct {
  indicator uint
  buttons []button
}

type button struct {
  mask uint
}

func (b button) toggle(indicator uint) uint {
  return (indicator & ^b.mask) | (^indicator & b.mask)
}

func parseMachine(machineStr string) (machine, error) {
  partsStrList := strings.Split(machineStr, " ")
  indicatorPartStr := partsStrList[0]
  buttonStrList := partsStrList[1: len(partsStrList) - 1]

  indicator, err := parseIndicatorLights(indicatorPartStr)
  if err != nil {
    return machine{}, err
  }

  buttons := make([]button, 0, len(buttonStrList))
  for _, buttonStr := range buttonStrList {
    button, err := parseButton(buttonStr, utf8.RuneCountInString(indicatorPartStr) - 2)
    if err != nil {
      return machine{}, err
    }
    buttons = append(buttons, button)
  }

  return machine{indicator, buttons}, nil
}

func parseIndicatorLights(indicatorStr string) (uint, error) {
  indicatorRunes := []rune(indicatorStr)
  indicatorValueRunes := indicatorRunes[1:len(indicatorRunes) - 1]
  var indicator uint

  for i, valueRune := range indicatorValueRunes {
    valueStr := string(valueRune)
    if valueStr == "#" {
      indicator |= (1 << (len(indicatorValueRunes) - 1 - i))
    }
  }

  return indicator, nil
}

func parseButton(buttonStr string, indicatorLength int) (button, error) {
  buttonRunes := []rune(buttonStr)
  buttonMaskRunes := buttonRunes[1: len(buttonRunes) - 1]
  buttonMaskStr := string(buttonMaskRunes)
  buttonMaskIndexStrList := strings.Split(buttonMaskStr, ",")
  var mask uint

  for _, buttonMaskIndexStr := range buttonMaskIndexStrList {
    maskIndex, err := strconv.Atoi(buttonMaskIndexStr)
    if err != nil {
      return button{}, err
    }

    mask |= 1 << (indicatorLength - 1 - maskIndex)
  }

  return button{mask}, nil
}

type machineStateNode struct {
  state uint
  cost int
  machine *machine
}

func (node machineStateNode) nextChild() <- chan machineStateNode {
  ch := make(chan machineStateNode)
  go func() {
    for _, button := range (*node.machine).buttons {
      ch <- machineStateNode{button.toggle(node.state), node.cost + 1, node.machine}
    }
    close(ch)
  }()

  return ch
}

type queue[T any] struct {
  slice []T
}

func (q *queue[T]) enqueue(value T) {
  q.slice = append(q.slice, value)
}

func (q *queue[T]) dequeue() T {
  value := q.slice[0]
  q.slice = slices.Delete(q.slice, 0, 1)
  return value
}

func bfs(startNode machineStateNode, isSolution func(node machineStateNode) bool) machineStateNode {
  searchQueue := queue[machineStateNode]{make([]machineStateNode, 0)}
  searchQueue.enqueue(startNode)

  for len(searchQueue.slice) > 0 {
    node := searchQueue.dequeue()
    if isSolution(node) {
      return node
    }

    for childNode := range node.nextChild() {
      searchQueue.enqueue(childNode)
    }
  }

  return machineStateNode{}
}

type solution struct {}

func (s solution) Part1(input string) string {
  machineStrList := strings.Split(input, "\n")
  machines := make([]machine, 0, len(machineStrList))
  for _, machineStr := range machineStrList {
    parsedMachine, err := parseMachine(machineStr)
    if err != nil {
      panic(err)
    }
    machines = append(machines, parsedMachine)
  }

  totalButtonPresses := 0
  var wg sync.WaitGroup
  for _, m := range machines {
    wg.Go(func() {
      totalButtonPresses += bfs(machineStateNode{0, 0, &m}, func (node machineStateNode) bool {
        return node.state == node.machine.indicator
      }).cost
    })
  }

  wg.Wait()

  return strconv.Itoa(totalButtonPresses)
}

func (s solution) Part2(input string) string {
  return ""
}

var Solution solution = solution{}
