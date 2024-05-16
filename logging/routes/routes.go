package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.io/nickslogging/handlers"
)

func NewRouter() http.Handler {
	mux := chi.NewRouter()

	mux.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))

	mux.Post("/broker-request", handlers.HandlersConf.InsertBrokerReqLog)
	mux.Post("/broker-response", handlers.HandlersConf.InsertBrokerRespLog)
	mux.Post("/auth-request", handlers.HandlersConf.InsertAuthReqLog)
	mux.Post("/auth-response", handlers.HandlersConf.InsertAuthRespLog)

	return mux
}
