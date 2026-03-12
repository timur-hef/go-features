package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MockStore struct {
	ExpectedResponse string
	t                *testing.T
}

func (s *MockStore) Fetch(ctx context.Context) (string, error) {
	RealResponse := make(chan string, 1)

	go func() {
		var result strings.Builder
		for _, c := range s.ExpectedResponse {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result.WriteRune(c)
			}
		}
		RealResponse <- result.String()
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-RealResponse:
		return res, nil
	}
}

// httptest.ResponseRecorder doesn't have a way of figuring this out so we'll have to roll our own mock
type MockResponseWriter struct {
	written bool
}

func (s MockResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s MockResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s MockResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
func TestServer(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		data := "hello, world"
		svr := Server(&MockStore{data, t})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &MockStore{ExpectedResponse: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// derive a new cancellingCtx from our request which returns us a cancel function
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)   // schedule that function to be called in 5 milliseconds
		request = request.WithContext(cancellingCtx) // we use this new context in our request

		response := MockResponseWriter{written: false}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("store supposed to be cancelled, without any response")
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &MockStore{ExpectedResponse: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})
}
