package http

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/warkop/up-meetup-clone/api/v1/user/domain"
	"github.com/warkop/up-meetup-clone/api/v1/user/gateway/presenter"
	"github.com/warkop/up-meetup-clone/api/v1/user/models"
)

type UserHandlerProto interface {
	Create(ctx *fiber.Ctx) error
	Fetch(ctx *fiber.Ctx) error
	FetchByID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type UserHandler struct {
	UseCase domain.UserUseCaseProto
}

func NewUserHandler(ucase domain.UserUseCaseProto) UserHandlerProto {
	return &UserHandler{
		UseCase: ucase,
	}
}

func (uh *UserHandler) Create(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(presenter.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
	}

	ucase := uh.UseCase.(*domain.UserUseCase)

	if err := ucase.Repo.Create(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(presenter.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(presenter.WebResponse{
		Code: http.StatusOK,
		Data: user,
	})
}

func (uh *UserHandler) Fetch(ctx *fiber.Ctx) error {
	ucase := uh.UseCase.(*domain.UserUseCase)
	users, err := ucase.Repo.Fetch()

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(presenter.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(presenter.WebResponse{
		Code: http.StatusOK,
		Data: users,
	})
}

func (uh *UserHandler) FetchByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(presenter.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ucase := uh.UseCase.(*domain.UserUseCase)
	user, err := ucase.Repo.FetchByID(int64(id))

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(presenter.WebResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(presenter.WebResponse{
		Code: http.StatusOK,
		Data: user,
	})
}

func (uh *UserHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(presenter.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(presenter.WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
	}

	user.Id = int64(id)

	ucase := uh.UseCase.(*domain.UserUseCase)
	err = ucase.Repo.Update(user.Id, user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(presenter.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(presenter.WebResponse{
		Code: http.StatusOK,
		Data: user,
	})
}

func (uh *UserHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(presenter.WebResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	ucase := uh.UseCase.(*domain.UserUseCase)

	if err := ucase.Repo.Delete(int64(id)); err != nil {
		return ctx.Status(http.StatusNotFound).JSON(
			presenter.WebResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
	}

	return ctx.Status(http.StatusOK).JSON(presenter.WebResponse{
		Code: http.StatusOK,
		Data: nil,
	})
}
