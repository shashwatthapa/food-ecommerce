package controllers

import (
	"net/http"

	"github.com/shashwatthapa/food-ecommerce/models"
	"github.com/shashwatthapa/food-ecommerce/responses"
)

func HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	allModels:= models.GetAllUsers()
	responses.SuccessResponse(w,allModels,"Sucessfully fetched user data")
}