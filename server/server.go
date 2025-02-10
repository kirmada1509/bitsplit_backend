package server

import (
	"bitsplit_backend/crud"
	"bitsplit_backend/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Server struct {
	Mux  *http.ServeMux
	CRUD *crud.CRUD
}

func NewServer(crudInstance *crud.CRUD) *Server {
	s := &Server{
		Mux:  http.NewServeMux(),
		CRUD: crudInstance,
	}

	// Define routes
	s.Mux.HandleFunc("/", s.HomeHandler)
	s.Mux.HandleFunc("/user", s.CreateUserHandler) 
	s.Mux.HandleFunc("/user/", s.GetUserByIdHandler)
	s.Mux.HandleFunc("/users", s.GetUsersHandler)
	return s
}

func (s *Server) Start(port string) {
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, s.Mux)
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Read request body
	body, err := func() ([]byte, error) {
		var r io.Reader = r.Body
		b := make([]byte, 0, 512)
		for {
			n, err := r.Read(b[len(b):cap(b)])
			b = b[:len(b)+n]
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				return b, err
			}
			if len(b) == cap(b) {
				b = append(b, 0)[:len(b)]
			}
		}
	}()
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON into user struct
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	} 
	
	// Create a user
	err = s.CRUD.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/"):] // Extract ID from path
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := s.CRUD.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)

}


func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	users, err := s.CRUD.GetUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}