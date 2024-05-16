package main

import (
	"fmt"
	"log"
	"net/http"

	"guthub.io/nicksbroker/database"
	"guthub.io/nicksbroker/handlers"
	"guthub.io/nicksbroker/logger"
	"guthub.io/nicksbroker/routers"
)

const (
	AllowedHosts = ""
	Port         = ":8000"
)

func main() {

	logger := logger.SetUpLogger()
	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	handlers.NewHandlConf(logger, db)

	srv := http.Server{
		Addr:    Port,
		Handler: routers.NewRouter(),
	}

	fmt.Println("Server is Up on port 8000!!!")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
