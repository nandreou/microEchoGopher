package main

import (
	"fmt"
	"log"
	"net/http"

	"guthub.io/nicksbroker/config"
	"guthub.io/nicksbroker/database"
	"guthub.io/nicksbroker/handlers"
	"guthub.io/nicksbroker/routers"
)

//TODO CREATE A MECHANISM THAT GIVES A TEMPORARY API KEY
//401 when you miss fields

func main() {

	var app config.App

	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	handlers.NewRepo(&app, db)

	srv := http.Server{
		Addr:    ":8000",
		Handler: routers.NewRouter(),
	}

	fmt.Println("Server is Up on port 8000!!!")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
