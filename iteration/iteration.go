package iteration

func Repeat(letter string, iter int) string{
  var result string
  for i := 0; i < iter; i ++ {
    result += letter
  }
  return result
}
