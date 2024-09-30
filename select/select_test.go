package selectpa

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
  //mocking our servers
  slowServer := mockServerFabric(time.Second*10)
  fastServer := mockServerFabric(time.Second*1)

  defer slowServer.Close()
  defer fastServer.Close()

  slowUrl := slowServer.URL
  fastUrl := fastServer.URL

  got := WebsiteRacer(fastUrl, slowUrl)
  want := fastUrl

  if got != want {
    t.Errorf("Received %q, want %q", got, want)
  }
}



func mockServerFabric(delay time.Duration) *httptest.Server{
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}
