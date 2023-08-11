package api

import "net/http"

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", s.handleHealth)

	return http.ListenAndServe(s.listenAddr, mux)
}
