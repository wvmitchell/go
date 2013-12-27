package wc

import (
  "reflect"
  "strings"
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
  split_words := strings.Split(words, " ")
  for i := range split_words {
    h[split_words[i]] += 1
  }
}

