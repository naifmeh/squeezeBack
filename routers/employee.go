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
	employeeRouter.HandleFunc("/employees/update", controllers.UpdateEmployee).Methods("PUT")
	employeeRouter.HandleFunc("/employees/face",controllers.AuthorizeEmployee).Methods("POST")
	employeeRouter.HandleFunc("/employees/list",controllers.GetEmployees).Methods("GET")
	employeeRouter.HandleFunc("/employees/image",controllers.SaveEmployeeFace).Methods("POST")
	employeeRouter.HandleFunc("/employees/remove",controllers.RemoveEmployee).Methods("DELETE")
	employeeRouter.HandleFunc("/employees/logs",controllers.GetLog).Methods("GET")
	employeeRouter.HandleFunc("/employees/logs", controllers.DeleteLogFile).Methods("DELETE")

	router.PathPrefix("/employees").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(employeeRouter),
	))
	return router
}


