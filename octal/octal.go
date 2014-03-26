package octal

import (
	"strconv"
)

func ToDecimal(octaldecimal string) int64 {
  n, _ := strconv.ParseInt(octaldecimal, 8, 64)
  return n
}
