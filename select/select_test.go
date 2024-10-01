package selectpa

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
  //mocking our servers
  slowServer := mockServerFabric(time.Millisecond*10)
  fastServer := mockServerFabric(time.Millisecond*1)

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

func TestRacer(t *testing.T) {
  t.Run("retornar erro se servidor nao responder com tempo especificado", func(t *testing.T){
    server := mockServerFabric(time.Millisecond*25)

    defer server.Close()

    _, err := ConfigurableRacer(server.URL, server.URL,time.Millisecond * 20)

    if err == nil {
      t.Errorf("Expected an error but didnt got one")
    }
  })
}

func mockServerFabric(delay time.Duration) *httptest.Server{
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}
