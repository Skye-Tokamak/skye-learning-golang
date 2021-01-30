package myselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()
	t.Run("test", func(t *testing.T) {
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})
	t.Run("return error if no response", func(t *testing.T) {
		server1 := makeDelayedServer(200 * time.Millisecond)
		server2 := makeDelayedServer(250 * time.Millisecond)
		defer server1.Close()
		defer server2.Close()
		_, err := ConfigurableRacer(server1.URL, server2.URL, 5*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
