package iteration

import "testing"


func TestRepeat(t * testing.T){
  got := Repeat("a", 5)
  want := "aaaaa"

  if got != want {
    t.Errorf("Expected %q but got %q", want, got)
  }

}
