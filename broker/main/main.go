package main

import (
	"fmt"
	"log"
	"net/http"

	"guthub.io/nicksbroker/database"
	"guthub.io/nicksbroker/handlers"
	"guthub.io/nicksbroker/loger"
	"guthub.io/nicksbroker/routers"
)

func main() {

	loger := loger.SetUpLogger()
	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	handlers.NewHandlConf(loger, db)

	srv := http.Server{
		Addr:    ":8000",
		Handler: routers.NewRouter(),
	}

	fmt.Println("Server is Up on port 8000!!!")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
