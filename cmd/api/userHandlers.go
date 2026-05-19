package main

import (
	"errors"
	"net/http"
	"time"

	"enderz.net/testcontainer-test/internal/data"
	"enderz.net/testcontainer-test/internal/logging"
	"enderz.net/testcontainer-test/internal/models"
	"enderz.net/testcontainer-test/internal/rest"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

type UserResponse struct {
	ID          mssql.UniqueIdentifier `json:"id"`
	Username    string                 `json:"username"`
	Email       string                 `json:"email"`
	CreatedAt   time.Time              `json:"created_at"`
	LastUpdated time.Time              `json:"last_updated"`
}

type UserListResponse struct {
	Users []*data.User `json:"users"`
}

type CreateUserRequest struct {
	ID       mssql.UniqueIdentifier `json:"id"`
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	Password string                 `json:"password"`
}

func (app *application) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logging.LoggerFromContext(ctx)

	var req CreateUserRequest
	err := rest.ReadJSON(r, &req)
	if err != nil {
		logger.Error("failed to read JSON", "error", err)
		rest.BadRequestResponse(w, r, "unable  to decode data from request")
		return
	}

	hashedPassword, err := rest.HashPassword(req.Password)
	if err != nil {
		logger.ErrorContext(ctx, "unable to hash password", "error", err)
		rest.ServerErrorResponse(w, r, err)
		return
	}

	user := &data.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	result, err := app.models.Users.Insert(ctx, user)
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

	rest.WriteJSON(
		w,
		http.StatusCreated,
		CreateUserRequest{
			ID:       result.ID,
			Username: result.Username,
			Email:    result.Email,
		},
		nil,
	)
}

func (app *application) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := logging.LoggerFromContext(ctx)

	users, err := app.models.Users.GetAll(ctx)
	if err != nil {
		logger.ErrorContext(ctx, "unable to retrieve users", "error", err)
		rest.ServerErrorResponse(w, r, err)
		return
	}

	logger.InfoContext(ctx, "returning users")

	rest.WriteJSON(
		w,
		http.StatusOK,
		UserListResponse{
			Users: users,
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

	user, err := app.models.Users.Get(ctx, id)
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

	rest.WriteJSON(
		w,
		http.StatusOK,
		UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
			LastUpdated: user.LastUpdated,
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

	err = app.models.Users.Delete(ctx, id)
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

	w.WriteHeader(http.StatusNoContent)
}
