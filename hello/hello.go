package main

import (
	"fmt"
	"strings"
)

func Hello (name string) string{
  if strings.Trim(name, " ") == "" {
    return "Hello, world!"
  }
  return fmt.Sprintf("Hello, %s!", name)
}

func main(){
  fmt.Println(Hello("Lucas"))
}
