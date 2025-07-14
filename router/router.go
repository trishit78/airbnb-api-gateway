package router

import (
	"AuthInGo/controllers"
	//"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux{
	router := chi.NewRouter();
	router.Get("/ping",controllers.PingHandler)

	return router

}