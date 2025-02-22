package crud

import (
	"bitsplit_backend/models"
)

func (crud CRUD) CreateUser(user models.User) error {
	query := `INSERT INTO USERS (user_id, user_name, email, firebase_uid) VALUES (?, ?, ?, ?)`
	_, err := crud.DB.Exec(query, user.UserID, user.UserName, user.Email, user.FirebaseUID)
	return err
}

func (crud CRUD) GetUserByID(userID string) (models.User, error) {
	var user models.User
	query := `SELECT user_id, user_name, email, firebase_uid FROM USERS WHERE user_id = ?`
	err := crud.DB.QueryRow(query, userID).Scan(&user.UserID, &user.UserName, &user.Email, &user.FirebaseUID)
	return user, err
}

func (crud CRUD) GetAllUsers() ([]models.User, error) {
	query := `SELECT user_id, user_name, email, firebase_uid FROM USERS`
	rows, err := crud.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.FirebaseUID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (crud CRUD) SearchInUsers(searchQuery string) ([]models.User, error) {
	query := `SELECT user_id, user_name, email, firebase_uid FROM USERS WHERE user_name LIKE ?`
	searchTerm := "%" + searchQuery + "%"
	rows, err := crud.DB.Query(query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.FirebaseUID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
