package controllers

import (
	"first_project/models"
	"first_project/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers() // Fetch users from the model
	utils.JSONResponse(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	user, err := models.GetUserByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, "User not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}
