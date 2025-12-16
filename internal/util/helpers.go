package util

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Solution interface {
  Part1(input string) string
  Part2(input string) string
}

func ReadFile(path string) (string, error) {
  data, err := os.ReadFile(path)
  if err != nil {
    return "", err
  }

  return strings.Trim(string(data), "\n"), nil
}

func ReadStdIn() (string, error) {
  data, err := io.ReadAll(os.Stdin)
  if err != nil {
    return "", err
  }

  return strings.Trim(string(data), "\n"), nil
}

func WriteFile(path string, content string) error {
  data := []byte(content)
  err := os.WriteFile(path, data, 0644)
  if err != nil {
    return err
  }

  return nil
}

func WriteStdOut(content string) {
  fmt.Print(content)
}
