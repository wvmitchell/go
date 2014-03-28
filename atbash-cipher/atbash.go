package atbash

func Atbash(s string) string {
  cipher_keys := make(map[rune]rune)
  buildCipherKeys(cipher_keys)

  s_runes := []rune(s)

  var cipher_runes []rune

  for _, c := range s_runes {
    cipher_runes = append(cipher_runes, cipher_keys[c])
  }

  return string(cipher_runes)
}

func buildCipherKeys(cipher_keys map[rune]rune){

  letters := "abcdefghijklmnopqrstuvwxyz"
  cipers := reverse(letters)

  for i, c := range letters {
    cs := []rune(cipers)
    cipher_keys[c] = cs[i]
  }

}

func reverse(s string) string {
  runes := []rune(s)

  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }

  return string(runes)
}
