package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController handles both GET and POST requests
func UserController(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		// Handle creating user
		CreateUserController(c)
	} else if c.Request.Method == http.MethodGet {
		// Handle getting users
		idParam := c.Param("id")

		if idParam != "" {
			GetUser(c)
		} else {
			GetUsers(c)
		}
	} else if c.Request.Method == http.MethodDelete {
		DeleteUser(c)
	} else {
		// Return 405 if the method is not allowed
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}
