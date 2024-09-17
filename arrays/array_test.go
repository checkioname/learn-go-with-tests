package arrays

import "testing"

func TestSumItems(t *testing.T) {
  got := SumItems()
  want := 5

  if got != want {
    t.Errorf("Expected %q but got %q", want, got)
  }
}
