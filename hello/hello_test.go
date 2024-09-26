package main

import (
	"fmt"
	"testing"
)


func TestHello(t *testing.T) {
  
  t.Run("saying hi to people", func(t *testing.T) {
    names := []string { "Lucas", "Giovanna", "Jose" }
    for _, name :=  range names{
    
      want := fmt.Sprintf("Hello, %s!", name)
      got := Hello(name)
      assertCorrectMessage(got, want, t)
    }
  })
 
  t.Run("say hello world when empty string is supplied", func(t *testing.T){
    got := Hello("")
    want := "Hello, world!"
    
    assertCorrectMessage(got, want, t)
 })

}


// Helper functions
func assertCorrectMessage(got string, want string, t testing.TB) {
  t.Helper()
  if got != want {
      t.Errorf("got %q want %q", got, want)
    }
}


func BenchmarkHello(b *testing.B) {
  for i := 0; i< b.N; i++ {
    Hello("Lucas")
  }
}
