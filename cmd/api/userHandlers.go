package main

import (
	"errors"
	"net/http"

	"enderz.net/testcontainer-test/internal/data"
	"enderz.net/testcontainer-test/internal/database"
	"enderz.net/testcontainer-test/internal/logging"
	"enderz.net/testcontainer-test/internal/models"
	"enderz.net/testcontainer-test/internal/rest"
	"github.com/google/uuid"
)

type UserResponse struct {
	Data data.User `json:"results"`
}

type UserListResponse struct {
	Data     []*data.User       `json:"results"`
	Metadata *database.Metadata `json:"metadata"`
}

type PostUserRequest struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func (app *application) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logging.LoggerFromContext(ctx)

	var req PostUserRequest
	err := rest.ReadJSON(r, &req)
	if err != nil {
		logger.Error("failed to read JSON", "error", err)
		rest.BadRequestResponse(w, r, "unable to decode data from request")
		return
	}

	user := &data.UserPatch{
		ID:       req.ID,
		Username: req.Username,
		Email:    req.Email,
	}

	result, err := app.repo.User.Update(ctx, *user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateUsername):
			rest.BadRequestResponse(w, r, "username already exists")
		case errors.Is(err, data.ErrDuplicateEmail):
			rest.BadRequestResponse(w, r, "email already exists")
		default:
			logger.ErrorContext(ctx, "unable to insert user into database", "error", err)
			rest.ServerErrorResponse(w, r, err)
		}
		return
	}

	logger.InfoContext(ctx, "user created")

	rest.RespondWithJSON(
		w,
		r,
		http.StatusCreated,
		UserResponse{
			Data: *result,
		},
		nil,
	)
}

func (app *application) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := logging.LoggerFromContext(ctx)

	result, metadata, err := app.repo.User.List(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "unable to retrieve users", "error", err)
		rest.ServerErrorResponse(w, r, err)
		return
	}

	logger.InfoContext(ctx, "returning users")

	rest.RespondWithJSON(
		w,
		r,
		http.StatusOK,
		UserListResponse{
			Data:     result,
			Metadata: metadata,
		},
		nil,
	)
}

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logging.LoggerFromContext(ctx)

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		rest.BadRequestResponse(w, r, "unable to parse id from path")
		return
	}

	user, err := app.repo.User.Read(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			rest.NotFoundResponse(w, r)
		default:
			logger.ErrorContext(ctx, "unable to get user", "error", err)
			rest.ServerErrorResponse(w, r, err)
		}
		return
	}

	logger.InfoContext(ctx, "returning user")

	rest.RespondWithJSON(
		w,
		r,
		http.StatusOK,
		UserResponse{
			Data: *user,
		},
		nil,
	)
}

func (app *application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logging.LoggerFromContext(ctx)

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		rest.BadRequestResponse(w, r, "unable to parse id from path")
		return
	}

	err = app.repo.User.Delete(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			rest.NotFoundResponse(w, r)
		default:
			logger.ErrorContext(ctx, "unable to delete user", "error", err)
			rest.ServerErrorResponse(w, r, err)
		}
		return
	}

	logger.InfoContext(ctx, "user deleted")

	rest.RespondWithJSON(w, r, http.StatusNoContent, nil, nil)
}
