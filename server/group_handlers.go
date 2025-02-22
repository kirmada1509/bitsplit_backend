package server

import (
	"bitsplit_backend/models"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := readBody(r)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var group models.Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	err = s.CRUD.CreateGroup(group)
	if err != nil {
		http.Error(w, "Failed to create group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Group created successfully"})
}


func (s *Server) GetAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	groups, err := s.CRUD.GetAllGroups()
	if err != nil {
		http.Error(w, "failed to get all groups " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}


func (s *Server) SearchInGroupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	searchQuery := r.URL.Query().Get("searchQuery")
	if searchQuery == "" {
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	groups, err := s.CRUD.SearchInGroups(searchQuery)
	if err != nil {
		http.Error(w, "Error searching groups "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}


func (s *Server) GetGroupsByOwnerIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ownerID := r.URL.Query().Get("owner_id")
	if ownerID == "" {
		http.Error(w, "Owner ID is required", http.StatusBadRequest)
		return
	}

	groups, err := s.CRUD.GetGroupsByOwnerID(ownerID)
	if err != nil {
		http.Error(w, "Error fetching groups by owner_id "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}

func (s *Server) GetGroupsUserIsInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ownerID := r.URL.Query().Get("user_id")
	if ownerID == "" {
		http.Error(w, "Owner ID is required", http.StatusBadRequest)
		return
	}

	groups, err := s.CRUD.GetGroupsUserIsIn(ownerID)
	if err != nil {
		http.Error(w,"error fetching groups user in in, " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}