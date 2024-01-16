package web

import (
	"datamanager/core"
	"fmt"
	"net/http"
	"shared/utils"
)

type server struct{}

var SERVER_API_PORT = utils.GetEnvAsIntOrDefault("SERVER_API_PORT", 0)

func (s *server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"healthCheck": "OK"}`))
}

func (s *server) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := core.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func App() {
	s := &server{}
	http.HandleFunc("/health_check", s.HealthCheck)
	http.HandleFunc("/books", s.GetBooks)
	http.ListenAndServe(fmt.Sprintf(":%d", SERVER_API_PORT), nil)
}
