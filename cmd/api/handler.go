package main

import (
	"net/http"

	"enderz.net/testcontainer-test/internal/rest"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	sampleData := "test"
	rest.RespondWithJSON(w, r, sampleData, http.StatusOK)
}