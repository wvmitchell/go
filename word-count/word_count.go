package wc

import (
	"reflect"
	"strings"
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
  re := regexp.MustCompile(/\W*/)
  cleaned := re.ReplaceAllString(words, "")
	split_words := strings.Split(words, " ")
	for i := range split_words {
		h[split_words[i]] += 1
	}
}
