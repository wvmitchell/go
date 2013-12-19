package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	var output = make(map[string]int)
	for score, letters := range input {
		assignLetterScores(output, letters, score)
	}
	return output
}

func assignLetterScores(output map[string]int, letters []string, score int) {
	for i := range letters {
		var letter = strings.ToLower(letters[i])
		output[letter] = score
	}
}
