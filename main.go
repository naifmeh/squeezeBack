package main

import (
	"squeezecnn/common"
	"github.com/codegangsta/negroni"
	"squeezecnn/routers"
	"net/http"
	"log"
	"fmt"
	"squeezecnn/controllers"
)

var context *controllers.Context

func main() {
	common.StartUp()

	context = controllers.NewContext()
	defer context.Close()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	fmt.Println("Starting server...")
	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}

	log.Println("Started server. Listening...")
	server.ListenAndServe()
}
