package server

import (
	"bitsplit_backend/crud"
	"fmt"
	"net/http"
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

	// Define routes for user-related operations
	s.Mux.HandleFunc("/", s.HomeHandler)
	s.Mux.HandleFunc("/create_user", s.CreateUserHandler)
	s.Mux.HandleFunc("/user/", s.GetUserByIdHandler)
	s.Mux.HandleFunc("/all_users", s.GetAllUsersHandler)
	s.Mux.HandleFunc("/search_in_users", s.SearchInUsersHandler)

	// Define routes for group-related operations
	s.Mux.HandleFunc("/create_group", s.CreateGroupHandler)
	s.Mux.HandleFunc("/group/", s.GetGroupByIDHandler)             
	s.Mux.HandleFunc("/update_group", s.UpdateGroupHandler)        
	s.Mux.HandleFunc("/delete_group", s.DeleteGroupHandler)        
	s.Mux.HandleFunc("/all_groups", s.GetAllGroupsHandler)         
	s.Mux.HandleFunc("/search_in_groups", s.SearchInGroupsHandler) 


	// Define routes for group-user-related operations
	s.Mux.HandleFunc("/add_user_to_group", s.AddUserToGroup)
	return s
}

func (s *Server) Start(port string) {
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, s.Mux)
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
