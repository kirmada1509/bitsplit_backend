package main

import (
	"bitsplit_backend/server"
    "bitsplit_backend/crud"
	"database/sql"
	"fmt"
	"log"
    _ "github.com/mattn/go-sqlite3"
)


func main() {
    // Connect to SQLite database
    db, err := sql.Open("sqlite3", "./mydatabase.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Successfully connected to SQLite database!")
    cd := crud.NewCRUD(db)
    mux := server.NewServer(&cd)
    mux.Start("8080")
}
