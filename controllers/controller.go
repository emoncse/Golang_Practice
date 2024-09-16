package controllers

import (
	"first_project/utils"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, "Welcome to My App!")
}
