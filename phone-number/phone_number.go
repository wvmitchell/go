package phonenumber

import (
  "fmt"
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
  number = remove_country_code(number)
  return fmt.Sprintf("(%s) %s-%s", AreaCode(number), number[3:6], number[6:])
}

func remove_country_code(number string) string {
  if len(number) > 10 {
    return number[1:]
  } else {
    return number
  }
}

