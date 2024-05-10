package main

import (
	"fmt"
	"log"
	"net/http"

	"github.io/nickslogging/handlers"
	"github.io/nickslogging/mongodb"
	"github.io/nickslogging/routes"
)

const port = ":8081"

func main() {

	db, err := mongodb.ConnectToDB()

	if err != nil {
		log.Panic(err)
	}

	handlers.NewHanldersConf(db)

	srv := http.Server{
		Addr:    port,
		Handler: routes.NewRouter(),
	}

	fmt.Println("Server is Up Listening to" + port)

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
