package server

import (
	"bitsplit_backend/models"
	"encoding/json"
	"io"
	"net/http"
)

func (s *Server) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
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

	// Parse JSON into Group struct
	var group models.Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}


	// Create a group
	err = s.CRUD.CreateGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get gid from URL parameters
	groupID := r.URL.Query().Get("gid")
	if groupID == "" {
		http.Error(w, "Group ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve the group by gid
	group, err := s.CRUD.GetGroupByID(groupID)
	if err != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	// Respond with the group data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}


func (s *Server) UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON into Group struct
	var group models.Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Update the group
	err = s.CRUD.UpdateGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group updated successfully"))
}


func (s *Server) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get gid from URL parameters
	groupID := r.URL.Query().Get("gid")
	if groupID == "" {
		http.Error(w, "Group ID is required", http.StatusBadRequest)
		return
	}

	// Delete the group by gid
	err := s.CRUD.DeleteGroup(groupID)
	if err != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group deleted successfully"))
}


func (s *Server) GetAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve all groups
	groups, err := s.CRUD.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of groups as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}


func (s *Server) SearchInGroupsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get search query from URL parameters
	searchQuery := r.URL.Query().Get("query")
	if searchQuery == "" {
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	// Search for groups by name
	groups, err := s.CRUD.SearchInGroups(searchQuery)
	if err != nil {
		http.Error(w, "Error searching groups", http.StatusInternalServerError)
		return
	}

	// Respond with the search results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}


func (s *Server) GetGroupsByOwnerIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get owner_id from URL parameters
	ownerID := r.URL.Query().Get("owner_id")
	if ownerID == "" {
		http.Error(w, "Owner ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve groups by owner_id
	groups, err := s.CRUD.GetGroupsByOwnerID(ownerID)
	if err != nil {
		http.Error(w, "Error fetching groups", http.StatusInternalServerError)
		return
	}

	// Respond with the list of groups as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}

func (s *Server) GetGroupsUserIsInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get owner_id from URL parameters
	ownerID := r.URL.Query().Get("user_id")
	if ownerID == "" {
		http.Error(w, "Owner ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve groups by owner_id
	groups, err := s.CRUD.GetGroupsUserIsIn(ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of groups as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"groups": groups})
}