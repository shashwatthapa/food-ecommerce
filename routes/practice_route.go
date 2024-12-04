package routes

import (
	"github.com/gorilla/mux"
	"github.com/shashwatthapa/food-ecommerce/controllers"
)

func InitAllPracticeRoutes(router *mux.Router){
	router.HandleFunc("/error",controllers.HandleErrorRequest)
	router.HandleFunc("/success",controllers.HandleSuccessRequest)
}