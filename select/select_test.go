package test

import "testing"






func TestWebsiteRacer(t *testing.T) {
  fastUrl := "http://google.com"
  slowUrl := "http://quii.dev"

  got := WebsiteRacer(fastUrl, slowUrl)
  want := fastUrl

  if got != want {
    t.Errorf("Received %q, want %q", got, want)
  }
}
