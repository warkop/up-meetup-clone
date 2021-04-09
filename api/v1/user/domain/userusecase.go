package domain

import (
	"github.com/warkop/up-meetup-clone/api/v1/user/repository"
)

type UserUseCaseProto interface{}

type UserUseCase struct {
	Repo repository.UserRepositoryProto
}

func NewUserUseCase(repo repository.UserRepositoryProto) *UserUseCase {
	return &UserUseCase{
		Repo: repo,
	}
}
