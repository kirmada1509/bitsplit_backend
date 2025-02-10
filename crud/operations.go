package crud

import (
	"bitsplit_backend/models"
	"database/sql"
)

type CRUD struct {
	DB *sql.DB
}

func NewCRUD(db *sql.DB) CRUD {
	return CRUD{DB: db}
}

// createUser creates a new user in the database
func (crud CRUD) CreateUser(user models.User) error {
	query := `INSERT INTO users (name, email, uid) VALUES (?, ?, ?)`
	_, err := crud.DB.Exec(query, user.Name, user.Email, user.UID)
	return err
}

func (crud CRUD) GetUserByID(userID int) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, uid FROM users WHERE id = ?`
	err := crud.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.UID)
	return user, err
}

func (crud CRUD) UpdateUser(user models.User) error {
    query := `UPDATE users SET name = ?, email = ?, uid = ? WHERE id = ?`
    _, err := crud.DB.Exec(query, user.Name, user.Email, user.UID, user.ID)
    return err
}

func (crud CRUD) DeleteUser(userID int) error {
    query := `DELETE FROM users WHERE id = ?`
    _, err := crud.DB.Exec(query, userID)
    return err
}

func (crud CRUD) GetUsers() ([]models.User, error) {
    query := `SELECT id, name, email, uid FROM users`
    rows, err := crud.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User

    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.UID); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}
