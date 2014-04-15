package phonenumber

import (
  "regexp"
  "strings"
)

func Number(raw string) string {
  rx := regexp.MustCompile(`\d`)
  digits := rx.FindAllString(raw, -1)

  if len(digits) != 10 {
    return "0000000000"
  } else {
    return strings.Join(digits, "")
  }
}

func AreaCode(number string) string {
  return string(number[0:3])
}

func Format(number string) string {
  return number
}

