package crud

import (
	"database/sql"
)

type CRUD struct {
	DB *sql.DB
}

func NewCRUD(db *sql.DB) CRUD {
	return CRUD{DB: db}
}
