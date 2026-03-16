package main

import {
	"net/http"

	"github.com/go-chi/chi"
}

type application struct {
	config config
}

//mount
func (app *application) mount() http.Handler {
	r := chi.newRouter()
}

//run

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}