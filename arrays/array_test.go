package arrays

import (
	"reflect"
	"testing"
)

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


func TestSumAll(t *testing.T) {
  t.Run("Sum two slices into 1", func(t *testing.T) {
    sliceOne := []int { 1, 1 }
    sliceTwo := []int { 2, 2 }

    got := SumAll(sliceOne, sliceTwo)
    want := []int {2,4}
    if !reflect.DeepEqual(got,want) {
      t.Errorf("Expected %v but got %v", want, got)
    }
  })
}
