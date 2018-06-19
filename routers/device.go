package routers

import (
	"github.com/gorilla/mux"
	"squeezecnn/controllers"
)

func SetDeviceRoutes( router *mux.Router) *mux.Router {
	router.HandleFunc("/device/register",controllers.Register).Methods("POST")
	router.HandleFunc("/device/authenticate",controllers.Authenticate).Methods("POST")
	return router
}
