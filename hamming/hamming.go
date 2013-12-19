package hamming

import (
	"strings"
)

func Distance(first_sequence, second_sequence string) int {
	first_letters := strings.Split(first_sequence, "")
	second_letters := strings.Split(second_sequence, "")
	return difference(first_letters, second_letters)
}

func difference(first, second []string) int {
	count := 0
	for i := range shorter(first, second) {
		if first[i] != second[i] {
			count++
		}
	}
	return count
}

func shorter(first, second []string) []string {
	if len(first) > len(second) {
		return second
	} else {
		return first
	}
}
