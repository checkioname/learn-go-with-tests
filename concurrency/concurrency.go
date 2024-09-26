package concurrency

import "time"
//Here's the setup: a colleague has written a function, CheckWebsites, that checks the status of a list of URLs.

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
      results[u] = wc(url)
    }(url)
	}

  time.Sleep(time.Second *2)

	return results
}



