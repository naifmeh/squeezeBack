package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"squeezecnn/common"
	"squeezecnn/controllers"
)

func SetEmployeeRoutes(router *mux.Router) *mux.Router {
	employeeRouter := mux.NewRouter()
	employeeRouter.HandleFunc("/recognition/register",controllers.RegisterEmployee).Methods("POST")
	employeeRouter.HandleFunc("/recognition/face",controllers.AuthorizeEmployee).Methods("POST")
	/*employeeRouter.HandleFunc("/recognition/face/{name}",controllers.DeleteEmployee).Methods("DELETE")
	employeeRouter.HandleFunc("/recognition/face/",controllers.UpdateEmployee).Methods("PUT")
	*/
	router.PathPrefix("/recognition").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(employeeRouter),
	))
	return router
}
