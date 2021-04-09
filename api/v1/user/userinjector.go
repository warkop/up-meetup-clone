package user

import (
	"github.com/warkop/up-meetup-clone/api/v1/user/domain"
	"github.com/warkop/up-meetup-clone/api/v1/user/gateway/handler/http"
	"github.com/warkop/up-meetup-clone/api/v1/user/repository"
	"github.com/warkop/up-meetup-clone/config"
	"github.com/warkop/up-meetup-clone/lib/db"
)

func ProvideUserRepository() repository.UserRepositoryProto {
	return repository.NewUserRepository(db.InitializeTransaction(config.Orm, "user_"))
}

func ProvideUserUseCase() domain.UserUseCaseProto {
	return domain.NewUserUseCase(ProvideUserRepository())
}

func ProvideUserHandler() http.UserHandlerProto {
	return http.NewUserHandler(ProvideUserUseCase())
}
