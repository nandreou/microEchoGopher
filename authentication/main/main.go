package main

import (
	"fmt"
	"log"
	"net/http"

	"github.io/nicksauth/database"
	"github.io/nicksauth/handlers"
	"github.io/nicksauth/routers"
)

func main() {
	//var app config.App

	db, err := database.ConnectToDb()

	if err != nil {
		log.Fatal("Server Failed to Start --> ", err)
	}

	handlers.NewRepo(db)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routers.NewRouter(),
	}

	fmt.Println("Server is Up port :8080 !!")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server Failed to Start: ", err)
	}
}
