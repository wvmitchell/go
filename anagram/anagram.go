package anagram

import (
  "strings"
  "sort"
)

func Detect(subject string, candidates []string) (anagrams []string) {
  for i := range candidates {
    if anagram(strings.ToLower(subject), strings.ToLower(candidates[i])){
      anagrams = append(anagrams, strings.ToLower(candidates[i]))
    }
  }
  return
}

func anagram(word, test string) bool {
  return !same_word(word, test) && same_letters(word, test)
}

func same_word(first, second string) bool {
  return first == second
}

func same_letters(first, second string) bool {
  return sort_letters(first) == sort_letters(second)
}

func sort_letters(word string) string {
  letters := strings.Split(word, "")
  sort.Strings(letters)
  return strings.Join(letters, "")
}
