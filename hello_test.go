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
      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
    }
  })
 
  t.Run("say hello world when empty string is supplied", func(t *testing.T){
    got := Hello("")
    want := "Hello, world!"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })



}

