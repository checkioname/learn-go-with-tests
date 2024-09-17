package arrays

func SumItems() int{
  numbers := [5]int {1,1,1,1,1}
  var result int
  
  for _, number := range numbers {
    result += number
  }

  return result
}
