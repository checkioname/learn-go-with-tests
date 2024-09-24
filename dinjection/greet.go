package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet (write io.Writer, name string) {
  fmt.Fprintf(write, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
  Greet(w, "web")
}

func main(){
  //Serving Greets to the web
  log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
