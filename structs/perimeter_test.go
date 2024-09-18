package structs

import "testing"


func TestPerimeter(t *testing.T) {
  rect := Rectangle {10.0, 10.0}
  got := rect.Perimeter()
  want := 40.0

  if got != want {
    t.Errorf("Expected %.2f but got %2.f", got, want)
  }
}

func TestArea(t *testing.T) {

  t.Run("testing Rectangle area", func(t *testing.T) {
    rect := Rectangle {10.0, 10.0}
    got := rect.Area()
    want := 20.0

    if got != want {
      t.Errorf("Expected %.2f but got %2.f", got, want)
    }
  })

 t.Run("testing Circle area", func(t *testing.T){
    rect := Circle {10}
    got := rect.Area()
    want := 314.1592653589793

    if got != want {
      t.Errorf("Expected %.2f but got %2.f", got, want)
    }
  })
}
