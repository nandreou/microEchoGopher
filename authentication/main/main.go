package main

import (
	"fmt"
	"log"
	"net/http"

	"github.io/nicksauth/database"
	"github.io/nicksauth/handlers"
	"github.io/nicksauth/logger"
	"github.io/nicksauth/routers"
)

func main() {

	db, err := database.ConnectToDb()

	if err != nil {
		log.Fatal("Server Failed to Start --> ", err)
	}

	handlers.NewHandlConf(logger.SetUpLogger(), db)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routers.NewRouter(),
	}

	fmt.Println("Server is Up port :8080 !!")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server Failed to Start: ", err)
	}
}
