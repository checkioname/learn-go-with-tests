package concurrency

import (
  "time"
  "fmt"
  "testing"
)


func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(50 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
    if i == 0 {
      fmt.Println("benchmarking started")
    }
	}
}
