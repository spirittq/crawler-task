package web

import "net/http"

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"healthCheck": "OK"}`))
} 

func App() {
	s := &server{}
	http.HandleFunc("/health_check", s.ServeHTTP)
	http.ListenAndServe(":3000", nil)
}