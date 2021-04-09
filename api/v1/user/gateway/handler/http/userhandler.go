package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/warkop/up-meetup-clone/api/v1/user/domain"
	"github.com/warkop/up-meetup-clone/api/v1/user/models"
)

type UserHandlerProto interface {
	Create(ctx echo.Context) error
	Fetch(ctx echo.Context) error
	FetchByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type UserHandler struct {
	UseCase domain.UserUseCaseProto
}

func NewUserHandler(ucase domain.UserUseCaseProto) UserHandlerProto {
	return &UserHandler{
		UseCase: ucase,
	}
}

func (uh *UserHandler) Create(ctx echo.Context) error {
	user := new(models.User)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(
			http.StatusUnprocessableEntity,
			err.Error(),
		)
	}

	ucase := uh.UseCase.(*domain.UserUseCase)

	if err := ucase.Repo.Create(user); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		user,
	)
}

func (uh *UserHandler) Fetch(ctx echo.Context) error {
	ucase := uh.UseCase.(*domain.UserUseCase)
	users, err := ucase.Repo.Fetch()

	if err != nil {
		return ctx.JSON(
			http.StatusOK,
			err.Error(),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		users,
	)
}

func (uh *UserHandler) FetchByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	ucase := uh.UseCase.(*domain.UserUseCase)
	user, err := ucase.Repo.FetchByID(int64(id))

	if err != nil {
		return ctx.JSON(
			http.StatusNotFound,
			err.Error(),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		user,
	)
}

func (uh *UserHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	user := new(models.User)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(
			http.StatusUnprocessableEntity,
			err.Error(),
		)
	}

	user.Id = int64(id)

	ucase := uh.UseCase.(*domain.UserUseCase)
	err = ucase.Repo.Update(user.Id, user)

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		user,
	)
}

func (uh *UserHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	ucase := uh.UseCase.(*domain.UserUseCase)

	if err := ucase.Repo.Delete(int64(id)); err != nil {
		return ctx.JSON(
			http.StatusNotFound,
			err.Error(),
		)
	}

	return ctx.NoContent(http.StatusNoContent)
}
