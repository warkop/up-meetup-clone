package repository

import (
	"github.com/warkop/up-meetup-clone/api/v1/user/models"
	"gorm.io/gorm"
)

type UserRepositoryProto interface {
	Create(user *models.User) (err error)
	Fetch() (response []*models.User, err error)
	FetchByID(id int64) (response *models.User, err error)
	Update(id int64, user *models.User) error
	Delete(id int64) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepositoryProto {
	return &UserRepository{
		db: connection,
	}
}

func (repo *UserRepository) Create(user *models.User) (err error) {
	err = repo.db.Create(user).Error

	return err
}

func (repo *UserRepository) Fetch() (response []*models.User, err error) {
	err = repo.db.Model(response).Find(&response).Error

	return response, err
}

func (repo *UserRepository) FetchByID(id int64) (response *models.User, err error) {
	err = repo.db.Model(response).Where(`id = ?`, id).Find(&response).Error
	return response, err
}

func (repo *UserRepository) Update(id int64, user *models.User) (err error) {
	ud := make(map[string]interface{})
	ud["name"] = user.Name
	ud["email"] = user.Email

	err = repo.db.Where("id = ?", id).Updates(ud).Error

	return err
}

func (repo *UserRepository) Delete(id int64) (err error) {
	err = repo.db.Where("id = ?", id).Delete(models.User{}).Error

	return err
}
