package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiError struct {
	Error string
}

// we use flow chart to represent it

type APIserver struct {
	listenAdd string
	store     storage
	// db connection string to be added
}

func newAPIserver(address string, store storage) *APIserver {
	return &APIserver{
		listenAdd: address,
		store:     store,
	}

}

func (s *APIserver) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandler(s.handleGetAccount))
	log.Println("the server is running on port: ", s.listenAdd)
	http.ListenAndServe(s.listenAdd, router)
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)

	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("meatho0d not allowed %s", r.Method)

}
func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	account := NewAccount("nikhl", "rathore")
	data := map[string]interface{}{
		"account": account,
		"id":      id,
	}

	return writeJson(w, http.StatusOK, data)
}
func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIserver) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Helper functions

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}
