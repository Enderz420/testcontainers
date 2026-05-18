package main

import (
	"net/http"

	"github.com/justinas/alice"
)


func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	standard := alice.New()

	mux.HandleFunc("GET /api/v1/healthcheck", app.HealthcheckHandler)

	handler := standard.Then(mux)
	return handler
}