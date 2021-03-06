package racer_test

import (
	"github.com/vicgrygorchyk/example_hello/racer"
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(20 * time.Millisecond)
        w.WriteHeader(http.StatusOK)
    }))

	defer slowServer.Close()

    fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

	defer fastServer.Close()

	slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got := racer.Racer(slowURL, fastURL)

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}