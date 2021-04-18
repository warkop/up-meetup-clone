package user

import (
	"github.com/warkop/up-meetup-clone/api/v1/user/domain"
	"github.com/warkop/up-meetup-clone/api/v1/user/gateway/handler/http"
	"github.com/warkop/up-meetup-clone/api/v1/user/repository"
	"gorm.io/gorm"
)

func ProvideUserRepository(db *gorm.DB) repository.UserRepositoryProto {
	return repository.NewUserRepository(db)
}

func ProvideUserUseCase(db *gorm.DB) domain.UserUseCaseProto {
	return domain.NewUserUseCase(ProvideUserRepository(db))
}

func ProvideUserHandler(db *gorm.DB) http.UserHandlerProto {
	return http.NewUserHandler(ProvideUserUseCase(db))
}
