package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIServer(listenAdress string, store Storage) *APIServer {
	return &APIServer{
		listenAdress: listenAdress,
		store:        store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/expanse", makeHTTPHandleFunc(s.handleExpanse))
	router.HandleFunc("/expanse/{id}", makeHTTPHandleFunc(s.handleGetExpanseById))

	log.Println("API Running on port ", s.listenAdress)

	http.ListenAndServe(s.listenAdress, router)
}

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func (s *APIServer) handleExpanse(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {

	case "GET":
		{
			return s.handleGetExpanse(w, r)

		}
	case "POST":
		{
			return s.handleCreateExpanse(w, r)

		}
	case "DELETE":
		{
			return s.handleDeleteExpanse(w, r)
		}
	}

	return fmt.Errorf(":( === method not allowed  %s", r.Method)
}

func (s *APIServer) handleGetExpanse(w http.ResponseWriter, r *http.Request) error {
	expanses, err := s.store.GetExpanses()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, expanses)
}

func (s *APIServer) handleGetExpanseById(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &Expanse{})
}

func (s *APIServer) handleCreateExpanse(w http.ResponseWriter, r *http.Request) error {

	CreateExpanseReq := new(CreateExpanseRequest)
	if err := json.NewDecoder(r.Body).Decode(CreateExpanseReq); err != nil {

		return err
	}
	expanse := NewExpanse(CreateExpanseReq.FirstName, CreateExpanseReq.LastName)
	if err := s.store.CreateExpanse(expanse); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, expanse)
}

func (s *APIServer) handleDeleteExpanse(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleModifyExpanse(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

type APIServer struct {
	listenAdress string
	store        Storage
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}
