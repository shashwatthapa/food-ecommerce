package routes

import (
	"github.com/gorilla/mux"
	"github.com/shashwatthapa/food-ecommerce/controllers"
)

func InitAllUserRoutes(router *mux.Router){
	router.HandleFunc("/users",controllers.HandleGetAllUsers)
}