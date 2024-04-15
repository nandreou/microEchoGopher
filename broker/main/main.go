package main

import (
	"fmt"
	"log"
	"net/http"

	"guthub.io/nicksbroker/config"
	"guthub.io/nicksbroker/handlers"
	"guthub.io/nicksbroker/routers"
)

func main() {

	var app config.App

	handlers.NewRepo(&app)

	srv := http.Server{
		Addr:    ":8080",
		Handler: routers.NewRouter(),
	}

	fmt.Println("Server is Up !!!")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
