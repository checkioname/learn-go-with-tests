package selectpa

import (
	"fmt"
	"net/http"
	"time"
)


func WebsiteRacer(a, b string) string {
  aDuration := measureResponseTime(a)
  bDuration := measureResponseTime(b)

  if aDuration < bDuration {
    return a
  }
  return b
}

func measureResponseTime(url string) time.Duration{
  startRequest := time.Now()
  http.Get(url)
  requestDuration := time.Since(startRequest)
  return requestDuration
}

// website racer using select 
func Racer(a, b string) (winner string, error error){
  select{
  case <- ping(a):
    return a, nil
  case <- ping(b):
    return b, nil
  case <- time.After(time.Second *10):
    return "", fmt.Errorf("timed out")
  }
}

func ping(url string) chan struct{} {
  ch := make(chan struct{})
  go func() {
    http.Get(url)
    close(ch)
  }()
  return ch
}
