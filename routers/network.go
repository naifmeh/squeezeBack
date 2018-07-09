package routers

import (
	"github.com/gorilla/mux"
	"squeezecnn/controllers"
	"github.com/codegangsta/negroni"
	"squeezecnn/common"
)

func SetNetworkRoutes(router *mux.Router) *mux.Router{
	networkRouter := mux.NewRouter()
	networkRouter.HandleFunc("/network/train",controllers.TrainNetwork).Methods("POST")
	networkRouter.HandleFunc("/network/pictures", controllers.RemovePics).Methods("DELETE")
	router.PathPrefix("/network").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(networkRouter),
	))
	return router
}
