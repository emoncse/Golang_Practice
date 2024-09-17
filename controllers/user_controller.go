package controllers

import (
	"first_project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers fetches all users and returns them as JSON
func GetUsers(c *gin.Context) {
	users, err := models.GetAllUsers() // Fetch users from the model
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser fetches a user by ID and returns the user as JSON
func GetUser(c *gin.Context) {
	// Extract the 'id' parameter from the URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Return 400 Bad Request if the ID is not a valid integer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the model function to get the user by ID
	user, err := models.GetUserByID(id)
	if err != nil {
		// If an error occurs (e.g., user not found), return a 404 Not Found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the user as JSON with a 200 OK status
	c.JSON(http.StatusOK, user)
}

func CreateUserController(c *gin.Context) {
	var user models.PostUser
	fmt.Print("XXXXXXXXXXXXXX")
	// Bind the incoming JSON to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Print("XXXXXXXXXXXXXX")

	// Print the bound object for debugging
	fmt.Printf("Parsed User Object: %+v\n", user)

	// Call the CreateUser function to insert the new user into the database
	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Return the created user object with the generated ID
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	// Extract the 'id' parameter from the URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Return a 400 Bad Request if the ID is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the model function to delete the user
	err = models.DeleteUser(id)
	if err != nil {
		// If an error occurs during deletion, return a 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID %d deleted successfully", id)})
}
