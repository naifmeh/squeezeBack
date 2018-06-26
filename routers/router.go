package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	/* Setting employee routes */
	router = SetEmployeeRoutes(router)
	/* Setting devices routes */
	router = SetDeviceRoutes(router)
	/* Setting network routes */
	router = SetNetworkRoutes(router)

	return router
}
