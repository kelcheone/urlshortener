package api

import "net/http"

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) GetURL(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/"):]

}
