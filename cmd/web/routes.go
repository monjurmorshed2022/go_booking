package main

import (
	"net/http"

	"github.com/monjurmorshed2022/go_booking/pkg/config"
	"github.com/monjurmorshed2022/go_booking/pkg/handlers"

	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
