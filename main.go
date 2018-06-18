package main

import (
	"squeezecnn/common"
	"github.com/codegangsta/negroni"
	"squeezecnn/routers"
	"net/http"
	"log"
)

func main() {
	common.StartUp()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}

	log.Println("Started server. Listening...")
	server.ListenAndServe()
}
