package models

import (
	"errors"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

func GetAllUsers() []User {
	return users
}

func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
