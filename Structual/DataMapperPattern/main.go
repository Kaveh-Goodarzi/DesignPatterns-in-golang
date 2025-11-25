package main

import (
	"database/sql"
	"fmt"
	"log"

)

type User struct {
	ID       string
	Username string
	Email    string
}

type UserMapper struct {
	DB *sql.DB
}

// Get user by ID
func (um *UserMapper) GetByID(id string) (*User, error) {
	row := um.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Save (insert) new user
func (um *UserMapper) Save(user *User) error {
	result, err := um.DB.Exec(
		"INSERT INTO users (username, email) VALUES (?, ?)",
		user.Username, user.Email,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = fmt.Sprintf("%d", id)
	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userMapper := &UserMapper{DB: db}

	// get user
	user, err := userMapper.GetByID("1")
	if err != nil {
		log.Fatal(err)
	}
	if user != nil {
		fmt.Println("User:", user.Username)
	}

	// create new user
	newUser := &User{Username: "john", Email: "john@example.com"}
	err = userMapper.Save(newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user with ID:", newUser.ID)
}