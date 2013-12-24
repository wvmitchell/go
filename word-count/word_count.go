package wc

import "reflect"

type Histogram struct{
  word_counts map[string]int
}

func NewHistogram(word_counts map[string]int) (h *Histogram) {
  for word, count := range word_counts {
    h.word_counts[word] = count
  }
  return
}

func (h Histogram) Equal(other Histogram) bool {
  return reflect.DeepEqual(h.word_counts, other.word_counts)
}

func WordCount(){
}
