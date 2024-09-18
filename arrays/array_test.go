package arrays

import "testing"

func TestSumItems(t *testing.T) {
  
  t.Run("Slice of 5 numbers", func(t *testing.T) {

    numbers := []int{1,1,1,1,1}
    got := SumItems(numbers)
    want := 5

    if got != want {
      t.Errorf("Expected %d but got %d, given %v", want, got, numbers)
    }
  })

  t.Run("Slice of 3 numbers", func(t *testing.T) {
    
    numbers := []int { 2, 3, 3}
    got := SumItems(numbers)
    want := 8

    if got != want {
      t.Errorf("Expected %d but got %d, given %v", want, got, numbers)
    }

  })
}
