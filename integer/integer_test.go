package integer

import "testing"

func TestAdd(t *testing.T) {
  t.Run("Should add the second input to the first one", func(t *testing.T) {
    got := Add(5,5)
    want := 10

    assertCorrectMessage(got, want, t) 
  })
}


func assertCorrectMessage(got, want int, t testing.TB) {
  t.Helper()
  if got != want {
    t.Errorf("Got %d should have got %d", got, want)
  }
}

func BenchmarkAdd(b *testing.B) {
  for i := 0; i< b.N; i++ {
    Add(5,5)
  }
}
