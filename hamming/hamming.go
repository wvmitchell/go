package hamming

import "strings"

func Distance(sequence1, sequence2 string) int {
	s1 := strings.Split(sequence1, "")
	s2 := strings.Split(sequence2, "")
	return difference(s1, s2)
}

func difference(s1, s2 []string) (count int) {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			count++
		}
	}
	return
}
