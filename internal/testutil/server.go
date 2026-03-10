package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Server struct {
	*httptest.Server
	mux *http.ServeMux
}

func NewServer(t *testing.T) *Server {
	t.Helper()
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return &Server{Server: srv, mux: mux}
}

func (s *Server) Handle(pattern string, statusCode int, body string) {
	s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(body))
	})
}
