package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"squeezecnn/common"
	"squeezecnn/controllers"
)

func SetEmployeeRoutes(router *mux.Router) *mux.Router {
	employeeRouter := mux.NewRouter()
	employeeRouter.HandleFunc("/employees/register",controllers.RegisterEmployee).Methods("POST")
	employeeRouter.HandleFunc("/employees/face",controllers.AuthorizeEmployee).Methods("POST")
	employeeRouter.HandleFunc("/employees/list",controllers.GetEmployees).Methods("GET")
	/*employeeRouter.HandleFunc("/recognition/face/{name}",controllers.DeleteEmployee).Methods("DELETE")
	employeeRouter.HandleFunc("/recognition/face/",controllers.UpdateEmployee).Methods("PUT")
	*/
	router.PathPrefix("/employees").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(employeeRouter),
	))
	return router
}


