package binary

import (
  "strings"
  "math"
)

func ToDecimal(binary string) int {
  if Valid(binary) {
    return Convert(binary)
  } else {
    return 0
  }
}

func Reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func Valid(s string) bool {
  for _, c := range strings.Split(s, "") {
    if c != "0" && c != "1" {
      return false
    }
  }
  return true
}

func Convert(s string) int{
  var total float64 = 0

  characters := strings.Split(Reverse(s), "")

  for i, char := range characters {
    if char == "1" {
      total += math.Pow(float64(2), float64(i))
    }
  }

  return int(total)
}
