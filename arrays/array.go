package arrays

func SumItems(numbers []int) int{
  result := 0

  for _, number := range numbers {
    result += number
  }
  return result
}


// We need a new function called SumAll which will take a varying number of slices, 
// returning a new slice containing the totals for each slice passed in.
// For example
// SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}

func SumAll( arraySum ...[]int) []int {
  lenght := len(arraySum)
  result := make([]int, lenght)

  for i, numbers := range arraySum {
    result[i] = SumItems(numbers)
  }

  return result
}
