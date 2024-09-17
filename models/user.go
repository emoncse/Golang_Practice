package models

import (
	"database/sql"
	"errors"
	"first_project/database"
	"fmt"
)

// User represents the structure of the user table
type GetUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
	Phone string `json:"phone"`
}

type PostUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"required"`
	ID    int    `json:"id"` // Included for the response, but not required in the request
}

// GetAllUsers fetches all users from the database
func GetAllUsers() ([]GetUser, error) {
	// Debugging: Print when the function starts
	fmt.Println("Starting GetAllUsers...")

	// Query the database and get all users
	rows, err := database.DB.Query("SELECT id, name, email, phone FROM users")
	if err != nil {
		fmt.Println("Error during database query:", err)
		return nil, err
	}
	defer rows.Close()

	// Debugging: Print if query was successful
	fmt.Println("Query successful, iterating over rows")

	var users []GetUser

	// Iterate over the result set and scan into the User struct
	for rows.Next() {
		var user GetUser
		// Debugging: Print each iteration of rows
		fmt.Println("Scanning user row...")

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		// Debugging: Print after each successful scan
		fmt.Printf("Scanned user: %+v\n", user)

		users = append(users, user)
	}

	// Debugging: Print after all rows have been processed
	fmt.Println("Finished iterating over rows")

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error during row iteration:", err)
		return nil, err
	}

	// Debugging: Print before returning the final result
	fmt.Println("Returning user list")
	return users, nil
}

func GetUserByID(id int) (GetUser, error) {
	var user GetUser

	// Query to fetch the user by ID
	err := database.DB.QueryRow("SELECT id, name, email, phone FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Phone)

	if err != nil {
		// If no rows are found, return a specific error
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func CreateUser(user *PostUser) error {
	fmt.Println("Executing CreateUser method...")

	query := `INSERT INTO users (name, email, phone) VALUES ($1, $2, $3) RETURNING id`

	// Execute the insert query and scan the returned ID into the user struct
	err := database.DB.QueryRow(query, user.Name, user.Email, user.Phone).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	// Execute the delete query
	res, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	// Check if any rows were affected (i.e., a user was actually deleted)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, the user does not exist
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	fmt.Printf("User with ID %d deleted successfully\n", id)
	return nil
}
