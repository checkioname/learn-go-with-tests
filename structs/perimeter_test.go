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

  //funcao helper para verificar area de um shape
  checkArea := func(t *testing.T, shape Shape, want float64) {
    t.Helper()
    got := shape.Area()

    if got != want {
      t.Errorf("Expected %.2f but got %.2f", want, got)
    }
  }

  areaTests := []struct {
    shape Shape
    want float64
  } {
      {Rectangle {10.0, 10.0},20.0},
      {Circle {10},314.1592653589793},
  }

  for _, item :=range areaTests {
    checkArea(t, item.shape, item.want)
  }
}


