package wc

import (
	"reflect"
	"strings"
  //"regexp"
)

type Histogram map[string]int

func (h Histogram) Equal(other Histogram) bool {
	return reflect.DeepEqual(h, other)
}

func WordCount(words string) Histogram {
	h := make(Histogram)
  cleaned := cleanWords(words)
	populateHistogram(h, cleaned)
	return h
}

func populateHistogram(h Histogram, words string) {
  words_collection := strings.Split(words, " ")
  for _, word := range words_collection {
    h[word] += 1
  }
}

func cleanWords(words string) string {
  return words
}
