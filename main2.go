package main

import (
	"database/sql"
	"fmt"
	"log"

	// _ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5"
)

var sqlDB *sql.DB

func TestWithDefault() {
	// Initialize the connection to the PostgreSQL database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	sqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Check if the connection is available
	if err = sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create the table (if it doesn't already exist)
	createTable()

	// Create a user
	createUser2("John Doe")

	// Read a user by ID
	readUser2(1)
}

// Create function to insert a new user
func createUser2(name string) {
	sqlStatement := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var id int
	err := sqlDB.QueryRow(sqlStatement, name).Scan(&id)
	if err != nil {
		log.Fatal("Error creating user:", err)
	}
	fmt.Printf("User created with ID: %d\n", id)
}

// Read function to retrieve a user by ID
func readUser2(id int) {
	var user User
	sqlStatement := `SELECT id, name, score FROM users WHERE id=$1`
	row := sqlDB.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
		} else {
			log.Fatal("Error reading user:", err)
		}
		return
	}

	fmt.Printf("User found: ID = %d, Name = %s, Score = %d\n", user.ID, user.Name, user.Score.Val)
}

// Create table function (if it doesn't already exist)
func createTable() {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		score NUMERIC
	);`
	_, err := sqlDB.Exec(sqlStatement)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}
