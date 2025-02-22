package server

import (
	"bitsplit_backend/models"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	print("creating user")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := readBody(r)


	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON format, "+err.Error(), http.StatusBadRequest)
		return
	}

	err = s.CRUD.CreateUser(user)
	if err != nil {
		http.Error(w,"error creating user " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	users, err := s.CRUD.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"users": users})
}

func (s *Server) SearchInUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	queries := r.URL.Query()
	users, err := s.CRUD.SearchInUsers(queries["q"][0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"users": users})
}
