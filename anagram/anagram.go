package anagram

import (
  "strings"
  "sort"
)

func Detect(subject string, candidates []string) []string {
  found_anagrams := make([]string, 0)
  for i := range candidates {
    if anagram(strings.ToLower(subject), strings.ToLower(candidates[i])){
      found_anagrams = append(found_anagrams, strings.ToLower(candidates[i]))
    }
  }
  return found_anagrams
}

func anagram(word, test string) bool {
  return !same_word(word, test) && same_letters(word, test)
}

func same_word(first, second string) bool {
  return first == second
}

func same_letters(first, second string) bool {
  first_letters := strings.Split(first, "")
  second_letters := strings.Split(second, "")
  sort.Strings(first_letters)
  sort.Strings(second_letters)
  new_first := strings.Join(first_letters, "")
  new_second := strings.Join(second_letters, "")
  return new_first == new_second
}

