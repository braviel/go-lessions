package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
func TestRacer(t *testing.T) {
	t.Run("Test response", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL, 40*time.Millisecond)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test time out", func(t *testing.T) {
		serverA := makeDelayServer(50 * time.Millisecond)
		serverB := makeDelayServer(60 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL, 1*time.Second)

		if err == nil {
			t.Error("Expected error when time out")
		}
	})

}
