package wc

import (
	"reflect"
	"regexp"
)

type Histogram map[string]int

func (h Histogram) Equal(other Histogram) bool {
	return reflect.DeepEqual(h, other)
}

func WordCount(words string) Histogram {
	h := make(Histogram)
	populateHistogram(h, words)
	return h
}

func populateHistogram(h Histogram, words string) {
	words_collection := collectWords(words)
	for _, word := range words_collection {
		h[word] += 1
	}
}

func collectWords(words string) []string {
	rx, _ := regexp.Compile("([a-zA-Z]+)")
	return rx.FindAllString(words, -1)
}
