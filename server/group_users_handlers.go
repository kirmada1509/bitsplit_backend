package server

import (
	"bitsplit_backend/models"
	"encoding/json"
	"io"
	"net/http"	
)

func (s *Server) AddUserToGroup(w http.ResponseWriter, r *http.Request){
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
		var groupUser models.GroupUser
		err = json.Unmarshal(body, &groupUser)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
	
		// Create a group
		err = s.CRUD.AddUserToGroup(groupUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
}