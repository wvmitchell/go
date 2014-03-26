package octal

import (
	"math"
	"strconv"
	"strings"
)

func ToDecimal(octaldecimal string) int64 {
	return convertDigits(octaldecimal)
}

func convertDigits(octaldecimal string) int64 {
	result := float64(0)
	digits := strings.Split(reverse(octaldecimal), "")

	for index, digit := range digits {
		as_int, _ := strconv.Atoi(digit)
		result += float64(as_int) * math.Pow(float64(8), float64(index))
	}
	return int64(result)
}

func reverse(octaldecimal string) string {
	runes := []rune(octaldecimal)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
