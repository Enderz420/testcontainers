package main

import (
	"log/slog"
	"net/http"

	"enderz.net/testcontainer-test/internal/data"
	"enderz.net/testcontainer-test/internal/database"
	"enderz.net/testcontainer-test/internal/logging"
	"enderz.net/testcontainer-test/internal/rest"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

type BlogpostResponse struct {
	Data data.Blogpost `json:"results"`
}

type BlogpostListResponse struct {
	Data     []*data.Blogpost   `json:"results"`
	Metadata *database.Metadata `json:"metadata"`
}

func (a application) GetBlogpostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		rest.BadRequestResponse(w, r, "unable to parse id in path")
		return
	}

	result, err := a.models.Blogpost.SelectOne(ctx, mssql.UniqueIdentifier(id))
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusOK, BlogpostResponse{Data: *result}, nil)
}
func (a application) ListBlogpostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, metadata, err := a.models.Blogpost.SelectAll(ctx)
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusOK, BlogpostListResponse{
		Data:     result,
		Metadata: metadata,
	}, nil)
}

func (a application) PostBlogpostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logging.LoggerFromContext(ctx)

	var blogpost data.BlogpostInput
	
	if err := rest.ReadJSON(r, &blogpost); err != nil {
		logger.Error("failed to read JSON", "error", err)  
		rest.BadRequestResponse(w, r, "unable to parse request body")
		return
	}
	logger.Log(ctx, slog.LevelInfo, "request", slog.Any("body", blogpost))
	
	result, err := a.models.Blogpost.Insert(ctx, blogpost)
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusCreated, BlogpostResponse{Data: *result}, nil)
}

func (a application) DeleteBlogpostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		rest.BadRequestResponse(w, r, "unable to parse id in path")
		return
	}

	err = a.models.Blogpost.Delete(ctx, mssql.UniqueIdentifier(id))
	if err != nil {
		rest.ServerErrorResponse(w, r, err)
		return
	}

	rest.RespondWithJSON(w, r, http.StatusOK, "deleted", nil)
}