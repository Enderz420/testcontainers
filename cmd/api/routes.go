package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	standard := alice.New()

	mux.HandleFunc("GET /api/v1/healthcheck", app.HealthcheckHandler)

	mux.HandleFunc("GET /api/v1/user/{id}", app.GetUserHandler)
	mux.HandleFunc("GET /api/v1/user", app.ListUserHandler)
	mux.HandleFunc("POST /api/v1/user", app.PostUserHandler)
	mux.HandleFunc("DELETE /api/v1/user/{id}", app.DeleteUserHandler)

	mux.HandleFunc("GET /api/v1/blogpost/{id}", app.GetBlogpostHandler)
	mux.HandleFunc("GET /api/v1/blogpost", app.ListBlogpostHandler)
	mux.HandleFunc("POST /api/v1/blogpost", app.PostBlogpostHandler)

	handler := standard.Then(mux)
	return handler
}
