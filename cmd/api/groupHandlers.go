package main

import (
	"net/http"

	"enderz.net/testcontainer-test/internal/data"
	"enderz.net/testcontainer-test/internal/database"
	"enderz.net/testcontainer-test/internal/rest"
	"github.com/google/uuid"
)

type GroupResponse struct {
	Data data.Group `json:"results"`
}

type GroupListResponse struct {
	Data     []*data.Group      `json:"results"`
	Metadata *database.Metadata `json:"metadata"`
}

func (app *application) GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		rest.BadRequestResponse(w, r, "cannot read ID")
		return
	}

	group, err := app.repo.Group.Read(ctx, id)
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusOK, GroupResponse{Data: *group}, nil)
}

func (app *application) ListGroupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	groups, metadata, err := app.repo.Group.List(ctx)
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusOK, GroupListResponse{Data: groups, Metadata: metadata}, nil)
}

func (app *application) PostGroupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var group data.GroupInsert
	err := rest.ReadJSON(r, &group)
	if err != nil {
		rest.BadRequestResponse(w, r, "unable to parse JSON")
		return
	}

	result, err := app.repo.Group.Insert(ctx, group)
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusCreated, GroupResponse{Data: *result}, nil)
}

func (app *application) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		rest.BadRequestResponse(w, r, "cannot read ID")
		return
	}

	err = app.repo.Group.Delete(ctx, id)
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusNoContent, nil, nil)
}
