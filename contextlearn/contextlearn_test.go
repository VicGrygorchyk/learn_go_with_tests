package contextlearn_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"github.com/vicgrygorchyk/example_hello/contextlearn"
	"testing"
	"strings"
	"time"
)

type SpyStore struct {
	response string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s* SpyStore) Cancel() {
	s.cancelled = true
}

func TestContext(t *testing.T) {
	t.Run("Check happypath", func(t *testing.T) {
		data := "Test"
		stub := SpyStore{response: data}
		svr := contextlearn.Server(&stub)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		responseBody := strings.TrimSpace(response.Body.String())
		if  responseBody != data {
			t.Errorf(`got "%s", want "%s"`, responseBody, data)
		}

		if stub.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := contextlearn.Server(store)
  
		request := httptest.NewRequest(http.MethodGet, "/", nil)
  
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
  
		response := httptest.NewRecorder()
  
		svr.ServeHTTP(response, request)
  
		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}
