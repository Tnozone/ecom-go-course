package main

import {
	"net/http"

	"github.com/go-chi/chi/v5"
}

type application struct {
	config config
}

//mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}

//run

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}