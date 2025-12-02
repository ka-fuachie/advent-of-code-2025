package util

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadInput(day int) (string, error) {
  path := filepath.Join("data", "day_" + strconv.Itoa(day) + ".txt")
  data, err := os.ReadFile(path)
  if err != nil {
    return "", err
  }

  return strings.Trim(string(data), "\n"), nil
}
